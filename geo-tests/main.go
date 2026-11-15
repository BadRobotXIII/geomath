package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/BadRobotXIII/geomath/geomath"
)

type EntryLayout struct {
	SizeW  float32
	SizeH  float32
	LocL   float32
	LocT   float32
	SpaceW float32
	SpaceH float32
}

type FloatEntry struct {
	widget.Entry
	data   binding.Float
	format string
}

func NewFloatEntry(data binding.Float, format string) *FloatEntry {
	e := &FloatEntry{
		data:   data,
		format: format,
	}
	e.ExtendBaseWidget(e)

	// Set initial formatted value
	e.syncFromData()
	return e
}

func (e *FloatEntry) syncFromData() {
	if val, err := e.data.Get(); err == nil {
		e.SetText(fmt.Sprintf(e.format, val))
	}
}

// FocusGained selects all text when the widget receives focus
func (e *FloatEntry) FocusGained() {
	e.Entry.FocusGained()
	e.TypedShortcut(&fyne.ShortcutSelectAll{})
}
func main() {

	entLayout := EntryLayout{SizeW: 120, SizeH: 35, LocL: 10, LocT: 20, SpaceW: 20, SpaceH: 30}

	//Create app window
	a := app.New()
	w := a.NewWindow("Geo World")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1000.0, 800.0))

	p1X, p1Y := binding.NewFloat(), binding.NewFloat()
	p1X.Set(100.0)
	p1Y.Set(700.0)
	p2X, p2Y := binding.NewFloat(), binding.NewFloat()
	p2X.Set(700.0)
	p2Y.Set(100.0)

	lblP1X, lblP1Y := widget.NewLabel("P1 X:"), widget.NewLabel("P1 Y:")
	entP1X, entP1Y := widget.NewEntryWithData(binding.FloatToStringWithFormat(p1X, "%.3f")),
		widget.NewEntryWithData(binding.FloatToStringWithFormat(p1Y, "%.3f"))
	lblP2X, lblP2Y := widget.NewLabel("P2 X:"), widget.NewLabel("P2 Y:")
	entP2X, entP2Y := widget.NewEntryWithData(binding.FloatToStringWithFormat(p2X, "%.3f")),
		widget.NewEntryWithData(binding.FloatToStringWithFormat(p2Y, "%.3f"))

	//P1 Locate and size widgets
	entP1X.Move(fyne.NewPos(entLayout.LocL, entLayout.LocT))
	entP1X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP1X.Move(fyne.NewPos(entP1X.Position().X, entP1X.Position().Y-entLayout.SpaceH))
	lblP1X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP1X.Validator = nil
	entP1X.AlwaysShowValidationError = false
	entP1X.Scroll = fyne.ScrollNone
	entP1X.TextStyle = fyne.TextStyle{Bold: true}

	entP1Y.Move(fyne.NewPos(entP1X.Position().X+entP1X.Size().Width+entLayout.SpaceW, entLayout.LocT))
	entP1Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP1Y.Move(fyne.NewPos(entP1Y.Position().X, entP1Y.Position().Y-entLayout.SpaceH))
	lblP1Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP1Y.Validator = nil

	//P2 Locate and size widgets
	entP2X.Move(fyne.NewPos(entP1Y.Position().X+entP1Y.Size().Width+entLayout.SpaceW, entLayout.LocT))
	entP2X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP2X.Move(fyne.NewPos(entP2X.Position().X, entP2X.Position().Y-entLayout.SpaceH))
	lblP2X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP2X.Validator = nil

	entP2Y.Move(fyne.NewPos(entP2X.Position().X+entP2X.Size().Width+entLayout.SpaceW, entLayout.LocT))
	entP2Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP2Y.Move(fyne.NewPos(entP2Y.Position().X, entP2Y.Position().Y-entLayout.SpaceH))
	lblP2Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP2Y.Validator = nil

	c := geomath.Circle{Diameter: 25.4, Radius: 53.7}
	c.Circumfrence = c.CalcCircum(57)
	lblCircum := widget.NewLabel(strconv.FormatFloat(c.Circumfrence, 'e', 6, 64))
	fmt.Printf("Circumfrence=%s", lblCircum.Text)
	line := canvas.NewLine(color.RGBA{R: 255, G: 0, B: 0, A: 50})
	line.StrokeWidth = 5
	//Point l set line XY from entry
	p1X.AddListener(binding.NewDataListener(func() {
		val, _ := p1X.Get()
		line.Position1.X = float32(val)
		line.Refresh()
	}))
	p1Y.AddListener(binding.NewDataListener(func() {
		val, _ := p1Y.Get()
		line.Position1.Y = float32(val)
		line.Refresh()
	}))
	//Point 2 set line XY from entrys
	p2X.AddListener(binding.NewDataListener(func() {
		val, _ := p2X.Get()
		line.Position2.X = float32(val)
		line.Refresh()
	}))

	p2Y.AddListener(binding.NewDataListener(func() {
		val, _ := p2Y.Get()
		line.Position2.Y = float32(val)
		line.Refresh()
	}))

	valP2X, _ := p2X.Get()
	line.Position2.X = float32(valP2X)
	valP2Y, _ := p2Y.Get()
	line.Position2.Y = float32(valP2Y)
	w.SetContent(container.NewWithoutLayout(lblP1X, entP1X, lblP1Y, entP1Y, lblP2X, entP2X, lblP2Y, entP2Y, line))
	line.Refresh()
	w.ShowAndRun()

}
