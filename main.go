package main

import (
	"embed"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game object used by ebiten
type game struct {
	ui *ebitenui.UI
}

var data embed.FS
var isStarted bool

func main() {

	ui := ebitenui.UI{
		Container: Container(),
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gopher Run")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game{
		ui: &ui,
	}

	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

// Update implements Ebiten's Update method.
func (g *game) Update() error {
	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *game) Draw(screen *ebiten.Image) {
	if !isStarted {
		g.ui.Draw(screen)
	} else {
		screen.DrawImage(img, nil)
	}
}
