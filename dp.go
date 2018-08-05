package offset

import (
	"github.com/intdxdt/geom"
)

//euclidean offset distance from dp - anchor line [i, j] to maximum
//vertex at i < k <= j - not maximum offset is may not  be perpendicular
//note : coordinates of node from begin 0 to n-1 find k
func MaxOffset(coordinates geom.Coords) (int, float64) {
	var n = coordinates.Len() - 1
	var index, offset = n, 0.0
	if n <= 1 {
		return index, offset
	}

	var dist float64
	var a, b = coordinates.Pt(0), coordinates.Pt(n)
	for k := 1; k < n; k++ { //exclusive range between 0 < k < n
		dist = geom.DistanceToPoint(a, b, coordinates.Pt(k))
		if dist >= offset {
			index, offset = k, dist
		}
	}

	return index, offset
}
