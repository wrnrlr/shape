package shape

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Line2(gtx *layout.Context, width float32, ps ...f32.Point) {
	var box BBox
	if len(ps) < 2 {
		return
	}
	var stack op.StackOp
	stack.Push(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	prev := ps[0]
	for _, p := range ps[1:] {
		b := strokeLine(&path, width, prev, p)
		box = box.Union(b)
		prev = p
	}
	path.End().Add(gtx.Ops)
	paint.ColorOp{red}.Add(gtx.Ops)
	paint.PaintOp{f32.Rectangle(box)}.Add(gtx.Ops)
	stack.Pop()
}

func strokeLine(path *clip.Path, width float32, a, b f32.Point) (box BBox) {
	var deltaPoints []f32.Point
	distance := width / 2
	tilt := angle(a, b)
	a1 := offsetPoint(a, distance, tilt+rad270)
	box = box.Add(a1)
	deltaPoints = append(deltaPoints, a1)
	path.Move(a1)
	prev := a1
	a2 := offsetPoint(a, distance, tilt+rad90)
	box = box.Add(a2)
	delta := a2.Sub(prev)
	deltaPoints = append(deltaPoints, delta)
	prev = a2
	path.Line(delta)
	b1 := offsetPoint(b, distance, tilt+rad90)
	box = box.Add(b1)
	delta = b1.Sub(prev)
	deltaPoints = append(deltaPoints, delta)
	prev = b1
	path.Line(delta)
	b2 := offsetPoint(b, distance, tilt+rad270)
	box = box.Add(b2)
	delta = b2.Sub(prev)
	deltaPoints = append(deltaPoints, delta)
	prev = b2
	path.Line(delta)
	delta = a1.Sub(prev)
	deltaPoints = append(deltaPoints, delta)
	path.Line(delta)
	fmt.Printf("Original Points: %v %v %v %v %v \n", a1, a2, b1, b2, a1)
	fmt.Printf("Delta Points: %v \n", deltaPoints)
	return box
}

type BBox f32.Rectangle

func (bb BBox) Add(p f32.Point) (bb2 BBox) {
	bb2.Min.X = Min(bb.Min.X, p.X)
	bb2.Min.Y = Min(bb.Min.Y, p.Y)
	bb2.Max.X = Max(bb.Max.X, p.X)
	bb2.Max.Y = Max(bb.Max.Y, p.Y)
	return bb2
}

func (bb BBox) Union(b BBox) (bb2 BBox) {
	bb2.Min.X = Min(bb.Min.X, b.Min.X)
	bb2.Min.Y = Min(bb.Min.Y, b.Min.Y)
	bb2.Max.X = Max(bb.Max.X, b.Max.X)
	bb2.Max.Y = Max(bb.Max.Y, b.Max.Y)
	return bb2
}
