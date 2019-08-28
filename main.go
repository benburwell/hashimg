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
	username := flag.String("username", "", "username to encode as image")
	flag.Parse()
	if strings.TrimSpace(*username) == "" {
		fmt.Println("no username")
		return
	}
	if !usernamePattern.MatchString(*username) {
		fmt.Println("no match")
		return
	}

	seed := hash(*username)
	err, img := generateImage(seed)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	err, str := encodeImage(img)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(str)
}

func hash(username string) [32]byte {
	return sha256.Sum256([]byte(username))
}

func generateImage(hash [32]byte) (error, image.Image) {
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
			on := hash[offset]&0xf0 > 0
			color := colorA
			if on {
				color = colorB
			}
			img.Set(x, y, color)
			offset++
		}
	}
	return nil, img
}

func encodeImage(img image.Image) (error, string) {
	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return err, ""
	}
	return nil, base64.StdEncoding.EncodeToString(b.Bytes())
}
