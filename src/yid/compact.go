
package yid

func memo_compact(visited map[*Grammar]bool, grammar *Grammar) *Grammar {
	if visited[grammar] {
		return grammar
	}
	
	switch g := grammar.Head().(type) {
	case *EmptyG:
		return Empty
		
	case *EpsG:
		return Eps
		
	case *TokenG:
		return grammar
		
	case *CatG:
		visited[grammar] = true
		
		// Empty . p == p . Empty ==> Empty
		// Eps . p == p . Eps ==> p
		switch first := memo_compact(visited, g.First).Head().(type) {
		case *EmptyG:
			return Empty
			
		case *EpsG:
			return memo_compact(visited, g.Second)
			
		default:
			switch second := memo_compact(visited, g.Second).Head().(type) {
			case *EmptyG:
				return Empty
				
			case *EpsG:
				return memo_compact(visited, g.First)
				
			default:
				g.First = makeH(first)
				g.Second = makeH(second)
				return makeH(g)
			}
		}
		
	case *AltG:
		visited[grammar] = true
		g.Left = memo_compact(visited, g.Left);
		g.Right = memo_compact(visited, g.Right)
		
		// Empty | p == p | Empty ==> p
		// Eps | Eps ==> Eps
		if g.Left == Empty {
			return g.Right
		} else if g.Right == Empty {
			return g.Left
		} else if (g.Left == Eps) && (g.Right == Eps) {
			return Eps
		} else {
			return makeH(g) // Already mutated in place
		}
		panic("Unreachable")
		
	default:
		panic("Non-grammar head")
	}
	panic("Unreachable")
}

func (grammar *Grammar) Compact() *Grammar {
	visited := make(map[*Grammar]bool)
	return memo_compact(visited, grammar)
}
