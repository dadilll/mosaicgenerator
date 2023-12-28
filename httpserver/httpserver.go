package httpserver

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	image "github.com/dadil/mosaicgenerator/pkg"
)

func ServeForm(w http.ResponseWriter) {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFile)
	filePath := filepath.Join(dir, "form.html")

	htmlContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading HTML file: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(htmlContent))
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	pixelSizeStr := r.FormValue("pixelSize")
	pixelSize, err := strconv.Atoi(pixelSizeStr)
	if err != nil {
		http.Error(w, "Invalid pixel size", http.StatusBadRequest)
		return
	}

	imageType := filepath.Ext(fileHeader.Filename)[1:]

	savePath := "uploaded_image." + imageType
	err = image.SaveImage(file, savePath, imageType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mosaicImagePath := "generated_mosaic." + imageType
	err = image.GenerateMosaic(savePath, mosaicImagePath, pixelSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, mosaicImagePath)
}
