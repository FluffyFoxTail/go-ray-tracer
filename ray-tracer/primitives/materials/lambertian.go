package materials

import (
	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives"
	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives/utils"

	"math/rand"
)

type Lambertian struct {
	Attenuation primitives.Color
}

// Bounce Lambertian reflectance
func (l Lambertian) Bounce(input primitives.Ray, hit primitives.Hit, rnd *rand.Rand) (bool, primitives.Ray) {
	direction := hit.Normal.Add(utils.VectorInUnitSphere(rnd))
	return true, primitives.Ray{Origin: hit.Point, Direction: direction}
}
func (l Lambertian) Color() primitives.Color {
	return l.Attenuation
}
