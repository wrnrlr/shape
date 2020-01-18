package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
)

type Triangle struct {
	A, B, C f32.Point
}

func (t Triangle) Fill(rgba color.RGBA, gtx *layout.Context) (bbox f32.Rectangle) {
	pp2 := t.B.Sub(t.A)
	pp3 := t.C.Sub(t.B)
	pp1 := t.A.Sub(t.C)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{rgba}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(t.A)
	path.Line(pp2)
	path.Line(pp3)
	path.Line(pp1)
	path.End().Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()
	return bbox
}

func (t Triangle) Stroke(rgba color.RGBA, width float32, gtx *layout.Context) (bbox f32.Rectangle) {
	p2 := t.B.Sub(t.A)
	p3 := t.C.Sub(t.B)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{rgba}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(t.A)
	path.Line(p2)
	path.Line(p3)
	path.End().Add(gtx.Ops)
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	return bbox
}
