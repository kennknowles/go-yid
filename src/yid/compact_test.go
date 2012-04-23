
package yid

import "testing"

var compact_tests = []struct {
	grammar *Grammar
	compacted *Grammar
}{
	{ Empty, Empty }, 
	{ Eps, Eps }, 

	{ Token("foo"), Token("foo") },

	{ Cat(Empty, Empty)       , Empty        },
	{ Cat(Eps, Empty)         , Empty        },
	{ Cat(Empty, Eps)         , Empty        },
	{ Cat(Eps, Eps)           , Eps          },
	{ Cat(Eps, Token("foo"))  , Token("foo") },
	{ Cat(Token("foo"), Eps)  , Token("foo") },
	
	{ Alt(Empty, Empty)           , Empty         },
	{ Alt(Empty, Eps)             , Eps           },
	{ Alt(Eps, Empty)             , Eps           },
	{ Alt(Eps, Eps)               , Eps           }, 
	{ Alt(Empty, Token("foo"))    , Token("foo")  },
	{ Alt(Token("foo"), Empty)    , Token("foo")  },

	{ rec_alt1             , rec_alt1             },
	{ Russ_cox_exponential , Russ_cox_exponential },
}

func TestCompact(t *testing.T) {
	for _, test_case := range compact_tests {
		compacted := test_case.grammar.Compact()
		if !compacted.Eq(test_case.compacted) {
			t.Errorf("Compacted %s failed.\nExpected: %s\nGot: %s", test_case.grammar.Pretty(), test_case.compacted.Pretty(), compacted.Pretty())
		}
	}
}
