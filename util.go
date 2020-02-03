package shape

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"image/color"
	"math"
)

var (
	red   = color.RGBA{255, 0, 0, 125}
	green = color.RGBA{0, 255, 0, 125}
	blue  = color.RGBA{0, 0, 255, 125}
	black = color.RGBA{0, 0, 0, 255}
	grey  = color.RGBA{125, 125, 125, 125}
)

func OffsetPoints(l Line, width float32) (offset []f32.Point) {
	distance := width
	var angles []float32
	var offsetPoints, originalPoints, deltaPoints []f32.Point
	var tilt float32
	var prevDelta f32.Point
	for i, point := range l {
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
		}
		angles = append(angles, tilt)
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
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
			tilt = bezel(point, nextPoint, prevPoint, true)
		}
		angles = append(angles, tilt)
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
	}
	point := l[0]
	nextPoint := l[1]
	tilt = angle(point, nextPoint) + rad225
	angles = append(angles, tilt)
	originalPoints = append(originalPoints, point)
	point = offsetPoint(point, distance, tilt)
	offsetPoints = append(offsetPoints, point)
	point = point.Sub(prevDelta)
	deltaPoints = append(deltaPoints, point)
	return offsetPoints
}

func printDegrees(radials []float32) {
	var degrees []float32
	for _, a := range radials {
		degrees = append(degrees, mod(a*180/math.Pi, 360))
	}
	fmt.Printf("Angles: %v\n", degrees)
}

func PaintPoints(ps []f32.Point, radius float32, gtx *layout.Context) {
	for _, p := range ps {
		Circle{p, radius}.Fill(black, gtx)
	}
}
