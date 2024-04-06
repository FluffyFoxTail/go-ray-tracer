package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error, s string) {
	if e != nil {
		log.Fatal(e, s)
	}
}

func main() {
	imageWidth := 400
	imageHeight := 300

	const color = 255.99

	f, err := os.Create("out.ppm")
	defer f.Close()
	check(err, "Error opening file")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	check(err, "Error writing to file")

	for row := imageHeight - 1; row >= 0; row-- {
		for col := 0; col < imageWidth; col++ {
			v := Vector{X: float64(col) / float64(imageWidth), Y: float64(row) / float64(imageHeight), Z: 0.0}

			ir := int(color * v.X)
			ig := int(color * v.Y)
			ib := int(color * v.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writing to file")
		}
	}
}
