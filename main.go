package main

import (
	"bytes"
	"flag"
	"image/jpeg"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/nfnt/resize"
	"gopkg.in/tokopedia/logging.v1"
)

var banner = `
    /\    / ____|/ ____|_   _|_   _|/ _|      
   /  \  | (___ | |      | |   | | | |_ _   _ 
  / /\ \  \___ \| |      | |   | | |  _| | | |
 / ____ \ ____) | |____ _| |_ _| |_| | | |_| |
/_/    \_\_____/ \_____|_____|_____|_|  \__, |
                                         __/ |
                                        |___/  v0.01	
`

func main() {
	flag.Parse()
	logging.LogInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	port = ":" + port

	http.HandleFunc("/", asciifyHandler)

	log.Println(banner)
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func asciifyHandler(w http.ResponseWriter, r *http.Request) {
	var rhint int
	rhints, ok := r.URL.Query()["rhint"]
	if ok && len(rhints) >= 1 {
		rhint, _ = strconv.Atoi(rhints[0])
	} else {
		rhint = 2
	}

	var scale int
	scales, ok := r.URL.Query()["scale"]
	if ok && len(scales) >= 1 {
		scale, _ = strconv.Atoi(scales[0])
	} else {
		scale = 10
	}

	var maxw int
	maxws, ok := r.URL.Query()["maxw"]
	if ok && len(maxws) >= 1 {
		maxw, _ = strconv.Atoi(maxws[0])
	} else {
		maxw = 175
	}

	var algo int
	algos, ok := r.URL.Query()["algo"]
	if ok && len(algos) >= 1 {
		algo, _ = strconv.Atoi(algos[0])
	} else {
		algo = 2
	}

	// algo has to be 1 or 2, this basically determines the ascii charset to be used
	if (algo != 1) && (algo != 2) {
		algo = 2
	}

	keys, ok := r.URL.Query()["image_url"]
	if !ok || len(keys[0]) < 1 {
		log.Println("err `image_url` parameter is missing")
		http.Error(w, "err `image_url` parameter is missing", 400)
		return
	}

	// read the image from the input URL
	imgOriginal, err := read(keys[0])
	if err != nil {
		log.Printf("err problem reading the image: %s\n", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	// maximum width should be 200 for rendering well on the browser.
	// two tweaks possible using params:
	// 		rhint = how much downsize should you do. more the value small the width.
	// 		maxw = override the max width from 200 to whatever you want.
	width := uint(math.Min(float64(maxw), float64(imgOriginal.Bounds().Max.X)/float64(rhint)))

	//log.Printf("maxw: [%d] rhint: [%d] width: [%d]\n", maxw, rhint, width)

	img := resize.Resize(width, 0, imgOriginal, resize.Lanczos3)

	// convert to grayscale and then asciify
	asciiImage := asciify(grayscale(img), algo, scale)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, asciiImage, nil); err != nil {
		log.Println("unable to encode image.")
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("info input image url = [%s]\n", keys[0])

	// // write back
	// if _, err := w.Write([]byte(ascii)); err != nil {
	// 	log.Println("unable to write image.")
	// 	http.Error(w, "unable to write image.", 500)
	// 	return
	// }

}
