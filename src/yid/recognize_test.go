
package yid

import "testing"

var recog_tests = []struct {
	grammar *Grammar
	input []string
	accept bool
}{
	{ Empty, []string{"a"}, false },
	{ Eps, []string{"b"}, false },
	{ Token("b"), []string{"b"}, true },

	{ Cat(Eps, Eps), []string{}, true },
	{ Cat(Eps, Empty), []string{}, false },
	{ Cat(Eps, Token("c")), []string{"c"}, true },
	{ Cat(Token("c"), Eps), []string{"c"}, true },
	{ Cat(Token("b"), Token("c")), []string{"b", "c"}, true },
	{ Cat(Token("b"), Token("a")), []string{"b", "c"}, false },

	{ Alt(Eps, Eps), []string{}, true },
	{ Alt(Eps, Token("foo")), []string{}, true },
	{ Alt(Eps, Token("foo")), []string{"foo"}, true },
	{ Alt(Eps, Token("baz")), []string{"foo"}, false },

	{ rec_alt1, []string{}, true },
	{ rec_alt1, []string{"biz"}, false },

	{ As_then_bs, []string{"a", "b"}, true },
	{ As_then_bs, []string{"a", "z"}, false },
	{ As_then_bs, []string{"a", "a", "a", "b", "b", "b"}, true },

	{ Russ_cox_exponential, []string{"one"}, true},
	{ Russ_cox_exponential, []string{"one", "+"}, false},
	{ Russ_cox_exponential, []string{"one", "+", "one"}, true},
	{ Russ_cox_exponential, []string{"one", "+", "one", "+"}, false},
	{ Russ_cox_exponential, []string{"one", "+", "one", "+", "one"}, true},
}

func TestRecognize(t *testing.T) {
	for _, test_case := range recog_tests {
		recog := test_case.grammar.Recognize(test_case.input)
		if recog != test_case.accept {
			t.Errorf("Recog test for %s on input %v (should be %t)\n", test_case.grammar.Pretty(), test_case.input, test_case.accept)
		}
	}
}

