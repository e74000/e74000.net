package main

import (
	"github.com/charmbracelet/log"
	"github.com/e74000/wshim"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"math"
	"math/rand"
	"time"
)

const (
	XScale = 200
	YScale = 100
)

var (
	c = 0.5
)

var (
	cmap   = colorgrad.RdBu()
	px, py int
)

const (
	bCondWall = iota
	bCondFoll
	bCondAnti
)

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

type Game struct {
	pixels []byte

	next []float64
	curr []float64
	last []float64

	wall []bool

	bCond int

	t time.Time
}

func (g *Game) Update() error {
	for y := 0; y < YScale; y++ {
		for x := 0; x < XScale; x++ {
			i := y*XScale + x

			ic, il := g.curr[i], g.last[i]
			u, d, l, r := g.getNeighborValues(x, y)

			g.next[i] = 2*ic - il +
				math.Pow(c, 2)*
					(l+u+r+d-4*ic)
		}
	}

	copy(g.last, g.curr)
	copy(g.curr, g.next)

	i := (py)*XScale + px

	g.curr[i] = 20 * math.Cos(4*time.Since(g.t).Seconds())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < YScale; y++ {
		for x := 0; x < XScale; x++ {
			i := y*XScale + x

			if g.wall[i] {
				g.pixels[4*i+0] = 0x05
				g.pixels[4*i+1] = 0x05
				g.pixels[4*i+2] = 0x05
				continue
			}

			col := cmap.At(sigmoid(g.curr[i]))

			g.pixels[4*i+0],
				g.pixels[4*i+1],
				g.pixels[4*i+2] = col.RGB255()

		}
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return XScale, YScale
}

func newGame() *Game {
	return &Game{
		pixels: make([]byte, 4*XScale*YScale),
		next:   make([]float64, XScale*YScale),
		curr:   make([]float64, XScale*YScale),
		last:   make([]float64, XScale*YScale),
		wall:   make([]bool, XScale*YScale),
		bCond:  bCondWall,
		t:      time.Now(),
	}
}

// My least favourite function
func (g *Game) getNeighborValues(x, y int) (float64, float64, float64, float64) {
	ui, di, li, ri := (y-1)*XScale+x, (y+1)*XScale+x, y*XScale+(x-1), y*XScale+(x+1)

	var u, d, l, r float64

	var bCond float64

	switch g.bCond {
	case bCondWall:
		bCond = 0.0
	case bCondFoll:
		bCond = g.curr[y*XScale+x]
	case bCondAnti:
		bCond = -g.curr[y*XScale+x]
	}

	if x > 0 {
		if !g.wall[li] {
			l = g.curr[li]
		} else {
			l = bCond
		}
	} else {
		l = bCond
	}

	if x < XScale-1 {
		if !g.wall[ri] {
			r = g.curr[ri]
		} else {
			r = bCond
		}
	} else {
		r = bCond
	}

	if y > 0 {
		if !g.wall[ui] {
			u = g.curr[ui]
		} else {
			u = bCond
		}
	} else {
		u = bCond
	}

	if y < YScale-1 {
		if !g.wall[di] {
			d = g.curr[di]
		} else {
			d = bCond
		}
	} else {
		d = bCond
	}

	return u, d, l, r
}

func main() {
	px = rand.Int() % XScale
	py = rand.Int() % YScale

	wshim.SetLogLevel(log.DebugLevel)

	wshim.Run(_main)
}

func _main() {
	g := newGame()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
