package ch1

import (
	"bytes"
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

var colorPurple = color.RGBA{R: 128, B: 128, A: 255}
var colorOrange = color.RGBA{R: 255, G: 140, A: 255}
var colorYellow = color.RGBA{R: 180, G: 174, B: 55, A: 255}

var palette6 = []color.Color{color.Black, colorPurple, colorOrange, colorYellow}

const (
	blackIndex6 = 0
	greenIndex6 = 1
	blueIndex6  = 2
	redIndex6   = 3
)

func setColorIndex(frameNum int) uint8 {
	colorIndex := uint8(blackIndex6)
	switch {
	case frameNum%9 <= 2:
		colorIndex = greenIndex6
	case frameNum%9 >= 3 && frameNum%9 <= 5:
		colorIndex = redIndex6
	case frameNum%9 >= 6 && frameNum%9 <= 8:
		colorIndex = blueIndex6
	}
	return colorIndex
}

func Ex6(args []string) {
	buf := &bytes.Buffer{}
	lissajous6(buf)
	if err := os.WriteFile("ch1/"+args[0]+".gif", buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
	println("Done, look at ch1 folder!")
}

func lissajous6(out io.Writer) {
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
		img := image.NewPaletted(rect, palette6)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), setColorIndex(i))
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
