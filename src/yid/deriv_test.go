
package yid

import "testing"

var deriv_tests = []struct {
	grammar *Grammar
	next string
	deriv *Grammar
}{
	{ Empty, "a", Empty },
	{ Eps, "b", Empty },

	{ Token("x"), "x", Eps },
	{ Token("x"), "y", Empty },

	{ Cat(Eps, Eps), "foo", Empty },
	{ Cat(Token("foo"), Eps), "foo", Eps },
	{ Cat(Token("foo"), Token("baz")), "foo", Token("baz") },

	{ Alt(Eps, Eps), "foo", Empty },
	{ Alt(Eps, Token("foo")), "foo", Eps },
	{ Alt(Eps, Token("baz")), "foo", Empty },

	//{ rec_alt1, "foo", rec_alt1_deriv },
	//{ as_then_bs, "a", as_then_bs_deriv_a },
	//{ as_then_bs, "b", Empty },
}

func TestDeriv(t *testing.T) {
	for _, test_case := range deriv_tests {
		derived := test_case.grammar.Deriv(test_case.next).Compact()
		if !derived.Eq(test_case.deriv) {
			t.Errorf("Deriv w.r.t %s of %s failed.\nExpected: %s\nGot: %s\n", test_case.next, test_case.grammar.Pretty(), test_case.deriv.Pretty(), derived.Pretty())
		}
	}
}

