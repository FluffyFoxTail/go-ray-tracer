package rayTracer

import "math/rand"

var UnitVector = Vector{1, 1, 1}

func VectorInUnitSphere() Vector {
	for {
		r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
		p := r.MultiplyScalar(2.0).Subtract(UnitVector)
		if p.SquaredLength() >= 1.0 {
			return p
		}
	}
}
