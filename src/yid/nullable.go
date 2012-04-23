
package yid

// De-HOFd, de-closured, Kleene fixpoint calculations
	
// I don't do this nearly as intelligently as I should
func improve_nullable(best_guess, visited map[*Grammar] bool, grammar *Grammar) (changed bool) {
	changed = false
		
	if best_guess[grammar] { // If the best_guess is true, then we are certainly done
		return
	} else if visited[grammar] { // If the grammar has been visited, then we are in a loop and it won't get better
		return
	} else { // Else we can try to make an improvement, cutting off loops via visited[grammar]
		visited[grammar] = true
		
		switch g := grammar.Head().(type) {
		case *EpsG:
			best_guess[grammar] = true
			changed = true

		case *CatG:
			changed = improve_nullable(best_guess, visited, g.First) || improve_nullable(best_guess, visited, g.Second)
			best_guess[grammar] = best_guess[g.First] && best_guess[g.Second]

		case *AltG:
			changed = improve_nullable(best_guess, visited, g.Left) || improve_nullable(best_guess, visited, g.Right)
			best_guess[grammar] = best_guess[g.Left] || best_guess[g.Right]
		}
		changed = changed || best_guess[grammar] // if this grammar became true or a sub-call improved, then something changed
	}
	return
}


func (grammar *Grammar) Nullable() bool {
	// A memo table full of falsity; actually this could be a global weakref hashtable; TODO
	best_guess := make(map[*Grammar]bool)

	for improved := true ; improved ; {
		improved = improve_nullable(best_guess, make(map[*Grammar]bool), grammar)
	}

	return best_guess[grammar]
}
