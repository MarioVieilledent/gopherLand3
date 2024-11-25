package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func manageInput(window *Window) {

	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		window.fullScreen = !window.fullScreen
		ebiten.SetFullscreen(window.fullScreen)
	}
	_, dy := ebiten.Wheel()
	window.scale += window.zoomFactor * dy
}
