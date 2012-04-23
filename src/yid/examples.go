
package yid

// S -> eps | S 
// ------------
// 1@( eps | 1 )
var rec_alt1 = func () *Grammar {
	s := &Grammar{}
	s.thunk = func() GrammarHead { return &AltG{ Eps, s } }
	return s
}()

// d wrt any token - simplifying to 
var rec_alt1_deriv = func() *Grammar {
	s := &Grammar{}
	s.thunk = func() GrammarHead { return &AltG{ Empty, s } }
	return s
}()

// S -> ab 
// S -> aSb
// -------
// 1@(
//    2@( 'a' . 'b' ) 
//    |
//    3@( 'a' . 4@( 1 . 'b' ) )
//   )
var As_then_bs = func() *Grammar {
	s := &Grammar{}

	terminal := Cat(Token("a"), Token("b"))
	recurse := Cat(Token("a"), Cat(s, Token("b")))

	s.thunk = func() GrammarHead { return &AltG{ terminal, recurse } }
	return s
}()

// D_a(as_then_bs)...
// S_a -> b
// S_a -> Sb
var as_then_bs_deriv_a = func() *Grammar {
	s := &Grammar{}
	
	terminal := Token("b")
	recurse  := Cat(s, Token("b"))

	s.thunk = func() GrammarHead { return &AltG{ terminal, recurse } }
	return s
}()

// S -> one
// S -> S + S
// ----------
// 1@(
//   'one'
//   |
//   2@( 1 . 3@( '+' . 1 ) )
//   )
var Russ_cox_exponential = func() *Grammar {
	s := &Grammar{}

	terminal := Token("one")
	recurse  := Cat(s, Cat(Token("+"), s))

	s.thunk = func() GrammarHead { return &AltG{ terminal, recurse } }
	return s
}()

