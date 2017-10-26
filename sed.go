package offset

import (
	"simplex/rng"
	"github.com/intdxdt/vect"
	"github.com/intdxdt/geom"
)

//computes Synchronized Euclidean Distance
func MaxSEDOffset(coordinates []*geom.Point, rng *rng.Range) (int, float64) {
	var t = 2
	var index, offset = rng.J(), 0.0
	var a, b = coordinates[rng.I()], coordinates[rng.J()]
	var opts = &vect.Options{A: a, B: b, At: &a[t], Bt: &b[t]}
	var segVect = vect.NewVect(opts)
	var sedVect *vect.Vect
	var dist float64

	if rng.Size() > 1 {
		for _, k := range rng.ExclusiveStride(1) {
			sedVect = segVect.SEDVector(coordinates[k], coordinates[k][t])
			dist = sedVect.Magnitude()
			if dist >= offset {
				index, offset = k, dist
			}
		}
	}
	return index, offset
}
