
package yid

// To keep it simplest to read in Go, everything is monomorphic where possible, and the AST is stratified into:
//
// - Grammar: always a lazy value
// - GrammarHead: the underlying struct for pattern matching

// Thunks with types specialized for this use

type Grammar struct {
	thunk func() GrammarHead
	head GrammarHead
}

type GrammarHead interface{} // Should actually be a pointer to one of EpsG, EmptyG, AltG, CatG

func (g *Grammar) Head() GrammarHead {
	if g.head == nil {
		g.head = g.thunk()
	}
	return g.head
}

func makeG(thunk func() GrammarHead) *Grammar {
	return &Grammar { thunk:thunk }
}

func makeH(head GrammarHead) *Grammar {
	return &Grammar { head: head }
}


// To express an ADT in Go, use a type interface {} and structs
// To make it lazy, always wrap in a function 
// (internal bits must convert the function to a thunk value)

// Representation

type EmptyG struct {}  // A grammaruage containing no strings
type EpsG struct {} // "epsilon", A grammaruage containing only the empty string
type TokenG struct { 
	Value string // Empty literal is not considered epsilon
}
type CatG struct { 
	First, Second *Grammar 
}
type AltG struct {
	Left, Right *Grammar 
}
type ReduceG struct {
	Input *Grammar
	Func func(string) string
}

// constructors should intern these
var EmptyH = &EmptyG{}
var EpsH = &EpsG{}

// Constructors - all delayed because it simplified composing them

var Empty *Grammar = makeG(func() GrammarHead { return EmptyH })
var Eps *Grammar = makeG(func() GrammarHead { return EpsH })
func Token(v string) *Grammar {
	return makeG(func() GrammarHead { return &TokenG{ v } })
}
func Cat(first, second *Grammar) *Grammar {
	return makeG(func() GrammarHead { return &CatG{ first, second } })
}
func Alt(left, right *Grammar) *Grammar {
	return makeG(func() GrammarHead { return &AltG{ left, right } })
}
func Reduce(input *Grammar, f func(string) string) *Grammar {
	return makeG(func() GrammarHead { return &ReduceG{ input, f } })
}

