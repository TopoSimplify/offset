package offset

import (
	"simplex/lnr"
	"simplex/rng"
	"github.com/intdxdt/vect"
)

//computes Synchronized Euclidean Distance
func MaxSEDOffset(polyline lnr.Linear, rng *rng.Range) (int, float64) {
	var t = 2
	var pln = polyline.Coordinates()
	var index, offset = rng.J(), 0.0
	var a, b = pln[rng.I()], pln[rng.J()]
	var opts = &vect.Options{A: a, B: b, At: &a[t], Bt: &b[t]}
	var segvect = vect.NewVect(opts)

	if rng.Size() > 1 {
		for _, k := range rng.ExclusiveStride(1) {
			sedvect := segvect.SEDVector(pln[k], pln[k][t])
			dist := sedvect.Magnitude()
			if dist >= offset {
				index, offset = k, dist
			}
		}
	}
	return index, offset
}
