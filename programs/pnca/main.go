package main

import (
	"github.com/charmbracelet/log"
	"github.com/e74000/wshim"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"golang.org/x/exp/rand"
)

type v2 struct {
	x, y int
}

func (v v2) add(w v2) v2 {
	return v2{
		x: v.x + w.x,
		y: v.y + w.y,
	}
}

var (
	temp  float64
	damp  float64
	sigma float64
)

func (v v2) inBounds(sx int, sy int) bool {
	return v.x >= 0 && v.x < sx && v.y >= 0 && v.y < sy
}

var offsets = [8]v2{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

type Game struct {
	probabilities []float64
	activations   []float64
	weights       [][8]float64

	sx, sy int

	pixels []byte
	cmap   colorgrad.Gradient
}

func (g *Game) Update() error {

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			i := y*g.sx + x
			iv := v2{x, y}
			act := g.activations[i]

			if act > 0.5 {
				for off := 0; off < 8; off++ {
					jv := iv.add(offsets[off])
					j := jv.y*g.sx + jv.x
					if jv.inBounds(g.sx, g.sy) {
						g.probabilities[j] += g.weights[i][off] * act * sigma
					}
				}
			}

			if rand.Float64() <= temp {
				g.probabilities[i] = 1 - g.probabilities[i]
				g.weights[i] = randWithMag(1)
			}
		}
	}

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			i := y*g.sx + x

			if rand.Float64() <= g.probabilities[i] {
				g.activations[i] = 1
			} else {
				g.activations[i] /= damp
			}

			g.probabilities[i] = 0
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, act := range g.activations {
		col := g.cmap.At(act)
		g.pixels[4*i+0], g.pixels[4*i+1], g.pixels[4*i+2] = col.RGB255()
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.sx, g.sy
}

func (g *Game) Init(sx, sy int) {
	g.sx, g.sy = sx, sy

	g.activations = make([]float64, sx*sy)
	g.probabilities = make([]float64, sx*sy)
	g.weights = make([][8]float64, sx*sy)

	g.pixels = make([]byte, 4*sx*sy)

	for i := 0; i < sx*sy; i++ {
		g.weights[i] = randWithMag(1)
	}

	g.cmap = colorgrad.CubehelixDefault()
}

func randWithMag(sigma float64) [8]float64 {
	out := [8]float64{}

	var sum float64

	for i := 0; i < 8; i++ {
		v := rand.Float64()
		out[i] = v
		sum += v
	}

	if sum == 0 {
		return randWithMag(sigma)
	}

	for i := 0; i < 8; i++ {
		out[i] *= sigma / sum
	}

	return out
}

func main() {
	temp = 0.0001
	damp = 2.01
	sigma = 1

	wshim.SetLogLevel(log.DebugLevel)

	wshim.Run(
		_main,
		wshim.FloatSlider("Temperature", 0, 0.01, 0.0001, &temp),
		wshim.FloatSlider("Damping", 1, 10, 0.001, &damp),
		wshim.FloatSlider("Sigma", 0, 2, 0.001, &sigma),
	)
}

func _main() {
	g := new(Game)
	g.Init(128, 64)

	ebiten.SetTPS(10)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
