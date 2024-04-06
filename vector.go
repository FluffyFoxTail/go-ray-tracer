package main

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
	l := v.Length()
	return Vector{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector) Cross(ov Vector) Vector {
	return Vector{
		v.Y*ov.Z - v.Z*ov.Y,
		v.Z*ov.X - v.X*ov.Z,
		v.X*ov.Y - v.Y*ov.Z,
	}
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
