
package yid

func (grammar *Grammar) Recognize(input []string) bool {
	curr := grammar
	for _, next := range input {
		curr = curr.Deriv(next).Compact()
	}
	return curr.Nullable()
}
