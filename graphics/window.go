package graphics

import (
	"bytes"
	"image"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var textureImage *ebiten.Image

//go:embed resources/textures.png
var textureImageData []byte

var textures map[string]*ebiten.Image

func init() {
	var err error
	textureImage, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(textureImageData))
	if err != nil {
		log.Fatal(err)
	}

	textures = make(map[string]*ebiten.Image)

	spriteRegions := map[string]image.Rectangle{
		"player":   image.Rect(0, 0, 32, 32),
		"enemy":    image.Rect(32, 0, 64, 32),
		"item":     image.Rect(64, 0, 96, 32),
		"platform": image.Rect(0, 32, 64, 64),
	}
	for key, rect := range spriteRegions {
		textures[key] = textureImage.SubImage(rect).(*ebiten.Image)
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
		scale:        4.0,
		zoomFactor:   config.Input.ZoomFactor,
	}

	ebiten.SetWindowSize(window.windowWidth, window.windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Gopher Land 3")
	if err := ebiten.RunGame(&window); err != nil {
		log.Fatal("Error running ebiten game: %w", err)
	}
}
