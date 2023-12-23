package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

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
				<label>Pixel Size:</label>
				<input type="number" name="pixelSize" value="10" min="1"> <!-- Added input for pixel size -->
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

	// Extract pixel size from the form
	pixelSizeStr := r.FormValue("pixelSize")
	pixelSize, err := strconv.Atoi(pixelSizeStr)
	if err != nil {
		http.Error(w, "Invalid pixel size", http.StatusBadRequest)
		return
	}

	// Save the uploaded image
	savePath := "uploaded_image.jpg"
	err = image.SaveImage(file, savePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create and save the pixelated image with the specified pixel size
	mosaicImagePath := "generated_mosaic.jpg"
	err = image.GenerateMosaic(savePath, mosaicImagePath, pixelSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the generated image in the response
	http.ServeFile(w, r, mosaicImagePath)
}
