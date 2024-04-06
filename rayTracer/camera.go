package rayTracer

type Camera struct {
	lowerLeft  Vector
	horizontal Vector
	vertical   Vector
	origin     Vector
}

func NewCamera() *Camera {
	return &Camera{
		lowerLeft:  Vector{-2.0, -1.0, -1.0},
		horizontal: Vector{4.0, 0.0, 0.0},
		vertical:   Vector{0.0, 2.0, 0.0},
		origin:     Vector{0.0, 0.0, 0.0},
	}
}

func (c *Camera) RayAt(u float64, v float64) Ray {
	position := c.position(u, v)
	direction := c.direction(position)
	return Ray{c.origin, direction}
}

func (c *Camera) position(u float64, v float64) Vector {
	horizontal := c.horizontal.MultiplyScalar(u)
	vertical := c.vertical.MultiplyScalar(v)
	return horizontal.Add(vertical)
}

// direction = lowerLeft + (u * horizontal) + (v * vertical)
func (c *Camera) direction(position Vector) Vector {
	return c.lowerLeft.Add(position)
}
