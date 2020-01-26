package shape

import (
	"gioui.org/f32"
	"testing"
)

func TestSlope(t *testing.T) {
	s := slope(f32.Point{0, 0}, f32.Point{1, 1})
	isTrue(t, s > 0)
	s = slope(f32.Point{0, 0}, f32.Point{0, -1})
	isTrue(t, s < 0)
	s = slope(f32.Point{0, 0}, f32.Point{0, 1})
	isTrue(t, s > 0)
	s = slope(f32.Point{0, 0}, f32.Point{1, -1})
	isTrue(t, s < 0)
}

func isEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf("not equal")
		panic(1)
	}
}

func isTrue(t *testing.T, expected bool) {
	if !expected {
		t.Fatalf("Should be true")
		panic(1)
	}
}
