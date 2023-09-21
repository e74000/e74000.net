package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	keys       []ebiten.Key
	mx, my     int
	ml, mr, mm bool
	wx, wy     float64
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.ml = true
	} else {
		g.ml = false
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.mr = true
	} else {
		g.mr = false
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		g.mm = true
	} else {
		g.mm = false
	}

	g.wx, g.wy = ebiten.Wheel()

	g.mx, g.my = ebiten.CursorPosition()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			`{
	"keys": %v,
	"mouse_x": %d,
	"mouse_y": %d,
	"mouse_left_pressed":   %t,
	"mouse_right_pressed":  %t,
	"mouse_middle_pressed": %t,
	"wheel_x": %f,
	"wheel_y": %f
}`,
			g.keys,
			g.mx,
			g.my,
			g.ml,
			g.mr,
			g.mm,
			g.wx,
			g.wy,
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	g := &Game{
		keys: make([]ebiten.Key, 0),
	}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("input test")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
