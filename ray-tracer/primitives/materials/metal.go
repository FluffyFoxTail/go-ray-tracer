package materials

import (
	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives"
	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives/utils"

	"math/rand"
)

type Metal struct {
	Attenuation primitives.Color
	Fuzz        float64
}

// Bounce specular reflection.
// incident_vector - (2 * dot(incident_vector, surface_normal) * surface_normal)
func (m Metal) Bounce(input primitives.Ray, hit primitives.Hit, rnd *rand.Rand) (bool, primitives.Ray) {
	direction := input.Direction.Reflect(hit.Normal)
	// optional ability
	fuzz := utils.VectorInUnitSphere(rnd).MultiplyScalar(m.Fuzz)
	bouncedRay := primitives.Ray{Origin: hit.Point, Direction: direction.Add(fuzz)}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}
func (m Metal) Color() primitives.Color {
	return m.Attenuation
}
