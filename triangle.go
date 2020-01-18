package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
)

type Triangle struct {
	A, B, C f32.Point
}

func (t Triangle) Fill(width float32, gtx *layout.Context) (bbox f32.Rectangle) {
	pp2 := t.B.Sub(t.A)
	pp3 := t.C.Sub(t.B)
	pp1 := t.A.Sub(t.C)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(t.A)
	path.Line(pp2)
	path.Line(pp3)
	path.Line(pp1)
	path.End().Add(gtx.Ops)
	return bbox
}

func (t Triangle) Stroke(width float32, gtx *layout.Context) (bbox f32.Rectangle) {
	p2 := t.B.Sub(t.A)
	p3 := t.C.Sub(t.B)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(t.A)
	path.Line(p2)
	path.Line(p3)
	path.End().Add(gtx.Ops)
	return bbox
}
