package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
	"math"
)

const (
	rad45  = float32(45 * math.Pi / 180)
	rad90  = float32(90 * math.Pi / 180)
	rad135 = float32(135 * math.Pi / 180)
	rad180 = float32(180 * math.Pi / 180)
	rad225 = float32(225 * math.Pi / 180)
	rad270 = float32(270 * math.Pi / 180)
	rad315 = float32(315 * math.Pi / 180)
	rad360 = float32(360 * math.Pi / 180)
)

type Line []f32.Point

// An algorithm for polylines outline construction
// http://old.cescg.org/CESCG99/SKrivograd/
//
// https://www.codeproject.com/Articles/226569/Drawing-polylines-by-tessellation
//
// A Realistic 2D Drawing System
// https://keithp.com/~keithp/talks/cairo2003.pdf
//
// https://stackoverflow.com/questions/5641769/how-to-draw-an-outline-around-any-line
//
// Drawing Lines is Hard
// https://mattdesl.svbtle.com/drawing-lines-is-hard
//
func (l Line) Stroke(c color.RGBA, width float32, gtx *layout.Context) (box f32.Rectangle) {

	//pointWidth := float32(gtx.Px(unit.Sp(2)))

	if len(l) < 2 {
		return box
	}
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{c}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	distance := width
	var angles []float32
	var offsetPoints, originalPoints, deltaPoints []f32.Point
	var tilt float32
	var prevDelta f32.Point
	for i, point := range l {
		distance := width
		if i == 0 {
			nextPoint := l[i+1]
			tilt = angle(point, nextPoint) + rad225
		} else if i == len(l)-1 {
			prevPoint := l[i-1]
			tilt = angle(prevPoint, point) + rad315
		} else {
			prevPoint := l[i-1]
			nextPoint := l[i+1]
			tilt = bezel(point, prevPoint, nextPoint, false)
			angles = append(angles, tilt)
		}
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
		if i == 0 {
			path.Move(newPoint)
		} else {
			path.Line(newPoint)
		}
	}
	for i := len(l) - 1; i >= 0; i-- {
		point := l[i]
		if i == 0 {
			nextPoint := l[i+1]
			tilt = angle(point, nextPoint) + rad135
		} else if i == len(l)-1 {
			prevPoint := l[i-1]
			tilt = angle(prevPoint, point) + rad45
		} else {
			point := l[i]
			prevPoint := l[i-1]
			nextPoint := l[i+1]
			//tilt = bezel(point, nextPoint, prevPoint, true)
			tilt = bezel(point, prevPoint, nextPoint, true) + rad180
			angles = append(angles, tilt)
		}
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
		path.Line(newPoint)
	}
	point := l[0]
	nextPoint := l[1]
	tilt = angle(point, nextPoint) + rad225
	//angles = append(angles, tilt)
	originalPoints = append(originalPoints, point)
	point = offsetPoint(point, distance, tilt)
	offsetPoints = append(offsetPoints, point)
	point = point.Sub(prevDelta)
	path.Line(point)
	deltaPoints = append(deltaPoints, point)
	//fmt.Printf("Original Points: %v\n", originalPoints)
	printDegrees(angles)
	//fmt.Printf("Offset Points:   %v\n", offsetPoints)
	for _, p := range offsetPoints {
		box.Min.X = Min(box.Min.X, p.X)
		box.Min.Y = Min(box.Min.Y, p.Y)
		box.Max.X = Max(box.Max.X, p.X)
		box.Max.Y = Max(box.Max.Y, p.Y)
	}
	//fmt.Printf("Min and Max:   %v\n", box)
	//fmt.Printf("Delta Points:    %v\n", deltaPoints)
	path.End().Add(gtx.Ops)
	//paint.PaintOp{f32.Rectangle{Max:f32.Point{w,h}}}.Add(gtx.Ops)
	paint.PaintOp{box}.Add(gtx.Ops)
	stack.Pop()
	//PaintPoints(offsetPoints[0:len(offsetPoints)/2], pointWidth, gtx)
	return box
}
