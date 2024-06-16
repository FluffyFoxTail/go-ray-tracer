package primitives

type Ray struct {
	Origin, Direction Vector
}

// At Point(t)=A+tb
func (r Ray) At(t float64) Vector {
	b := r.Direction.MultiplyScalar(t)
	return r.Origin.Add(b)
}
