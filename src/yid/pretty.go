
package yid;

import "fmt"
import "strconv"

type nonterminal int64

func consume_nonterminal(used *nonterminal) nonterminal {
	*used = nonterminal(int64(*used) + 1)
	return *used
}

func pretty(nonterminals map[*Grammar]nonterminal, used *nonterminal, grammar *Grammar) string {
	if nonterminals[grammar] > 0 {
		return strconv.FormatInt(int64(nonterminals[grammar]), 10)
	} 

	switch g := grammar.Head().(type) {
	case *EmptyG:
		return "∅"

	case *EpsG:
		return "𝜺"

	case *TokenG:
		return fmt.Sprintf("'%s'", g.Value)

	case *CatG:
		// Set up the current nonterminal to cut off loops in recursion, and consume a couple for the lhs and rhs
		nonterminals[grammar] = consume_nonterminal(used)
		return fmt.Sprintf("%d@( %s . %s )", nonterminals[grammar], pretty(nonterminals, used, g.First), pretty(nonterminals, used, g.Second))

	case *AltG:
		nonterminals[grammar] = consume_nonterminal(used)
		return fmt.Sprintf("%d@( %s | %s )", nonterminals[grammar], pretty(nonterminals, used, g.Left), pretty(nonterminals, used, g.Right))
	}

	panic(fmt.Sprintf("Not a grammar: %s", grammar))
}

func (grammar *Grammar) Pretty() string {
	used := new(nonterminal)
	*used = 0
	nonterminals := make(map[*Grammar]nonterminal)
	return pretty(nonterminals, used, grammar)
}
