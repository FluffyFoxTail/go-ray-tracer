package rayTracer

type Ray struct {
	Origin, Direction Vector
}

// At P(t)=A+tb
func (r Ray) At(t float64) Vector {
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}
