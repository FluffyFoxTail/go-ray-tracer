package rayTracer

type Metal struct {
	C    Vector
	Fuzz float64
}

// Bounce specular reflection.
// incident_vector - (2 * dot(incident_vector, surface_normal) * surface_normal)
func (m Metal) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := Reflect(input.Direction, hit.Normal)
	// optional ability
	fuzz := VectorInUnitSphere().MultiplyScalar(m.Fuzz)
	bouncedRay := Ray{hit.Point, direction.Add(fuzz)}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}
func (m Metal) Color() Vector {
	return m.C
}
