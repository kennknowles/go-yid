
package yid

import "testing"
	
// Testing that equality is reflexive, symmetric, and terminates on recursive structures

type eq_test struct {
	grammar1 *Grammar
	grammar2 *Grammar
	equal bool
}

var atomic_eq_tests = []eq_test {
	{ Empty, Empty, true },
	{ Eps, Eps, true },
	{ Token("foo"), Token("foo"), true },

	{ Empty, Eps, false },
	{ Empty, Token("baz"), false },
	{ Eps, Token("baz"), false },
	{ Token("foo"), Token("baz"), false },
}

var recursive_tests = []eq_test {
	{ rec_alt1, rec_alt1, true },
	{ As_then_bs, As_then_bs, true },
	{ Russ_cox_exponential, As_then_bs, false },
}

func compound_eq_tests(lhs_tests, rhs_tests []eq_test) []eq_test {
	tests := []eq_test {}

	for _, lhs := range lhs_tests {
		for _, rhs := range rhs_tests {
			tests = append(tests, []eq_test { 
				{ Cat(lhs.grammar1, rhs.grammar1), Cat(lhs.grammar2, rhs.grammar2), lhs.equal && rhs.equal },
				{ Alt(lhs.grammar1, rhs.grammar1), Alt(lhs.grammar2, rhs.grammar2), lhs.equal && rhs.equal },
			}...)
		}
	}
	
	return tests
}

func depth_n_tests(n int) (tests []eq_test) {
	tests = []eq_test {}
	if n > 0 {
		tests = append(atomic_eq_tests, compound_eq_tests(depth_n_tests(n-1), depth_n_tests(n-1))...)
	}
	return
}

func TestEq(t *testing.T) {
	do_test := func(idx int, test_case eq_test) {
		if (!test_case.grammar1.Eq(test_case.grammar1)) {
			t.Errorf("Eq test at index %d failed reflexivity!", idx, test_case.equal)
		}

		if (!test_case.grammar2.Eq(test_case.grammar2)) {
			t.Errorf("Eq test at index %d failed reflexivity!", idx, test_case.equal)
		}

		if (test_case.grammar1.Eq(test_case.grammar2) != test_case.equal) {
			t.Errorf("Eq test at index %d failed (should be %b)", idx, test_case.equal)
		}
		
		if (test_case.grammar2.Eq(test_case.grammar1) != test_case.equal) {
			t.Errorf("Eq test at index %d failed in reverse (should be %b)", idx, test_case.equal)
		}
	}
	
	for idx, test_case := range depth_n_tests(2) {
		do_test(idx, test_case)
	}

	for idx, test_case := range recursive_tests {
		do_test(idx, test_case)
	}
}
