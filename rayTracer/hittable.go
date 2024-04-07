package rayTracer

type Hit struct {
	T      float64
	Point  Vector // point
	Normal Vector
	Material
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, Hit)
}
