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



func creaGIF()	{
	// funzione per la generazione di un file GIF animato usando 10 file JPG indicati in un vettore di stringhe
	// il file GIF viene salvato nella cartella corrente con il nome "animazione.gif"
	// il tempo di visualizzazione di ogni frame è di 100 ms
	// il loop è infinito
	// il numero di frame è uguale al numero di elementi del vettore di stringhe
	// il numero di cicli è uguale al numero di elementi del vettore di stringhe
	// il numero di cicli è uguale a 0 se si vuole un loop infinito
	// il numero di cicli è uguale a -1 se si vuole un loop infinito
	// il numero di cicli è uguale a 1 se si vuole un loop infinito
	// il numero di cicli è uguale a 2 se si vuole un loop infinito	

	// importa il package "image" per la gestione delle immagini
	// importa il package "image/gif" per la gestione dei file GIF
	// importa il package "image/jpeg" per la gestione dei file JPG
	// importa il package "os" per la gestione del sistema operativo
	// importa il package "fmt" per la gestione dei messaggi di errore
	// importa il package "time" per la gestione del tempo
	// importa il package "strconv" per la gestione delle conversioni di tipo
	// importa il package "strings" per la gestione delle stringhe
	// importa il package "log" per la gestione dei messaggi di errore
	// importa il package "io/ioutil" per la gestione dei file

	// crea un vettore di stringhe con i nomi dei file JPG
	var files []string
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")
	files = append(files, "mappa_radar.jpg")

	// crea un vettore di immagini JPG
	var images []*image.Image

	// carica le immagini JPG
	for _, file := range files {
		// apre il file JPG
		reader, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// decodifica il file JPG
		image, err := jpeg.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		// chiude il file JPG
		reader.Close()

		// aggiunge l'immagine JPG al vettore di immagini
		images = append(images, &image)
	}

	// crea un file GIF
	file, err := os.Create("animazione.gif")
	if err != nil {
		log.Fatal(err)
	}

	// chiude il file GIF
	defer file.Close()
	





}
