package ui

import (
	"ztm/pixl/apptype"
	"ztm/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
