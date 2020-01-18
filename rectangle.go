package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
)

type Rectangle struct {
	P1, P2 f32.Point
}

func (r Rectangle) Fill(rgba color.RGBA, gtx *layout.Context) f32.Rectangle {
	p1, p2 := r.P1, r.P2
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{rgba}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.End().Add(gtx.Ops)
	box := f32.Rectangle{Min: p1, Max: p2}
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	return box
}

func (r Rectangle) Stroke(rgba color.RGBA, lineWidth float32, gtx *layout.Context) f32.Rectangle {
	p1, p2 := r.P1, r.P2
	box := f32.Rectangle{Min: p1, Max: p2}
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{rgba}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.Move(f32.Point{X: lineWidth, Y: lineWidth})
	p2.X -= lineWidth * 2
	p2.Y -= lineWidth * 2
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.End().Add(gtx.Ops)
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	return box
}
