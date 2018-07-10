package offset

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/vect"
	"github.com/intdxdt/geom"
)

//@formatter:off
//computes Synchronized Euclidean Distance
func MaxSEDOffset(coordinates []*geom.Point) (int, float64) {

	var n = len(coordinates) - 1
	var index, offset = n, 0.0
	if n <= 1 {
		return index, offset
	}

	var pt *geom.Point
	var dist, m, ptx, pty, ptt, px, py float64

	var a, b = coordinates[0], coordinates[n]
	var ax, ay  = a[geom.X], a[geom.Y]
	var at , bt = a[geom.Z], b[geom.Z]

	var v = vect.NewVector(a, b)
	var mij = v.Magnitude()
	var fb = direction(v[geom.X], v[geom.Y])
	var vx, vy = math.Cos(fb), math.Sin(fb)
	var dt = bt - at

	for k := 1; k < n; k++ { //exclusive range between 0 < k < n
		pt = coordinates[k]
		ptx, pty, ptt = pt[geom.X], pt[geom.Y], pt[geom.Z]

		m = (mij / dt) * (ptt - at)
		px, py = ax + (m * vx), ay + (m * vy)
		dist = math.Hypot(ptx - px, pty - py)

		if dist >= offset {
			index, offset = k, dist
		}
	}
	return index, offset
}

//Dir computes direction in radians - counter clockwise from x-axis.
func direction(x, y float64) float64 {
	var d = math.Atan2(y, x)
	if d < 0 {
		d += math.Tau
	}
	return d
}