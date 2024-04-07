package rayTracer

func Reflect(v Vector, n Vector) Vector {
	b := 2 * v.Dot(n)
	return v.Subtract(n.MultiplyScalar(b))
}
