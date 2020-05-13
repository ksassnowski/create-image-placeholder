package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stdout, "usage: preview [path-to-image]")
		os.Exit(0)
	}

	path := os.Args[1]
	str, err := generate(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %q\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, str)
	os.Exit(0)
}

func generate(path string) (string, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return "", err
	}

	img = imaging.Resize(img, 32, 0, imaging.NearestNeighbor)
	dst := imaging.Blur(img, 5)

	var buff bytes.Buffer

	jpeg.Encode(&buff, dst, &jpeg.Options{
		Quality: 100,
	})

	return base64.StdEncoding.EncodeToString(buff.Bytes()), nil
}
