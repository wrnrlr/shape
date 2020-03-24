package shape

import (
	"gioui.org/f32"
	"image/color"
)

type Canvas struct {
	FillColor *color.RGBA

	StrokeWidth int
	StrokeColor *color.RGBA
	StrokeStyle StrokeStyle

	capper Capper
	joiner Joiner
}

func (c *Canvas) DrawLine(p1, p2 f32.Point) {

}

type StrokeStyle int

const (
	Solid StrokeStyle = iota
	Dotted
	Dashed
)

type LineJoin int

const (
	None LineJoin = iota
	Round
)

type LineCap int

const (
	Butt LineCap = iota
	Square
)
