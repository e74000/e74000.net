package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	g := &Game{
		keys: make([]ebiten.Key, 0),
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type Game struct {
	keys        []ebiten.Key
	mouseLeft   bool
	mouseRight  bool
	mouseMiddle bool
	mouseX      int
	mouseY      int
	wheelX      float64
	wheelY      float64
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys)
	g.mouseLeft = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	g.mouseRight = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	g.mouseMiddle = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)

	g.mouseX, g.mouseY = ebiten.CursorPosition()
	g.wheelX, g.wheelY = ebiten.Wheel()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"KEYS: %v\nMOUSE LEFT: %t\nMOUSE RIGHT: %t\nMOUSE MIDDLE: %t\nMOUSE X: %d\nMOUSE Y: %d\nWHEEL X: %f\nWHEEL Y: %f",
		g.keys,
		g.mouseLeft,
		g.mouseRight,
		g.mouseMiddle,
		g.mouseX,
		g.mouseY,
		g.wheelX,
		g.wheelY,
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
