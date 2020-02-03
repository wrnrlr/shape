package shape

import (
	"fmt"
	"gioui.org/f32"
	"math"
)

func slope(p1, p2 f32.Point) float32 {
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func angle(p1, p2 f32.Point) float32 {
	return float32(math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X)))
}

func cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func pointDeltas(points []f32.Point) []f32.Point {
	deltas := make([]f32.Point, len(points))
	var prev f32.Point
	for i, p := range points {
		deltas[i] = p.Sub(prev)
		prev = p
	}
	return deltas
}

func offsetPoint(point f32.Point, distance, angle float32) f32.Point {
	//fmt.Printf("Point X: %f, Y: %f, Angle: %f\n", point.X, point.Y, angle)
	x := point.X + distance*cos(angle)
	y := point.Y + distance*sin(angle)
	//fmt.Printf("Point X: %f, Y: %f \n", x, y)
	return f32.Point{X: x, Y: y}
}

// Return angle in radian between the vectors pq and pr
func across(p, q, r f32.Point) float32 {
	return atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
}

func bezel(p, q, r f32.Point, below bool) float32 {
	angle := atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
	if angle < -rad180 || angle > rad180 { // concave
		//fmt.Printf("Concave point\n")
	} else { // convex
		//fmt.Printf("Convex point\n")
	}
	s1 := slope(q, p)
	s2 := slope(p, r)
	if s1 > 0 {
		fmt.Printf("Slope 1: Down, ")
	} else {
		fmt.Printf("Slope 1: Up, ")
	}
	if s2 > 0 {
		fmt.Printf("Slope 2: Down \n")
	} else {
		fmt.Printf("Slope 2: Up \n")
	}
	if angle > 0 && !below {
		angle = angle * -1
	} else if angle < 0 && below {
		angle = angle * -1
	}
	//if angle < 0 {
	//	angle = mod(angle + rad360, 360)
	//} else {
	//	angle = angle
	//}
	//if s1 >= 0 {
	//	angle = angle * -1
	//}
	//if s2 < 0 {
	//	angle = angle * -1
	//}
	return angle
}

func bezel2(p, q, r f32.Point) float32 {
	angle := atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
	if angle < 0 {
		angle = angle + 360
	}
	return angle
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

func Overlay(rs ...f32.Rectangle) f32.Rectangle {
	rr := rs[0]
	for _, r := range rs[1:] {
		rr = f32.Rectangle{
			Min: f32.Point{Min(rr.Min.Y, r.Min.Y), Min(rr.Min.Y, r.Min.Y)},
			Max: f32.Point{Max(rr.Max.Y, r.Max.Y), Max(rr.Max.Y, r.Max.Y)},
		}
	}
	return rr
}

func Min(x, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

func Max(x, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}

func Inf(sign int) float32 {
	return float32(math.Inf(sign))
}

func IsNaN(sign float32) bool {
	return math.IsNaN(float64(sign))
}

func IsInf(f float32, sign int) bool {
	return math.IsInf(float64(f), sign)
}

//func pointAngles(points []f32.Point) []float32 {
//	var angles []float32
//	for i, point := range points {
//		if i == 0 {
//			angles = append(angles, rad225)
//		} else if i == len(points)-1 {
//			angles = append(angles, rad315)
//		} else {
//			prevPoint := points[i-1]
//			nextPoint := points[i+1]
//			a := across(point, nextPoint, prevPoint) - rad90
//			angles = append(angles, a)
//		}
//	}
//	for i := len(points) - 1; i >= 0; i-- {
//		if i == 0 {
//			angles = append(angles, rad135)
//		} else if i == len(points)-1 {
//			angles = append(angles, rad45)
//		} else {
//			point := points[i]
//			prevPoint := points[i-1]
//			nextPoint := points[i+1]
//			a := across(point, nextPoint, prevPoint) + rad90
//			angles = append(angles, a)
//		}
//	}
//	return angles
//}

//func pointOffsets(line []f32.Point, distance float32) []f32.Point {
//	fmt.Printf("==================================================================\n")
//	fmt.Printf("Original Points: %v\n", line)
//	angles := pointAngles(line)
//	//fmt.Println(angles)
//	var pps []f32.Point
//	j := 0
//	for i, p := range line {
//		a := angles[i]
//		pps = append(pps, offsetPoint(p, distance, a))
//		j++
//	}
//	lastIndex := len(pps) - 1
//	for i := lastIndex; i >= 0; i-- {
//		p := line[i]
//		a := angles[j]
//		pps = append(pps, offsetPoint(p, distance, a))
//		j++
//	}
//	pps = append(pps, pps[0])
//	fmt.Printf("Delta Points:    %v\n", pps)
//	pos := pointDeltas(pps)
//	fmt.Printf("Offset Points:   %v\n", pos)
//	fmt.Printf("==================================================================\n")
//	return pos
//}
