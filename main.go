package main

import (
	"flag"
	"log"
	"math"
	"net/http"
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

	http.HandleFunc("/", asciifyHandler)

	log.Println(banner)
	log.Println("Listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func asciifyHandler(w http.ResponseWriter, r *http.Request) {
	var rhint int
	rhints, ok := r.URL.Query()["rhint"]
	if ok && len(rhints) >= 1 {
		rhint, _ = strconv.Atoi(rhints[0])
	} else {
		rhint = 5
	}

	var maxw int
	maxws, ok := r.URL.Query()["maxw"]
	if ok && len(maxws) >= 1 {
		maxw, _ = strconv.Atoi(maxws[0])
	} else {
		maxw = 200
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

	log.Printf("maxw: [%d] rhint: [%d] width: [%d]\n", maxw, rhint, width)

	img := resize.Resize(width, 0, imgOriginal, resize.Lanczos3)

	// convert to grayscale and then asciify
	ascii := asciify(grayscale(img))

	// write back
	if _, err := w.Write([]byte(ascii)); err != nil {
		log.Println("unable to write image.")
		http.Error(w, "unable to write image.", 500)
		return
	}

}
