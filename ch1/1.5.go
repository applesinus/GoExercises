package ch1

import (
	"bytes"
	"fmt"
	"image"
	color "image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var colorGreen = color.RGBA{G: 222, A: 255}
var palette = []color.Color{color.Black, colorGreen}

const (
	blackIndex = 0
	greenIndex = 1
)

func Ex5(args []string) {
	buf := &bytes.Buffer{}
	lissajous(buf)
	if err := os.WriteFile("ch1/"+args[0]+".gif", buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("Done, look at ch1 folder!")
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
