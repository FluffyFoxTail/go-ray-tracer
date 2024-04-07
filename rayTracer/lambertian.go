package rayTracer

type Lambertian struct {
	C Vector
}

// Bounce Lambertian reflectance
func (l Lambertian) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := hit.Normal.Add(VectorInUnitSphere())
	return true, Ray{hit.Point, direction}
}
func (l Lambertian) Color() Vector {
	return l.C
}
