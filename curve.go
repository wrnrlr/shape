package shape

import "gioui.org/f32"

// https://stackoverflow.com/questions/14398178/plotting-graphs-using-bezier-curves
// https://en.wikipedia.org/wiki/Cubic_Hermite_spline

type Cube struct {
	Point0, Point1, Point2, Point3 f32.Point
}

type Quad struct {
	Point0, Point1, Point2 f32.Point
}
