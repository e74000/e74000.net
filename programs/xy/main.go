package main

import (
	"github.com/charmbracelet/log"
	"github.com/e74000/wshim"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lucasb-eyer/go-colorful"
	"math"
	"math/rand"
)

const (
	XScale = 128
	YScale = 64
)

var (
	temperature = 0.5
	external    = 0.0
	interaction = 0.2
)

type Game struct {
	pixels []byte
	spins  []float64
}

func (g *Game) Layout(_, _ int) (int, int) {
	return XScale, YScale
}

func mod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func fmod(d, m float64) float64 {
	res := math.Mod(d, m)
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func torus(mx, my, x, y int) (int, int) {
	return mod(x, mx), mod(y, my)
}

func (g *Game) Update() error {
	for i := 0; i < XScale*YScale; i++ {
		px := rand.Int() % XScale
		py := rand.Int() % YScale

		f := 0.0

		for _, j := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			tx, ty := torus(XScale, YScale, px+j[0], py+j[1])
			f += math.Sin(g.spins[py*XScale+px] - g.spins[ty*XScale+tx])
		}

		f *= interaction
		f += temperature * (rand.Float64()*2 - 1)
		f += external * math.Sin(g.spins[py*XScale+px])

		g.spins[py*XScale+px] = fmod(g.spins[py*XScale+px]-f, 2*math.Pi)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, 4*XScale*YScale)

		for i := 0; i < XScale*YScale; i++ {
			g.pixels[4*i+3] = 0xff
		}
	}

	for y := 0; y < YScale; y++ {
		for x := 0; x < XScale; x++ {
			i := y*XScale + x

			c := colorful.Hsv(g.spins[i]/(math.Pi)*180, 1, 1)

			g.pixels[4*i+0] = uint8(c.R * 255)
			g.pixels[4*i+1] = uint8(c.G * 255)
			g.pixels[4*i+2] = uint8(c.B * 255)
		}
	}

	screen.WritePixels(g.pixels)
}

func main() {
	wshim.SetLogLevel(log.DebugLevel)

	wshim.Run(_main,
		wshim.FloatSlider("Temperature", 0, 2, 0.001, &temperature),
		wshim.FloatSlider("Interaction Strength", -1, 1, 0.001, &interaction),
		wshim.FloatSlider("External Field", -0.5, 0.5, 0.001, &external),
	)
}

func _main() {
	g := &Game{
		pixels: make([]byte, 4*XScale*YScale),
		spins:  make([]float64, XScale*YScale),
	}

	for i := 0; i < XScale*YScale; i++ {
		g.spins[i] = rand.Float64() * 2 * math.Pi
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
