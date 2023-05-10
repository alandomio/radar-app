package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Struct per il JSON
type ImageData struct {
	Bbox   string `json:"bbox"`
	Date   string `json:"date"`
	Height string `json:"height"`
	Hhmm   string `json:"hhmm"`
	Mode   string `json:"mode"`
	Path   string `json:"path"`
	Valid  string `json:"valid"`
	Width  string `json:"width"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	images := getData()
	// Verifica se ci sono immagini nel JSON
	if len(images) == 0 {
		http.Error(w, "Nessuna immagine trovata nel JSON", http.StatusBadRequest)
		return
	}

	image := images[0]

	// Costruisce l'URL completo per il download dell'immagine
	imageURL := "http://www.meteo.si" + image.Path

	// Scarica l'immagine
	resp, err := http.Get(imageURL)
	if err != nil {
		http.Error(w, "Errore nel download dell'immagine.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Legge il contenuto dell'immagine
	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del contenuto dell'immagine.", http.StatusInternalServerError)
		return
	}

	// Codifica l'immagine in base64
	imgBase64 := base64.StdEncoding.EncodeToString(imgData)

	// Costruisce la pagina HTML
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Immagine</title>
	</head>
	<body>
		<img src="data:image/png;base64,%s">
	</body>
	</html>
`
	html = strings.ReplaceAll(html, "%s", imgBase64)

	// Invia la pagina HTML al client
	fmt.Fprint(w, html)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func getData() []ImageData {
	response, err := http.Get("http://www.meteo.si/uploads/probase/www/nowcast/inca/inca_si0zm_data.json?prod=si0zm")
	if err != nil {
		fmt.Println("Errore nella richiesta HTTP:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Errore nella lettura della risposta:", err)
		return nil
	}

	var data []ImageData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Errore nella decodifica JSON:", err)
		return nil
	}

	return data
}
