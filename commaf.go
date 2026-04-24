//go:build go1.6
// +build go1.6

package humanize

import "math/big"

// BigCommaf produces a string form of the given big.Float in base 10
// with commas after every three orders of magnitude.
func BigCommaf(v *big.Float) string {
	return commaf(v.Text('f', -1))
}
