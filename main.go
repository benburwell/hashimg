package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"regexp"
	"strings"
)

var usernamePattern = regexp.MustCompile("[a-zA-Z][a-zA-Z0-9]+")

func main() {
	username := flag.String("username", "", "username to encode as image (must contain only letters a-zA-Z0-9 and start with a letter)")
	flag.Parse()
	if strings.TrimSpace(*username) == "" {
		flag.Usage()
		return
	}

	if !usernamePattern.MatchString(*username) {
		flag.Usage()
		return
	}

	seed := hash(*username)
	img := makeBigger(repeat(generateImage(seed)))
	str, err := encodeImage(img)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(str)
}

func hash(username string) [32]byte {
	return sha256.Sum256([]byte(username))
}

func generateImage(hash [32]byte) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 5, 5))
	colorA := color.NRGBA{
		R: hash[0],
		G: hash[1],
		B: hash[2],
		A: 255,
	}
	colorB := color.NRGBA{
		R: hash[3],
		G: hash[4],
		B: hash[5],
		A: 255,
	}
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

func encodeImage(img image.Image) (string, error) {
	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func repeat(img image.Image) image.Image {
	newImg := image.NewNRGBA(image.Rect(0, 0, 10, 10))
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			newImg.Set(x, y, img.At(x, y))
			newImg.Set(9-x, y, img.At(x, y))
			newImg.Set(x, 9-y, img.At(x, y))
			newImg.Set(9-x, 9-y, img.At(x, y))
		}
	}
	return newImg
}

func makeBigger(img image.Image) image.Image {
	newImg := image.NewNRGBA(image.Rect(0, 0, 200, 200))
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for i := 0; i < 20; i++ {
				for j := 0; j < 20; j++ {
					newImg.Set(x*20+i, y*20+j, img.At(x, y))
				}
			}
		}
	}
	return newImg
}
