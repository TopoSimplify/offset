package offset

import (
	"simplex/rng"
	"github.com/intdxdt/geom"
)

//euclidean offset distance from dp - anchor line [i, j] to maximum
//vertex at i < k <= j - not maximum offset is may not  be perpendicular
func MaxOffset(coordinates []*geom.Point, rng *rng.Range) (int, float64) {
	var seg = geom.NewSegment(coordinates[rng.I()], coordinates[rng.J()])
	var index, offset = rng.J(), 0.0

	if rng.Size() > 1 {
		for _, k := range rng.ExclusiveStride(1) {
			var dist = seg.DistanceToPoint(coordinates[k])
			if dist >= offset {
				index, offset = k, dist
			}
		}
	}
	return index, offset
}
