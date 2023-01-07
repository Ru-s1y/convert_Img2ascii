package main

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"golang.org/x/image/draw"
)

const (
	asciiChars = "MNHQ$OC?7>!:-;. "
	imageWidth  = 100
	imageHeight = 100
)

func main() {
	// Open the input image file
	input := os.Args[1]
	fmt.Println(input)
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the image
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// Scale the image to the desired width and height
	img = scaleImage(img)

	// Convert the image to grayscale
	img = toGrayscale(img)

	// Build the ASCII representation of the image
	var ascii strings.Builder
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			ascii.WriteByte(getAsciiChar(img.At(x, y)))
		}
		ascii.WriteByte('\n')
	}

	// Print the ASCII representation of the image
	textBuffer := []byte(ascii.String())
	ioutil.WriteFile("./output.txt", textBuffer, 0644)
	println(ascii.String())
}

// scaleImage scales an image to the specified width and height
func scaleImage(img image.Image) image.Image {
	imgDst := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return imgDst
}

// toGrayscale converts an image to grayscale
func toGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y;  y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}
	return dest
}

// getAsciiChar returns the ASCII character that is closest in intensity to the specified color
func getAsciiChar(c color.Color) byte {
	// Convert input color to grayscale
	r, g, b, _ := c.RGBA()
	intensity := (r*299 + g*587 + b*114) / 1000

	// Divide intensity range into 16 equal parts
	numChars := len(asciiChars)
	step := 255 / numChars

	// Return ASCII character based on intensity
	for i := 0; i < numChars; i++ {
		if intensity < uint32(step * (i + 1)) {
			return asciiChars[i]
		}
	}
	return ' '
}
