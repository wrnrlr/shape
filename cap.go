package shape

import (
	"gioui.org/f32"
)

type Capper func(offset float32, p1, p2 f32.Point) f32.Point

func ButtCapper(offset float32, p1, p2 f32.Point) f32.Point {
	return f32.Point{}
}

func RoundCapper(offset float32, p1, p2 f32.Point) f32.Point {
	return f32.Point{}
}

func SquareCapper(offset float32, p1, p2 f32.Point) float32 {
	tilt := angle(p1, p2)
	if slope(p1, p2) > 0 {
		tilt += rad315
	} else {
		tilt += rad225
	}
	return tilt
}
