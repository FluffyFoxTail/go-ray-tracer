package render

import (
	"image"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/Fluffyfoxtail/go-ray-tracer/ray-tracer/primitives"
)

const (
	maxDepth = 50
	tMin     = 0.001
)

// Do performs the render, sampling each pixel the provided number of times
func Do(world primitives.Hittable, camera *primitives.Camera, cpus, samples, width, height int, ch chan<- int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

			for row := i; row < height; row += cpus {
				for col := 0; col < width; col++ {
					rgb := sample(world, camera, rnd, samples, width, height, col, row)
					img.Set(col, height-row-1, rgb)
				}
				ch <- 1
			}
		}(i)
	}

	wg.Wait()
	return img
}

// sample create samples rays to prevent antialiasing
func sample(world primitives.Hittable, camera *primitives.Camera, rnd *rand.Rand, samples, width, height, i, j int) primitives.Color {
	rgb := primitives.Color{}

	for s := 0; s < samples; s++ {
		u := (float64(i) + rnd.Float64()) / float64(width)
		v := (float64(j) + rnd.Float64()) / float64(height)

		ray := camera.RayAt(u, v, rnd)
		rgb = rgb.Add(rayColor(ray, world, rnd, 0))
	}

	// average
	return rgb.DivideScalar(float64(samples))
}

func rayColor(r primitives.Ray, world primitives.Hittable, rnd *rand.Rand, depth int) primitives.Color {
	isHit, record := world.Hit(r, tMin, math.MaxFloat64)
	if isHit {
		if depth < maxDepth {
			isBounced, bouncedRay := record.Bounce(r, record, rnd)
			if isBounced {
				newColor := rayColor(bouncedRay, world, rnd, depth+1)
				return record.Material.Color().Multiply(newColor)
			}
		}
		return primitives.Black
	}
	return primitives.Gradient(primitives.White, primitives.Blue, r.Direction.Normalize().Y)
}
