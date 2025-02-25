package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x55, 0x66, 0xBB, 0xff}}

const (
	blackIndex = 0
	whiteIndex = 1

	randomIndex = 2
)

func main() {
	cycles := 5
	if len(os.Args) > 1 && os.Args[1] == "web" {

		handler := func(w http.ResponseWriter, r *http.Request) {

			for k, v := range r.Form {
				if k == "cycles" {
					cycles, _ = strconv.Atoi(v[0])
				}

			}
			lissajous(w, float64(cycles))
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, float64(cycles))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
