package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const COLOR = 255.99

func checkErr(e error, s string) {
	if e != nil {
		log.Fatal(e, s)
	}
}

func writeColor(w io.Writer, v Vector) {
	_, err := fmt.Fprintf(w, "%d %d %d\n", int(v.X*COLOR), int(v.Y*COLOR), int(v.Z*COLOR))
	checkErr(err, "Error writing to file")
}

func main() {
	imageWidth := 400
	imageHeight := 200

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

	lowerLeft := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	origin := Vector{0.0, 0.0, 0.0}

	////writes each pixel with r/g/b values
	////from top left to bottom right
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth)
			v := float64(j) / float64(imageHeight)

			position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

			// direction = lowerLeft + (u * horizontal) + (v * vertical)
			direction := lowerLeft.Add(position)

			rgb := Ray{origin, direction}.Color()
			writeColor(f, rgb)
		}
	}
}
