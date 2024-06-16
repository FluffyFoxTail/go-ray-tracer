package primitives

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Normalize() Vector {
	return v.DivideScalar(v.Length())
}

func (v Vector) Cross(ov Vector) Vector {
	return Vector{
		v.Y*ov.Z - v.Z*ov.Y,
		v.Z*ov.X - v.X*ov.Z,
		v.X*ov.Y - v.Y*ov.Z,
	}
}

func (v Vector) Reflect(ov Vector) Vector {
	b := 2 * v.Dot(ov)
	return v.Subtract(ov.MultiplyScalar(b))
}

func (v Vector) Refract(ov Vector, n float64) (bool, Vector) {
	uv := v.Normalize()
	uo := ov.Normalize()
	dt := uv.Dot(uo)
	discriminant := 1.0 - (n*n - (1 - dt*dt))
	if discriminant > 0 {
		a := uv.Subtract(ov.MultiplyScalar(dt)).MultiplyScalar(n)
		b := ov.MultiplyScalar(math.Sqrt(discriminant))
		return true, a.Subtract(b)
	}
	return false, Vector{}
}

func (v Vector) Add(ov Vector) Vector {
	return Vector{
		v.X + ov.X, v.Y + ov.Y, v.Z + ov.Z}
}

func (v Vector) Subtract(ov Vector) Vector {
	return Vector{v.X - ov.X, v.Y - ov.Y, v.Z - ov.Z}
}

func (v Vector) Multiply(ov Vector) Vector {
	return Vector{v.X * ov.X, v.Y * ov.Y, v.Z * ov.Z}
}

func (v Vector) Divide(ov Vector) Vector {
	return Vector{v.X / ov.X, v.Y / ov.Y, v.Z / ov.Z}
}

func (v Vector) AddScalar(t float64) Vector {
	v.X += t
	v.Y += t
	v.Z += t
	return v
}

func (v Vector) SubtractScalar(t float64) Vector {
	v.X -= t
	v.Y -= t
	v.Z -= t
	return v
}

func (v Vector) MultiplyScalar(t float64) Vector {
	v.X *= t
	v.Y *= t
	v.Z *= t
	return v
}

func (v Vector) DivideScalar(t float64) Vector {
	v.X /= t
	v.Y /= t
	v.Z /= t
	return v
}
