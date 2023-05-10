package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"math"
)

func findClosestStorm(img gocv.Mat, x int, y int) (int, int) {
	// Converte l'immagine in scala di grigi
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// Applica una soglia per estrarre i pixel che rappresentano i temporali
	thresh := gocv.NewMat()
	defer thresh.Close()
	gocv.Threshold(gray, &thresh, 200, 255, gocv.ThresholdBinary)

	// Trova i contorni nella mappa dei temporali
	contours := gocv.FindContours(thresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	// Calcola la distanza di ogni temporale dal punto di interesse
	distances := make([]float64, len(contours))
	for i, contour := range contours {
		// Calcola il centroide del contorno
		m := gocv.Moments(contour, true)
		cX := m["m10"] / m["m00"]
		cY := m["m01"] / m["m00"]

		// Calcola la distanza Euclidea dal punto di interesse
		distance := math.Sqrt(math.Pow(cX-float64(x), 2) + math.Pow(cY-float64(y), 2))
		distances[i] = distance
	}

	// Trova l'indice del temporale più vicino
	minDistanceIndex := 0
	minDistance := distances[0]
	for i, distance := range distances {
		if distance < minDistance {
			minDistance = distance
			minDistanceIndex = i
		}
	}

	// Restituisce le coordinate del temporale più vicino
	centroid := gocv.Moments(contours[minDistanceIndex], true)
	cX := centroid["m10"] / centroid["m00"]
	cY := centroid["m01"] / centroid["m00"]

	return int(cX), int(cY)
}

func main() {
	// Carica l'immagine radar
	img := gocv.IMRead("mappa_radar.jpg", gocv.IMReadColor)
	defer img.Close()

	// Coordinate del punto di interesse
	poiX := 100
	poiY := 200

	// Trova le coordinate del temporale più vicino
	closestX, closestY := findClosestStorm(img, poiX, poiY)

	fmt.Println("Coordinate del temporale più vicino:")
	fmt.Println("X:", closestX)
	fmt.Println("Y:", closestY)
}
