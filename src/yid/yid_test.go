
package yid

/*

import "testing"
	
// S = eps | S
func rec_alt1() Grammar {
	s := Alt{ Eps, Eps }
	s.Right = s
	return s
}

// S = eps . S
func rec_cat1() Grammar {
	s := Cat{ Eps, Eps }
	s.Second = s
	return s
}

// S = S . eps
func rec_cat2() Grammar {
	s := &Cat{ &Eps{}, &Eps{} }
	s.First = s
	return s
}

//
// Nullable
//

var nullable_tests = []struct {
	grammar Grammar
	nullable bool
}{
	{ Empty, false },
	{ Eps, true },
	{ Cat(func() Grammar { return Empty }, func() Grammar { return Eps }), false },
	{ Cat(func() Grammar { return Eps }, func() { return Empty }), false },
	{ Cat(func() Grammar { return Eps}, func() { return Eps }}, true },

	{ rec_alt1(), true },
	{ rec_cat1(), false },
	{ rec_cat2(), false },
}

func TestNullable(t *testing.T) {
	for idx, test_case := range nullable_tests {
		if (Nullable(test_case.grammar) != test_case.nullable) {
			t.Errorf("Nullable test at index %d failed test (should be %b)", idx, test_case.nullable)
		}
	}
}

//
// Compact
//

var compact_tests = []struct {
	grammar Grammar
	compacted Grammar
}{
	{ &Empty{}, TheEmpty }, 
	{ &Eps{}, TheEps }, 

	{ &Lit{ "hello" }, &Lit{ "hello" } },

	{ &Cat{ &Empty{}, &Empty{} }, TheEmpty },
	{ &Cat{ &Eps{}, &Empty{} }, TheEmpty },
	{ &Cat{ &Empty{}, &Eps{} }, TheEmpty },
	{ &Cat{ &Eps{}, &Eps{} }, TheEps },
	{ &Cat{ &Eps{}, &Lit{ "hello" } }, &Lit{ "hello" } },
	{ &Cat{ &Lit{ "hello" }, &Eps{} }, &Lit{ "hello" } },
	
	{ &Alt{ &Empty{}, &Empty{} }, TheEmpty },
	{ &Alt{ &Empty{}, &Eps{} }, TheEps },
	{ &Alt{ &Eps{}, &Empty{} }, TheEps },
	{ &Alt{ &Eps{}, &Eps{} }, TheEps }, 
	{ &Alt{ &Empty{}, &Lit{ "foo" } }, &Lit{ "foo" } },
	{ &Alt{ &Lit{ "foo" }, &Empty{} }, &Lit{ "foo" } },
}

func TestCompact(t *testing.T) {
	for idx, test_case := range compact_tests {
		if !Eq(Compact(test_case.grammar), test_case.compacted) {
			t.Errorf("Compact test at index %d failed", idx)
		}
	}
}

//
// Deriv
//

var deriv_tests = []struct {
	grammar Grammar
	next string
	deriv Grammar
}{
	{ &Empty{}, "a", &Empty{} },
	{ &Eps{}, "b", &Empty{} },

	{ &Lit{"x"}, "x", &Eps{} },
	{ &Lit{"x"}, "y", &Empty{} },

	{ &Cat{ &Eps{}, &Eps{} }, "foo", &Empty{} },
	{ &Cat{ &Lit{ "foo" }, &Eps{} }, "foo", &Eps{} },
	{ &Cat{ &Lit{ "foo" }, &Lit{ "baz" } }, "foo", &Lit{ "baz" } },
}

func TestDeriv(t *testing.T) {
	for idx, test_case := range deriv_tests {
		if !Eq(Compact(Deriv(test_case.next, test_case.grammar)), test_case.deriv) {
			t.Errorf("Deriv test at index %d failed test", idx)
		}
	}
}

//
// TODO: test equiv-acceptance on many strings of a grammar and its compaction
//
*/
