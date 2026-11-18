package goutilities

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/BadRobotXIII/geomath/geomath"
)

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
	fmt.Print("Focus gained\n")
	e.Entry.FocusGained()
	e.TypedShortcut(&fyne.ShortcutSelectAll{})
}

// FocusLost formats the number back to %.3f when the user clicks away
func (e *FloatEntry) FocusLost() {
	e.Entry.FocusLost()
	e.syncFromData()
}

func main() {

	//Set and define world origin
	Origin := geomath.Point{Xw: 0.0, Yw: 0.0, Zw: 0.0, Xp: 100, Yp: 700}

	p1X, p1Y := binding.NewFloat(), binding.NewFloat()
	p1X.Set(Origin.Xp)
	p1Y.Set(Origin.Yp)
	p2X, p2Y := binding.NewFloat(), binding.NewFloat()
	line := canvas.NewLine(color.RGBA{R: 255, G: 0, B: 0, A: 100})
	line.StrokeWidth = 5

	//Point l set line XY from entry
	p1X.AddListener(binding.NewDataListener(func() {
		fmt.Print("P1 X istener\n")
		valX, _ := p1X.Get()
		line.Position1.X = float32(valX)
		line.Refresh()

	}))
	p1Y.AddListener(binding.NewDataListener(func() {
		valY, _ := p1Y.Get()
		line.Position1.Y = float32(valY)
		line.Refresh()

	}))
	//Point 2 set line XY from entrys
	p2X.AddListener(binding.NewDataListener(func() {
		fmt.Print("P2 X Listener\n")
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

	pOrigin := canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 255, A: 100})
	pOrigin.Move(fyne.NewPos(float32(Origin.Xp), float32(Origin.Yp)))
	pOrigin.FillColor = color.RGBA{R: 0, G: 0, B: 255, A: 100}
	pOrigin.Resize(fyne.NewSize(10, 10))
}
