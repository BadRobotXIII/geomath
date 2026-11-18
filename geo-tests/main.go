package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/BadRobotXIII/geomath/geomath"
)

type ObjectLayout struct {
	SizeW  float32
	SizeH  float32
	LocL   float32
	LocT   float32
	SpaceW float32
	SpaceH float32
}

type Vec3 struct{ X, Y, Z float64 }

func main() {

	entLayout := ObjectLayout{SizeW: 120, SizeH: 35, LocL: 10, LocT: 20, SpaceW: 20, SpaceH: 30}
	//Start keyboard listener
	// err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	// 	if key.Code == keys.CtrlC || key.Code == keys.Escape {
	// 		fmt.Println("Exiting...")
	// 		return true, nil // Stop listening
	// 	}
	// 	return false, nil
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//Create app window
	a := app.New()
	w := a.NewWindow("Geo World")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1000.0, 800.0))

	//Set and define world origin
	Origin := geomath.Point{Xw: 0.0, Yw: 0.0, Zw: 0.0, Xp: 100, Yp: 700}

	p1X, p1Y := binding.NewFloat(), binding.NewFloat()
	p1X.Set(Origin.Xp)
	p1Y.Set(Origin.Yp)
	p2X, p2Y := binding.NewFloat(), binding.NewFloat()
	p2X.Set(700.0)
	p2Y.Set(100.0)

	//P1 entry widget and label init
	lblP1X, lblP1Y := widget.NewLabel("P1 X:"), widget.NewLabel("P1 Y:")
	entP1X, entP1Y := widget.NewEntryWithData(binding.FloatToStringWithFormat(p1X, "%.3f")),
		widget.NewEntryWithData(binding.FloatToStringWithFormat(p1Y, "%.3f"))

	//P2 entry widget and label init
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
	entP1Y.AlwaysShowValidationError = false
	entP1Y.Scroll = fyne.ScrollNone
	entP1Y.AlwaysShowValidationError = false
	entP1Y.TextStyle = fyne.TextStyle{Bold: true}

	//P2 Locate and size widgets
	entP2X.Move(fyne.NewPos(entP1Y.Position().X+entP1Y.Size().Width+entLayout.SpaceW, entLayout.LocT))
	entP2X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP2X.Move(fyne.NewPos(entP2X.Position().X, entP2X.Position().Y-entLayout.SpaceH))
	lblP2X.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP2X.Validator = nil
	entP2X.AlwaysShowValidationError = false
	entP2X.Scroll = fyne.ScrollNone
	entP2X.AlwaysShowValidationError = false
	entP2X.TextStyle = fyne.TextStyle{Bold: true}

	entP2Y.Move(fyne.NewPos(entP2X.Position().X+entP2X.Size().Width+entLayout.SpaceW, entLayout.LocT))
	entP2Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	lblP2Y.Move(fyne.NewPos(entP2Y.Position().X, entP2Y.Position().Y-entLayout.SpaceH))
	lblP2Y.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	entP2Y.Validator = nil
	entP2Y.AlwaysShowValidationError = false
	entP2Y.Scroll = fyne.ScrollNone
	entP2Y.AlwaysShowValidationError = false
	entP2Y.TextStyle = fyne.TextStyle{Bold: true}

	slideX := widget.NewSlider(0.0, 360.0)
	slideX.Step = 1.0
	slideX.Move(fyne.NewPos(lblP2Y.Position().X+lblP2Y.Size().Width+entLayout.SpaceW, entLayout.LocT))
	slideX.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))

	lblSlideX := widget.NewLabel("Angle X:")
	lblSlideX.Move(fyne.NewPos(slideX.Position().X, slideX.Position().Y-entLayout.SpaceH))
	lblSlideX.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))

	slideY := widget.NewSlider(0.0, 360.0)
	slideY.Step = 1.0
	slideY.Move(fyne.NewPos(slideX.Position().X+slideX.Size().Width+entLayout.SpaceW, entLayout.LocT))
	slideY.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))

	lblSlideY := widget.NewLabel("Angle Y:")
	lblSlideY.Move(fyne.NewPos(slideY.Position().X, slideY.Position().Y-entLayout.SpaceH))
	lblSlideY.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))

	//Point l set line XY from entry
	p1X.AddListener(binding.NewDataListener(func() {
		val, _ := p1X.Get()
		fmt.Printf("%f\n", val)
	}))
	p1Y.AddListener(binding.NewDataListener(func() {
		val, _ := p1Y.Get()
		fmt.Printf("%f\n", val)
	}))
	//Point 2 set line XY from entrys
	p2X.AddListener(binding.NewDataListener(func() {
		val, _ := p2X.Get()
		fmt.Printf("%f\n", val)
	}))

	p2Y.AddListener(binding.NewDataListener(func() {
		val, _ := p2Y.Get()
		fmt.Printf("%f\n", val)
	}))

	width, height := 1000, 800
	buffer := image.NewRGBA(image.Rect(0, 0, width, height))

	btnClearBuff := widget.NewButton("Clear Buffer", func() { clearBuffer(buffer) })
	btnClearBuff.Move(fyne.NewPos(slideY.Position().X+slideY.Size().Width+entLayout.SpaceW, entLayout.LocT))
	btnClearBuff.Resize(fyne.NewSize(entLayout.SizeW, entLayout.SizeH))
	// Display the raw pixel buffer using a Fyne canvas.Image
	img := canvas.NewImageFromImage(buffer)
	img.FillMode = canvas.ImageFillOriginal
	contWidgets := container.NewWithoutLayout(lblP1X, entP1X, lblP1Y, entP1Y, lblP2X, entP2X, lblP2Y, entP2Y, lblSlideX, slideX, lblSlideY, slideY, btnClearBuff)
	contWidgets.Refresh()
	content := container.NewBorder(contWidgets, nil, nil, nil, img)
	w.SetContent(content)
	w.Resize(fyne.NewSize(float32(width), float32(height)))

	angleLastX, angleLastY := 0.0, 0.0
	var pxTrail [8][8][2]int
	angleX := 0.0
	angleY := 0.0
	captureVerticies := false
	// Simple cube vertices
	cube := []Vec3{
		{-1, -1, -1}, {1, -1, -1}, {1, 1, -1}, {-1, 1, -1},
		{-1, -1, 1}, {1, -1, 1}, {1, 1, 1}, {-1, 1, 1},
	}
	i := 0
	// Render Loop
	go func() {

		angleX += 0.1

		for range time.Tick(time.Second / 30) {
			// 1. Clear pixel buffer (fast black fill)
			// clear(buffer.Pix)

			// 2. Rotate & Project 3D -> 2D
			angleX = slideX.Value * math.Pi / 180
			angleY = slideY.Value * math.Pi / 180

			cosX, sinX := math.Cos(angleX), math.Sin(angleX)
			cosY, sinY := math.Cos(angleY), math.Sin(angleY)

			for vi, v := range cube {
				// Step A: X-axis rotation (modifies Y and Z, leaves X alone)
				y1 := v.Y*cosX - v.Z*sinX
				z1 := v.Y*sinX + v.Z*cosX

				// Step B: Y-axis rotation using z1 from Step A (modifies X and Z, leaves y1 alone)
				rx := v.X*cosY + z1*sinY
				rz := -v.X*sinY + z1*cosY

				// Step C: Push forward away from the camera after rotating
				rz += 4.0

				// Step D: Perspective projection using rx, y1, and rz
				fov := 300.0
				px := int((rx/rz)*fov) + width/2
				py := int((y1/rz)*fov) + height/2
				if angleX != angleLastX || angleY != angleLastY || captureVerticies == true {
					captureVerticies = true
					angleLastX = angleX
					angleLastY = angleY
					fmt.Printf("Moving Angle X: %f Angle Y: %f\n", angleX, angleY)
					if vi < 8 {
						if vi < 8 {
							pxTrail[i][vi][0] = px
							pxTrail[i][vi][1] = py

						} else {
							captureVerticies = false

						}
					}
					if vi == 7 {
						i++
						captureVerticies = false
						clearDot(buffer, pxTrail, color.RGBA{R: 23, G: 23, B: 24, A: 255})
					}
				}
				// Draw a 3x3 dot for each vertex
				drawDot(buffer, px, py, color.RGBA{0, 255, 128, 255})
			}

			// 3. Tell Fyne to repaint the updated buffer on the UI thread
			fyne.Do(func() {
				img.Refresh()
			})
		}
	}()

	w.ShowAndRun()

}

func drawDot(img *image.RGBA, cx, cy int, c color.Color) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			img.Set(cx+dx, cy+dy, c)
		}
	}
}

func clearBuffer(buff *image.RGBA) {
	clear(buff.Pix)
}

func clearDot(img *image.RGBA, cxy [8][8][2]int, col color.Color) {
	for i, xy := range cxy {
		if len(xy) < 8 {
			continue
		}
		for _, px := range xy {
			if len(px) < 2 {
				continue
			}
			if i == 8 {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						img.Set(px[0]+dx, px[1]+dy, col)
					}
				}
			}
		}
	}
}
