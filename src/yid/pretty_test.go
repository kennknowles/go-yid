
package yid

import "testing"
	
// Testing that equality is reflexive, symmetric, and terminates on recursive structures

type pretty_test struct {
	name string
	grammar *Grammar
	pretty string
}

var pretty_tests = []pretty_test {
	{ "empty", Empty, "∅" },
	{ "eps"  , Eps, "𝜺" },
	{ "foo"  , Token("foo"), "'foo'" },

	{ "cat"  , Cat(Eps, Eps), "1@( 𝜺 . 𝜺 )" },
	{ "alt"  , Alt(Eps, Empty), "1@( 𝜺 | ∅ )" },

	{ "catcat", Cat(Eps, Cat(Token("foo"), Empty)), "1@( 𝜺 . 2@( 'foo' . ∅ ) )"},

	{ "rec_alt1", rec_alt1, "1@( 𝜺 | 1 )" },
	//{ "as_then_bs", as_then_bs, "1@( 2@( 'a' . 'b' ) | 3@( 'a' . 4@( 1 . 'b' ) ) )" },
	// This randomly passes based on surrounding code; presumably a misunderstanding of == // { "russ_cox",  russ_cox_exponential, "1@( 'one' | 2@( 1 . 3@( '+' . 1 ) ) )" },
}

func TestPretty(t *testing.T) {
	do_test := func(test_case pretty_test) {
		printed := test_case.grammar.Pretty()
		if (printed != test_case.pretty) {
			t.Errorf("Pretty test %s failed.\nExpected:\n%s\n\nGot:\n%s\n\n", test_case.name, test_case.pretty, printed)
		}
	}
	
	for _, test_case := range pretty_tests {
		do_test(test_case)
	}
}
