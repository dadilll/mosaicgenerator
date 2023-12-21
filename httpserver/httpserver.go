package httpserver

import (
	"fmt"
	"net/http"

	image "github.com/dadil/mosaicgenerator/pkg"
)

func ServeForm(w http.ResponseWriter) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Mosaic Generator</title>
		</head>
		<body>
			<form method="post" action="/" enctype="multipart/form-data">
				<label>Select an image to upload:</label>
				<input type="file" name="image" accept="image/*">
				<br>
				<input type="submit" value="Generate Mosaic">
			</form>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Save the uploaded image
	savePath := "uploaded_image.jpg"
	err = image.SaveImage(file, savePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create and save the pixelated image
	mosaicImagePath := "generated_mosaic.jpg"
	err = image.GenerateMosaic(savePath, mosaicImagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the generated image in the response
	http.ServeFile(w, r, mosaicImagePath)
}
