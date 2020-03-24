package shape

import (
	"gioui.org/f32"
	"math"
)

func slope(p1, p2 f32.Point) float32 {
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func slope2(p1, p2 f32.Point) float32 {
	p1.Y *= -1
	p2.Y *= -1
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func angle(p1, p2 f32.Point) float32 {
	return atan2(p2.Y-p1.Y, p2.X-p1.X)
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
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

type Vector struct {
	a, b float32
}

func (v Vector) Normalize() Vector {
	m := v.Magnitude()
	return Vector{v.a / m, v.b / m}
}

func (v Vector) Magnitude() float32 {
	return sqrt(pow2(v.a) + pow2(v.b))
}

//func ToUnitVector(v Vector) Vector {
//
//}

func miter(p1, p2, p3 f32.Point) {

}

func pow2(f float32) float32 {
	return float32(math.Pow(float64(f), 2))
}

func sqrt(f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

// Return angle in radian between the vectors pq and pr
//func across(p, q, r f32.Point) float32 {
//	return atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
//}

func bezel(p, r, q f32.Point, below bool) float32 {
	angle := angle(q, p) - angle(r, p)
	//if angle < 0 {
	//	angle = 2*math.Pi + angle
	//}
	//if angle > rad180  && below {
	//	angle -= rad180
	//} else if angle < rad180 && !below {
	//	angle += rad180
	//}
	//flip := float32(1)
	//run1 := p.X - q.X
	//rise1 := p.Y - q.Y
	//if below {
	//	flip = -1
	//}
	//s1 := rise1 / run1 * flip
	//if s1 == 0 {
	//	angle = angle + math.Pi + rad45
	//}
	//if s1 < 0  {
	//	angle = angle + math.Pi
	//}
	//if s1 == 0 && below {
	//	angle = angle + math.Pi + rad45
	//} else if s1 < 0 && below {
	//	angle = angle + math.Pi
	//}
	return angle
}

func slopes(p, q, r f32.Point, below bool) (float32, float32) {
	if below {
		q, r = r, q
	}
	//if q.X > p.X || q.Y > p.Y {
	//	q, p = p, q
	//}
	s1 := slope(q, p)
	s2 := slope(p, r)
	return s1, s2
}

func metrict(p1, p2 f32.Point) (float32, float32) {
	// Run && Rise
	// Left-to-right, Ascending,
	return p2.X - p1.X, p2.Y - p1.Y
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
