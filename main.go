package main

import (
	"fmt"
	"goRayTracer/rayTracer"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
)

const (
	COLOR       = 255.99
	imageWidth  = 400
	imageHeight = 200
	numS        = 100 // num of try to antialiasing
)

var (
	white = rayTracer.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	blue  = rayTracer.Vector{X: 0.5, Y: 0.7, Z: 1.0}

	camera = rayTracer.NewCamera()

	sphere = rayTracer.Sphere{
		Center:   rayTracer.Vector{Z: -1},
		Radius:   0.5,
		Material: rayTracer.Lambertian{C: rayTracer.Vector{X: 0.8, Y: 0.3, Z: 0.3}},
	}
	floor = rayTracer.Sphere{
		Center:   rayTracer.Vector{Y: -100.5, Z: -1},
		Radius:   100,
		Material: rayTracer.Lambertian{C: rayTracer.Vector{X: 0.8, Y: 0.8}},
	}
	leftSphere = rayTracer.Sphere{
		Center:   rayTracer.Vector{X: -1, Z: -1},
		Radius:   0.5,
		Material: rayTracer.Metal{C: rayTracer.Vector{X: 0.8, Y: 0.8, Z: 0.8}, Fuzz: 0.0},
	}
	rightSphere = rayTracer.Sphere{
		Center:   rayTracer.Vector{X: 1, Z: -1},
		Radius:   0.5,
		Material: rayTracer.Metal{C: rayTracer.Vector{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 0.3},
	}

	world = rayTracer.World{Elems: []rayTracer.Hittable{&sphere, &floor, &leftSphere, &rightSphere}}
)

func checkErr(e error, s string) {
	if e != nil {
		log.Fatal(e, s)
	}
}

func writeColor(w io.Writer, v rayTracer.Vector) {
	_, err := fmt.Fprintf(w, "%d %d %d\n", int(v.X*COLOR), int(v.Y*COLOR), int(v.Z*COLOR))
	checkErr(err, "Error writing to file")
}

func gradient(r rayTracer.Ray) rayTracer.Vector {
	// unit vector
	unitDirection := r.Direction.Normalize()

	// scale t to be between 0.0 and 1.0
	t := 0.5 * (unitDirection.Y + 1.0)

	// linear blend: blended_value = (1 - t) * white + t * blue
	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func rayColor(r rayTracer.Ray, world rayTracer.Hittable, depth int) rayTracer.Vector {
	isHit, record := world.Hit(r, 0.001, math.MaxFloat64)
	if isHit {
		if depth < 50 {
			isBounced, bouncedRay := record.Bounce(r, record)
			if isBounced {
				newColor := rayColor(bouncedRay, world, depth+1)
				return record.Material.Color().Multiply(newColor)
			}
		}
		return rayTracer.Vector{}
	}
	return gradient(r)
}

func main() {

	f, err := os.Create("out.ppm")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	checkErr(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	checkErr(err, "Error writing to file: %v\n")

	////writes each pixel with r/g/b values
	////from top left to bottom right
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			rgb := rayTracer.Vector{}

			for s := 0; s < numS; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth)
				v := (float64(j) + rand.Float64()) / float64(imageHeight)

				r := camera.RayAt(u, v)
				color := rayColor(r, &world, 0)
				rgb = rgb.Add(color)
			}
			// calc avg color
			rgb = rgb.DivideScalar(float64(numS))
			writeColor(f, rgb)
		}
	}
}
