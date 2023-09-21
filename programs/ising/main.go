package main

import (
	"github.com/charmbracelet/log"
	"github.com/e74000/wshim"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
)

// Set X and Y resolution
const (
	XScale = 200
	YScale = 100
)

// Initialise temperature to critical value
var temperature = 2.26918531421

// Game Type - Stores all data needed in main loop.
type Game struct {
	pixels []byte
	grid   []float64
}

// Update runs each frame
func (g *Game) Update() error {
	// For each position in grid
	for i := 0; i < XScale; i++ {
		for j := 0; j < YScale; j++ {

			// Pick a random position and get value s
			a := rand.Int() % XScale
			b := rand.Int() % YScale
			s := g.grid[b*XScale+a]

			// Get neighbors and calculate cost
			nb := g.grid[b*XScale+FW(a+1, XScale)] + g.grid[b*XScale+FW(a-1, XScale)] + g.grid[FW(b+1, YScale)*XScale+a] + g.grid[FW(b-1, YScale)*XScale+a]
			cost := 2 * s * nb

			// Flip cell if cost is negative, or otherwise randomly
			if cost < 0 {
				s *= -1
			} else if rand.Float64() < math.Exp(-cost/temperature) {
				s *= -1
			}

			// Update the grid
			g.grid[b*XScale+a] = s
		}
	}

	return nil
}

// Layout outputs the resolution of the program
func (g *Game) Layout(_, _ int) (int, int) {
	return XScale, YScale
}

// Draw is run each frame, to update the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Set pixels depending on state of grid
	for i, f := range g.grid {
		if f == 1 {
			g.pixels[4*i+0] = 0xff
			g.pixels[4*i+1] = 0xff
			g.pixels[4*i+2] = 0xff
		} else {
			g.pixels[4*i+0] = 0x05
			g.pixels[4*i+1] = 0x05
			g.pixels[4*i+2] = 0x05
		}
	}

	screen.WritePixels(g.pixels)
}

func FW(a, b int) int {
	if a < 0 {
		return a + b
	} else if a >= b {
		return a - b
	}
	return a
}

func main() {
	wshim.SetLogLevel(log.DebugLevel)

	// Initialise program with one slider
	wshim.Run(_main, wshim.FloatSlider(
		"Temperature", 0.001, 10, 0.001, &temperature,
	))
}

// The main program
func _main() {
	// Initialise game
	g := &Game{
		pixels: make([]byte, 4*XScale*YScale),
		grid:   make([]float64, XScale*YScale),
	}

	// Initialise grid randomly
	for i := range g.grid {
		if rand.Int()%2 == 1 {
			g.grid[i] = 1
		} else {
			g.grid[i] = -1
		}
	}

	// Run simulation
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
