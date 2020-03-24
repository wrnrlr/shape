package main

import (
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

	var a, b, c, d, e, f, g f32.Point
	var l shape.Line

	//w, h := float32(gtx.Constraints.Width.Max), float32(gtx.Constraints.Height.Max)

	a = f32.Point{0, 100}
	b = f32.Point{100, 0}
	c = f32.Point{200, 100}
	l = shape.Line{a, b, c}
	l.Stroke(red, bold, gtx)

	//a = f32.Point{500, 100}
	//b = f32.Point{400, 0}
	//c = f32.Point{300, 100}
	//l1a := shape.Line{a, b, c}
	//l1a.Stroke(red, bold, gtx)

	a = f32.Point{600, 0}
	b = f32.Point{700, 100}
	c = f32.Point{800, 0}
	l2x := shape.Line{a, b, c}
	l2x.Stroke(red, bold, gtx)

	//a = f32.Point{1100, 0}
	//b = f32.Point{1000, 100}
	//c = f32.Point{900, 0}
	//l2b := shape.Line{a, b, c}
	//l2b.Stroke(red, bold, gtx)

	//a = f32.Point{0, 200}
	//b = f32.Point{200, 200}
	//c = f32.Point{200, 300}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{300, 200}
	//b = f32.Point{500, 200}
	//c = f32.Point{500, 300}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{600, 300}
	//b = f32.Point{600, 200}
	//c = f32.Point{800, 200}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{900, 200}
	//b = f32.Point{900, 300}
	//c = f32.Point{1100, 300}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{0, 400}
	//b = f32.Point{100, 500}
	//c = f32.Point{0, 600}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{200, 600}
	//b = f32.Point{300, 500}
	//c = f32.Point{200, 400}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{500, 400}
	//b = f32.Point{400, 500}
	//c = f32.Point{500, 600}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)
	//
	//a = f32.Point{700, 600}
	//b = f32.Point{600, 500}
	//c = f32.Point{700, 400}
	//l = shape.Line{a, b, c}
	//l.Stroke(red, bold, gtx)

	//a = f32.Point{0, 500}
	//b = f32.Point{100, 600}
	//c = f32.Point{200, 500}
	//d := f32.Point{300, 600}
	//e := f32.Point{400, 500}
	//f := f32.Point{500, 600}
	//g := f32.Point{600, 500}
	//l = shape.Line{a, b, c, d, e, f, g}
	//l.Stroke(red, bold, gtx)

	a = f32.Point{0, 800}
	b = f32.Point{100, 700}
	c = f32.Point{200, 800}
	d = f32.Point{300, 700}
	e = f32.Point{400, 800}
	f = f32.Point{500, 700}
	g = f32.Point{600, 800}
	l = shape.Line{a, b, c, d, e, f, g}
	l.Stroke(red, bold, gtx)

	//shape.PaintPoints(l, light, gtx)
	//shape.PaintPoints(shape.OffsetPoints(l, bold), light, gtx)

	//fmt.Println("=================================================")
}
