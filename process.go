package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"

	"log"

	"github.com/gobuffalo/packr"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	//fontfile = flag.String("fontfile", "../../testdata/luxisr.ttf", "filename of the ttf font")
	hinting = flag.String("hinting", "none", "none | full")
	size    = flag.Float64("size", 12, "font size in points")
	spacing = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb    = flag.Bool("whiteonblack", false, "white text on a black background")
)

// convert to grayscale
func grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	//log.Println(bounds)

	imgSet := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			imgSet.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	return imgSet
}

// conver to ascii using the charmap defined
func asciify(img image.Image, algo int, scale int) image.Image {
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

			if algo == 1 {
				ascii[y] += charmap[avg]
			} else {
				ascii[y] += charmap2[avg]
			}
		}
	}

	return rasterize(ascii, width, height, scale)
}

func rasterize(ascii []string, width, height, scale int) image.Image {
	//log.Println(width, height)
	//aspect := float64(width) / float64(height)
	//width = 800
	//height = int((1.0 / aspect) * float64(width))
	width = scale * width
	height = scale * height
	//log.Println(width, height)

	box := packr.NewBox("./fonts")
	fontBytes, err := box.FindString("monospace.ttf")
	if err != nil {
		log.Println("err", "finding the font file", err)
		return nil
	}

	f, err := truetype.Parse([]byte(fontBytes))
	if err != nil {
		log.Println(err)
		return nil
	}

	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	h := font.HintingNone
	switch *hinting {
	case "full":
		h = font.HintingFull
	}
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    *size,
			DPI:     *dpi,
			Hinting: h,
		}),
	}

	y := 10
	for _, s := range ascii {
		d.Dot = fixed.P(0, y)
		d.DrawString(s)
		y += 10
	}

	var img image.Image = rgba
	return img

	// // Save that RGBA image to disk.
	// log.Println("saving the file")
	// outFile, err := os.Create("/tmp/out.png")
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
	// defer outFile.Close()
	// b := bufio.NewWriter(outFile)
	// err = jpeg.Encode(b, rgba)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
	// err = b.Flush()
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
}
