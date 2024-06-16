package materials

import (
	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives"
	"math"
	"math/rand"
)

type Glass struct {
	Index float64
}

func (g Glass) Bounce(input primitives.Ray, hit primitives.Hit, rnd *rand.Rand) (bool, primitives.Ray) {
	var outwardNormal primitives.Vector
	var niOverNt, cosine float64

	if input.Direction.Dot(hit.Normal) > 0 {
		outwardNormal = hit.Normal.MultiplyScalar(-1)
		niOverNt = g.Index

		a := input.Direction.Dot(hit.Normal) * g.Index
		b := input.Direction.Length()

		cosine = a / b
	} else {
		outwardNormal = hit.Normal
		niOverNt = 1.0 / g.Index

		a := input.Direction.Dot(hit.Normal) * g.Index
		b := input.Direction.Length()

		cosine = -a / b
	}

	var success bool
	var refracted primitives.Vector
	var reflectProbability float64

	if success, refracted = input.Direction.Refract(outwardNormal, niOverNt); success {
		reflectProbability = g.schlick(cosine)
	} else {
		reflectProbability = 1.0
	}

	if rnd.Float64() < reflectProbability {
		reflected := input.Direction.Reflect(hit.Normal)
		return true, primitives.Ray{Origin: hit.Point, Direction: reflected}
	}

	return true, primitives.Ray{Origin: hit.Point, Direction: refracted}
}

func (g Glass) schlick(cos float64) float64 {
	r0 := (1 - g.Index) / (1 + g.Index)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func (g Glass) Color() primitives.Color {
	return primitives.Color{R: 1.0, G: 1.0, B: 1.0}
}
