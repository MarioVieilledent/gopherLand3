package graphics

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawScreen(window *Window, screen *ebiten.Image) {

	screen.Fill(color.RGBA{0xdd, 0x99, 0x55, 0xff})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(window.scale, window.scale)

	screen.DrawImage(stoneImg, op)
	op.GeoM.Translate(
		window.scale*float64(window.tileSize),
		0,
	)
	screen.DrawImage(woodImg, op)

	ebitenutil.DebugPrint(screen, "Yo")
	ebitenutil.DebugPrint(screen, "\nfullScreen = "+strconv.FormatBool(window.fullScreen))
}
