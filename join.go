package shape

import (
	"gioui.org/f32"
	"gioui.org/op/clip"
)

type Joiner func(path clip.Path, offset float32, p1, p2, qp3 f32.Point)

func RoundJoiner(path clip.Path, offset float32, p1, p2, qp3 f32.Point) {}

func SquareJoiner(path clip.Path, offset float32, p1, p2, qp3 f32.Point) {}

func MitterJoiner(path clip.Path, offset float32, p1, p2, qp3 f32.Point) {}
