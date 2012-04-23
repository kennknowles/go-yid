
package yid

import "testing"
	
var nullable_tests = []struct {
	grammar *Grammar
	nullable bool
}{
	{ Empty, false },
	{ Eps, true },
	{ Token("foo"), false },

	{ Cat(Empty, Eps), false },
	{ Cat(Eps, Empty), false },
	{ Cat(Eps, Eps), true },
	{ Cat(Eps, Token("baz")), false },

	{ Alt(Eps, Empty), true },
	{ Alt(Empty, Eps), true },
	{ Alt(Empty, Empty), false },
	{ Alt(Eps, Token("foo")), true },

	{ rec_alt1, true },
	{ As_then_bs, false },
	{ Russ_cox_exponential, false },
}

func TestNullable(t *testing.T) {
	for _, test_case := range nullable_tests {
		if (test_case.grammar.Nullable() != test_case.nullable) {
			t.Errorf("Nullable test failed (should be %t): %s", test_case.nullable, test_case.grammar.Pretty())
		}
	}
}

