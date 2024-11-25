package graphics

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var stoneImg *ebiten.Image
var woodImg *ebiten.Image

//go:embed resources/stone.png
var stoneImageData []byte

//go:embed resources/wood.png
var woodImageData []byte

func init() {
	var err error
	stoneImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(stoneImageData))
	if err != nil {
		log.Fatal(err)
	}

	woodImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(woodImageData))
	if err != nil {
		log.Fatal(err)
	}
}

type Window struct {
	fullScreen   bool
	windowWidth  int
	windowHeight int
	tileSize     int
	scale        float64
	zoomFactor   float64
}

func (w *Window) Update() error {
	manageInput(w)
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	DrawScreen(w, screen)
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func StartWindow() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config file: %w", err)
	}
	window := Window{
		fullScreen:   false,
		windowWidth:  config.Window.DefaultWindowWidth,
		windowHeight: config.Window.DefaultWindowHeight,
		tileSize:     config.Graphics.TileSize,
		scale:        1.0 / 4.0,
		zoomFactor:   config.Input.ZoomFactor,
	}

	ebiten.SetWindowSize(window.windowWidth, window.windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Gopher Land 3")
	if err := ebiten.RunGame(&window); err != nil {
		log.Fatal("Error running ebiten game: %w", err)
	}
}
