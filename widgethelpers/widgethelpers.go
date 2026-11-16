package widgethelpers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Entry(w *widget.Entry, scoll bool, validation bool, size fyne.Size, position fyne.Position, style fyne.TextStyle) {
	w.Move(fyne.NewPos(position.X, position.Y))
	w.Resize(fyne.NewSize(size.Width, size.Height))

	w.Scroll = fyne.ScrollNone
	w.TextStyle = fyne.TextStyle{Bold: true}

	switch validation {
	case true:
		w.AlwaysShowValidationError = true
	case false:
		w.Validator = nil
		w.AlwaysShowValidationError = false
	}

}
