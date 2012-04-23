
package yid

// De-HOFd, de-closured, memoized derivative calculations

func memo_deriv(memo map[*Grammar]*Grammar, next string, grammar *Grammar) *Grammar {
		
	if memo[grammar] != nil {
		return memo[grammar]
	} 

	switch g := grammar.Head().(type) {
	case *EmptyG:
		return grammar // save an allocation
		
	case *EpsG:
		return Empty

	case *TokenG:
		if g.Value == next {
			return Eps
		} else {
			return Empty
		}
		
	case *CatG:
		if g.First.Nullable() {
			memo[grammar] = makeG(func() GrammarHead { return Alt(memo_deriv(memo, next, g.Second),  Cat(memo_deriv(memo, next, g.First),  g.Second)).Head() })
		} else {
			memo[grammar] = makeG(func() GrammarHead { return Cat(memo_deriv(memo, next, g.First), g.Second).Head() })
		}
		return memo[grammar]

	case *AltG:
		memo[grammar] = makeG(func() GrammarHead { return Alt(memo_deriv(memo, next, g.Left), memo_deriv(memo, next, g.Right)).Head() })
		return memo[grammar]
		
	default:
		panic("Non-grammar in deriv")
	}
	
	panic("Unreachable")
}

func (grammar *Grammar) Deriv(next string) *Grammar {
	return memo_deriv(make(map[*Grammar]*Grammar), next, grammar)
}
