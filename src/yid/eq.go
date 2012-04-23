
package yid;

type maybe_bool int

const (
	DUNNO            = 0
	DEFINITELY_FALSE = 1
	DEFINITELY_TRUE  = 2
)

func memo_eq(memo map[*Grammar]map[*Grammar]maybe_bool, grammar1, grammar2 *Grammar) bool {
	if memo[grammar1] == nil {
		memo[grammar1] = make(map[*Grammar]maybe_bool)
	}

	if grammar1 == grammar2 {
		memo[grammar1][grammar2] = DEFINITELY_TRUE
		return true
	}

	if memo[grammar1][grammar2] != DUNNO {
		return (memo[grammar1][grammar2] == DEFINITELY_TRUE)
	}

	switch g1 := grammar1.Head().(type) {
	case *EmptyG:
		switch grammar2.Head().(type) {
		case *EmptyG:
			memo[grammar1][grammar2] = DEFINITELY_TRUE
		default:
			memo[grammar1][grammar2] = DEFINITELY_FALSE
		}
		
	case *EpsG:
		switch grammar2.Head().(type) {
		case *EpsG:
			memo[grammar1][grammar2] = DEFINITELY_TRUE
		default:
			memo[grammar1][grammar2] = DEFINITELY_FALSE
		}

	case *TokenG:
		switch g2 := grammar2.Head().(type) {
		case *TokenG:
			if g1.Value == g2.Value {
				memo[grammar1][grammar2] = DEFINITELY_TRUE
			} else {
				memo[grammar1][grammar2] = DEFINITELY_FALSE
			}
		default:
			memo[grammar1][grammar2] = DEFINITELY_FALSE
		}

	case *CatG:
		switch g2 := grammar2.Head().(type) {
		case *CatG:
			memo[grammar1][grammar2] = DEFINITELY_TRUE // If we loop, then they are equal
			if ! (memo_eq(memo, g1.First, g2.First) && memo_eq(memo, g1.Second, g2.Second)) {
				memo[grammar1][grammar2] = DEFINITELY_FALSE
			}
		}

	case *AltG:
		switch g2 := grammar2.Head().(type) {
		case *AltG:
			memo[grammar1][grammar2] = DEFINITELY_TRUE // If we loop, then they are equal
			if ! (memo_eq(memo, g1.Left, g2.Left) && memo_eq(memo, g1.Right, g2.Right)) {
				memo[grammar1][grammar2] = DEFINITELY_FALSE
			}
		}

	default:
		memo[grammar1][grammar2] = DEFINITELY_FALSE
	}

	return (memo[grammar1][grammar2] == DEFINITELY_TRUE)
}

func (grammar1 *Grammar) Eq(grammar2 *Grammar) bool {
	memo := make(map[*Grammar]map[*Grammar]maybe_bool)
	return memo_eq(memo, grammar1, grammar2)
}
