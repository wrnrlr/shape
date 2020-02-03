package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/wrnrlr/shape"
	"image/color"
)

var (
	red   = color.RGBA{255, 0, 0, 125}
	green = color.RGBA{0, 255, 0, 125}
	blue  = color.RGBA{0, 0, 255, 125}
	black = color.RGBA{0, 0, 0, 255}
	grey  = color.RGBA{125, 125, 125, 125}
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		gofont.Register()
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				layout.UniformInset(unit.Dp(30)).Layout(gtx, func() {
					painting(gtx)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func painting(gtx *layout.Context) {
	//defaultWidth := float32(gtx.Px(unit.Sp(2)))

	//width := float32(gtx.Px(unit.Sp(10)))
	//thin := float32(gtx.Px(unit.Sp(1)))
	//light := float32(gtx.Px(unit.Sp(2)))
	//thick := float32(gtx.Px(unit.Sp(5)))
	bold := float32(gtx.Px(unit.Sp(20)))

	var a, b, c f32.Point

	//w, h := float32(gtx.Constraints.Width.Max), float32(gtx.Constraints.Height.Max)

	a, b = f32.Point{0, 100}, f32.Point{100, 0}
	c = f32.Point{200, 100}
	l2y := shape.Line{a, b, c}
	l2y.Stroke(red, bold, gtx)

	a, b = f32.Point{500, 100}, f32.Point{400, 0}
	c = f32.Point{300, 100}
	l1a := shape.Line{a, b, c}
	l1a.Stroke(red, bold, gtx)

	a, b = f32.Point{600, 0}, f32.Point{700, 100}
	c = f32.Point{800, 0}
	l2x := shape.Line{a, b, c}
	l2x.Stroke(red, bold, gtx)

	a, b = f32.Point{1100, 0}, f32.Point{1000, 100}
	c = f32.Point{900, 0}
	l2b := shape.Line{a, b, c}
	l2b.Stroke(red, bold, gtx)

	a, b = f32.Point{0, 200}, f32.Point{200, 200}
	c = f32.Point{200, 300}
	l2y = shape.Line{a, b, c}
	l2y.Stroke(red, bold, gtx)

	//shape.PaintPoints(l2y, light, gtx)
	//shape.PaintPoints(shape.OffsetPoints(l2y, bold), light, gtx)

	fmt.Println("=================================================")
}
