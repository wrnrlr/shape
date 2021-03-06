package shape

import (
	"gioui.org/op"
	"image/color"
)

type Shape interface {
	Add(ops *op.Ops)
}

type Stroke interface {
	Shape
	Stroke(width float32, style StrokeType, rgba color.RGBA, ops *op.Ops)
}

type Fill interface {
	Stroke
	Fill(rgba color.RGBA, ops *op.Ops)
}

type StrokeType int

// stroke := shape.Stroke{red}
// fill := shape.Solid{red}
// shape.Rectangle{a,b}.Layout(gtx, stroke, fill)
//
// shape.Label{"Hello"}.Layout(gtx, stroke, fill)
