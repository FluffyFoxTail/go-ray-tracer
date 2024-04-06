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

	sphere = rayTracer.Sphere{Center: rayTracer.Vector{Z: -1}, Radius: 0.5}
	floor  = rayTracer.Sphere{Center: rayTracer.Vector{Y: -100.5, Z: -1}, Radius: 100}

	world = rayTracer.World{Elems: []rayTracer.Hitable{&sphere, &floor}}
)

func checkErr(e error, s string) {
	if e != nil {
		log.Fatal(e, s)
	}
}

func gradient(v *rayTracer.Vector) rayTracer.Vector {
	// scale t to be between 0.0 and 1.0
	t := 0.5 * (v.Y + 1.0)

	// linear blend: blended_value = (1 - t) * white + t * blue
	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func rayColor(r *rayTracer.Ray, h rayTracer.Hitable) rayTracer.Vector {
	isHit, record := h.Hit(r, 0.0, math.MaxFloat64)
	if isHit {
		return record.Normal.AddScalar(1.0).MultiplyScalar(0.5)
	}

	// make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()
	return gradient(&unitDirection)
}

func writeColor(w io.Writer, v rayTracer.Vector) {
	_, err := fmt.Fprintf(w, "%d %d %d\n", int(v.X*COLOR), int(v.Y*COLOR), int(v.Z*COLOR))
	checkErr(err, "Error writing to file")
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
				color := rayColor(&r, &world)
				rgb = rgb.Add(color)
			}
			// calc avg color
			rgb = rgb.DivideScalar(float64(numS))
			writeColor(f, rgb)
		}
	}
}
