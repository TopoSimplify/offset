package offset

import (
	"simplex/lnr"
	"simplex/rng"
	"github.com/intdxdt/geom"
)

//@formatter:off

//euclidean offset distance from dp - anchor line [i, j] to maximum
//vertex at i < k <= j - not maximum offset is may not  be perpendicular
func MaxOffset(linear lnr.Linear, rng *rng.Range) (int, float64) {
	var polyline           = linear.Coordinates()
	var seg           = geom.NewSegment(polyline[rng.I()], polyline[rng.J()])
	var index, offset = rng.J(), 0.0

	if rng.Size() > 1 {
		for _, k := range rng.ExclusiveStride(1) {
			dist := seg.DistanceToPoint(polyline[k])
			if dist >= offset {
				index, offset = k, dist
			}
		}
	}
	return index, offset
}

