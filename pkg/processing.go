package image

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

const (
	width     = 640
	height    = 520
	pixelSize = 10
)

func SaveImage(file multipart.File, savePath string) error {
	imageData, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, imageData, nil)
	if err != nil {
		return err
	}

	return nil
}

func GenerateMosaic(inputPath, outputPath string, pixelSize int) error {
	img, err := loadImage(inputPath)
	if err != nil {
		return err
	}

	resized := resize.Resize(width, height, img, resize.NearestNeighbor)

	mosaic := createMosaic(resized, pixelSize)

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, mosaic, nil)
	if err != nil {
		return err
	}

	return nil
}

func createMosaic(img image.Image, pixelSize int) *image.RGBA {
	bounds := img.Bounds()
	mosaic := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x += pixelSize {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += pixelSize {

			avgColor := averageColor(img, x, y, pixelSize)

			draw.Draw(mosaic, image.Rect(x, y, x+pixelSize, y+pixelSize), &image.Uniform{avgColor}, image.Point{}, draw.Src)
		}
	}

	return mosaic
}

func averageColor(img image.Image, x, y, size int) color.RGBA {
	var r, g, b, a uint32

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c := color.NRGBAModel.Convert(img.At(x+i, y+j)).(color.NRGBA)
			r += uint32(c.R)
			g += uint32(c.G)
			b += uint32(c.B)
			a += uint32(c.A)
		}
	}

	pixelCount := uint32(size * size)
	return color.RGBA{uint8(r / pixelCount), uint8(g / pixelCount), uint8(b / pixelCount), uint8(a / pixelCount)}
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
