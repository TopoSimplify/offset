package offset

import (
	"github.com/intdxdt/vect"
	"github.com/intdxdt/geom"
)

//computes Synchronized Euclidean Distance
func MaxSEDOffset(coordinates []*geom.Point) (int, float64) {
	var t = 2
	var n = len(coordinates) - 1
	var index, offset = n, 0.0
	if n <= 1 {
		return index, offset
	}

	var a, b = coordinates[0], coordinates[n]
	var opts = &vect.Options{A: a, B: b, At: &a[t], Bt: &b[t]}
	var dist float64

	var pt *geom.Point
	var vec = vect.NewVect(opts)
	for k := 0; k <= n; k++ {
		pt = coordinates[k]
		dist = vec.SEDVector(pt, pt[geom.Z]).Magnitude()
		if dist >= offset {
			index, offset = k, dist
		}
	}
	return index, offset
}
