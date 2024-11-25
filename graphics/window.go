package graphics

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	_ "image/png" // Ensure the PNG format is supported
)

var img *ebiten.Image

//go:embed resources/brick_2.png
var brickImageData []byte

func init() {
	var err error
	log.Println(len(brickImageData))
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(brickImageData))
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
	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		w.fullScreen = !w.fullScreen
		ebiten.SetFullscreen(w.fullScreen)
	}
	_, dy := ebiten.Wheel()
	w.scale += w.zoomFactor * dy
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
