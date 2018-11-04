package main

import (
	"image"
	"image/color"
	"log"
	"strings"
)

// convert to grayscale
func grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	log.Println(bounds)

	imgSet := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			imgSet.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	return imgSet
}

// conver to ascii using the charmap defined
func asciify(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var ascii = make([]string, height, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// in golang, image library premultiplies the values by 257 to prevent overflow
			// on operations on the pixels. hence the values are stored in 32 bit integers.
			// this is needed to bring them back to [0-255] range.
			//
			// read more here: https://blog.golang.org/go-image-package
			// + Note that the R field of an RGBA is an 8-bit alpha-premultiplied color in the range [0, 255].
			// + RGBA satisfies the Color interface by multiplying that value by 0x101 to generate a 16-bit
			// + alpha-premultiplied color in the range [0, 65535].
			//

			r = r / 257 // 0x101
			g = g / 257 // 0x101
			b = b / 257 // 0x101

			avg := int((r + g + b) / 3)

			ascii[y] += charmap[avg]
		}
	}

	return strings.Join(ascii, "\n")
}
