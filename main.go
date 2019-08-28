package main

import (
	// Bag 4: bytes, crypto, encoding, image, regexp
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"regexp"

	// Pantry: flag
	"flag"
	"fmt"
)

var usernamePattern = regexp.MustCompile("[a-zA-Z][a-zA-Z0-9]*")

func main() {
	username := flag.String("username", "", "username to encode as image (must contain only letters a-zA-Z0-9 and start with a letter)")
	mult := flag.Int("mult", 20, "the multiplier to apply to the 10x10 generated image (must be >= 1)")
	flag.Parse()

	if !usernamePattern.MatchString(*username) {
		flag.Usage()
		return
	}

	if *mult < 1 {
		flag.Usage()
		return
	}

	hash := sha256.Sum256([]byte(*username))
	img := makeBigger(tesselate(generateImage(hash)), *mult)
	str, err := encodePNG(img)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(str)
}

func newColor(b []byte) color.NRGBA {
	if len(b) != 3 {
		panic("cannot make color without exactly 3 bytes")
	}
	return color.NRGBA{
		R: b[0],
		G: b[1],
		B: b[2],
		A: 255,
	}
}

func generateImage(hash [32]byte) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 5, 5))
	colorA := newColor(hash[0:3])
	colorB := newColor(hash[3:6])
	offset := 6
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			on := hash[offset] > 127
			color := colorA
			if on {
				color = colorB
			}
			img.Set(x, y, color)
			offset++
		}
	}
	return img
}

func encodePNG(img *image.NRGBA) (string, error) {
	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func tesselate(img *image.NRGBA) *image.NRGBA {
	newImg := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Max.X*2, img.Bounds().Max.Y*2))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			x1 := img.Bounds().Max.X*2 - 1
			y1 := img.Bounds().Max.Y*2 - 1
			newImg.Set(x, y, img.At(x, y))
			newImg.Set(x1-x, y, img.At(x, y))
			newImg.Set(x, y1-y, img.At(x, y))
			newImg.Set(x1-x, y1-y, img.At(x, y))
		}
	}
	return newImg
}

func makeBigger(img *image.NRGBA, mult int) *image.NRGBA {
	newImg := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Max.X*mult, img.Bounds().Max.Y*mult))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			for i := 0; i < mult; i++ {
				for j := 0; j < mult; j++ {
					newImg.Set(x*mult+i, y*mult+j, img.At(x, y))
				}
			}
		}
	}
	return newImg
}
