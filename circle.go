package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
)

const c = 0.55228475 // 4*(sqrt(2)-1)/3

type Circle struct {
	Center f32.Point
	Radius float32
}

func (cc Circle) Stroke(col color.RGBA, width float32, gtx *layout.Context) f32.Rectangle {
	r := cc.Radius
	//scale := width // base value on lineWidth and r
	w, h := r*2, r*2
	p := cc.Center
	box := f32.Rectangle{Max: f32.Point{X: p.X + w, Y: p.Y + h}}
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{col}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: p.X + w, Y: p.Y + r})

	//path.Move(f32.Point{X: w, Y: h - r})
	path.Cube(f32.Point{X: 0, Y: r * c}, f32.Point{X: -r + r*c, Y: r}, f32.Point{X: -r, Y: r})    // SE
	path.Cube(f32.Point{X: -r * c, Y: 0}, f32.Point{X: -r, Y: -r + r*c}, f32.Point{X: -r, Y: -r}) // SW
	path.Cube(f32.Point{X: 0, Y: -r * c}, f32.Point{X: r - r*c, Y: -r}, f32.Point{X: r, Y: -r})   // NW
	path.Cube(f32.Point{X: r * c, Y: 0}, f32.Point{X: r, Y: r - r*c}, f32.Point{X: r, Y: r})      // NE
	// Return to origin
	path.Move(f32.Point{X: -w, Y: -r})
	scale := (r - width*2) / r
	path.Move(f32.Point{X: w * (1 - scale) * .5, Y: h * (1 - scale) * .5})
	w *= scale
	h *= scale
	r *= scale
	path.Move(f32.Point{X: 0, Y: h - r})
	path.Cube(f32.Point{X: 0, Y: r * c}, f32.Point{X: +r - r*c, Y: r}, f32.Point{X: +r, Y: r})      // SW
	path.Cube(f32.Point{X: +r * c, Y: 0}, f32.Point{X: +r, Y: -r + r*c}, f32.Point{X: +r, Y: -r})   // SE
	path.Cube(f32.Point{X: 0, Y: -r * c}, f32.Point{X: -(r - r*c), Y: -r}, f32.Point{X: -r, Y: -r}) // NE
	path.Cube(f32.Point{X: -r * c, Y: 0}, f32.Point{X: -r, Y: r - r*c}, f32.Point{X: -r, Y: r})     // NW
	path.End().Add(gtx.Ops)
	//paint.PaintOp{f32.Rectangle{Max:f32.Point{w,h}}}.Add(gtx.Ops)
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	return box
}

func (cc Circle) Fill(col color.RGBA, gtx *layout.Context) f32.Rectangle {
	p := cc.Center
	r := cc.Radius
	w, h := r*2, r*2
	var stack op.StackOp
	stack.Push(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: p.X, Y: p.Y})
	path.Move(f32.Point{X: w, Y: h - r})
	path.Cube(f32.Point{X: 0, Y: r * c}, f32.Point{X: -r + r*c, Y: r}, f32.Point{X: -r, Y: r})    // SE
	path.Cube(f32.Point{X: -r * c, Y: 0}, f32.Point{X: -r, Y: -r + r*c}, f32.Point{X: -r, Y: -r}) // SW
	path.Cube(f32.Point{X: 0, Y: -r * c}, f32.Point{X: r - r*c, Y: -r}, f32.Point{X: r, Y: -r})   // NW
	path.Cube(f32.Point{X: r * c, Y: 0}, f32.Point{X: r, Y: r - r*c}, f32.Point{X: r, Y: r})      // NE
	path.End().Add(gtx.Ops)
	box := f32.Rectangle{Max: f32.Point{X: p.X + w, Y: p.Y + h}}
	//paint.PaintOp{f32.Rectangle{Max:f32.Point{w,h}}}.Add(gtx.Ops)
	paint.ColorOp{col}.Add(gtx.Ops)
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	return box
}
