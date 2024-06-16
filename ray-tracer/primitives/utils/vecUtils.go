package utils

import (
	rt "github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives"
	"math/rand"
)

var UnitVector = rt.Vector{X: 1, Y: 1, Z: 1}

func VectorInUnitSphere(rnd *rand.Rand) rt.Vector {
	for {
		r := rt.Vector{X: rnd.Float64(), Y: rnd.Float64(), Z: rnd.Float64()}
		p := r.MultiplyScalar(2.0).Subtract(UnitVector)
		if p.SquaredLength() >= 1.0 {
			return p
		}
	}
}
