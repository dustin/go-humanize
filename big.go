package humanize

import (
	"math/big"
)

// order of magnitude (to a max order)
func oomm(n, b *big.Int, maxmag int) (float64, int) {
	mag := 0
	m := &big.Int{}
	for n.Cmp(b) >= 0 {
		n.DivMod(n, b, m)
		mag++
		if mag == maxmag && maxmag >= 0 {
			break
		}
	}
	// Stopping at maxmag leaves n unreduced, so it can be larger than an
	// int64 and Int64 would return garbage. Convert through big.Float
	// instead, which rounds the same way as float64(int64) does when the
	// value fits.
	val, _ := new(big.Float).SetInt(n).Float64()
	return val + (float64(m.Int64()) / float64(b.Int64())), mag
}

// total order of magnitude
// (same as above, but with no upper limit)
func oom(n, b *big.Int) (float64, int) {
	mag := 0
	m := &big.Int{}
	for n.Cmp(b) >= 0 {
		n.DivMod(n, b, m)
		mag++
	}
	return float64(n.Int64()) + (float64(m.Int64()) / float64(b.Int64())), mag
}
