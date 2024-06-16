package primitives

import (
	"math"
	"math/rand"
)

type Camera struct {
	lowerLeft  Vector
	horizontal Vector
	vertical   Vector
	origin     Vector
	u          Vector
	v          Vector
	w          Vector
	lensRadius float64
}

func NewCamera(lookFrom Vector, lookAt Vector, vFov float64, aspectRatio float64, aperture float64) *Camera {
	c := Camera{}

	c.origin = lookFrom
	c.lensRadius = aperture / 2

	// calc from vertical field of view
	theta := vFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspectRatio * halfHeight

	w := lookFrom.Subtract(lookAt).Normalize()
	u := Vector{Y: 1}.Cross(w).Normalize()
	v := w.Cross(u)

	focusDist := lookFrom.Subtract(lookAt).Length()

	x := u.MultiplyScalar(halfWidth * focusDist)
	y := v.MultiplyScalar(halfHeight * focusDist)

	c.lowerLeft = c.origin.Subtract(x).Subtract(y).Subtract(w.MultiplyScalar(focusDist))
	c.horizontal = x.MultiplyScalar(2)
	c.vertical = y.MultiplyScalar(2)

	c.w = w
	c.u = u
	c.v = v
	return &c
}

func (c *Camera) RayAt(s float64, t float64, rnd *rand.Rand) Ray {
	var randomUnitDisc Vector
	for {
		randomUnitDisc = Vector{X: rnd.Float64(), Y: rnd.Float64()}.MultiplyScalar(2).Subtract(Vector{X: 1, Y: 1})
		if randomUnitDisc.Dot(randomUnitDisc) < 1.0 {
			break
		}
	}

	rd := randomUnitDisc.MultiplyScalar(c.lensRadius)
	offset := c.u.MultiplyScalar(rd.X).Add(c.v.MultiplyScalar(rd.Y))

	horizontal := c.horizontal.MultiplyScalar(s)
	vertical := c.vertical.MultiplyScalar(t)

	origin := c.origin.Add(offset)
	direction := c.lowerLeft.Add(horizontal).Add(vertical).Subtract(c.origin).Subtract(offset)
	return Ray{Origin: origin, Direction: direction}

}
