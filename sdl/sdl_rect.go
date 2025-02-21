package sdl

import "math"

// Rect is a rectangle, with the origin at the upper left (using integers).
type Rect struct {
	X, Y, W, H int32
}

// FRect is a rectangle, with the origin at the upper left (using floating point values).
type FRect struct {
	X, Y, W, H float32
}

type FPoint struct {
	X, Y float32
}

type Point struct {
	X, Y int32
}

// RectToFRect converts an integer rectangle to a float rectangle.
func RectToFRect(rect Rect) FRect {
	return FRect{float32(rect.X), float32(rect.Y), float32(rect.W), float32(rect.H)}
}

// PointInRect returns true if p is located inside r.
func PointInRect(p Point, r Rect) bool {
	return (p.X >= r.X) && (p.X < (r.X + r.W)) && (p.Y >= r.Y) && (p.Y < (r.Y + r.H))
}

// PointInRectFloat returns true if p is located inside r.
func PointInRectFloat(p FPoint, r FRect) bool {
	return (p.X >= r.X) && (p.X < (r.X + r.W)) && (p.Y >= r.Y) && (p.Y < (r.Y + r.H))
}

// RectEmpty returns true if r has no area.
func RectEmpty(r Rect) bool {
	return r.W <= 0 || r.H <= 0
}

// RectEmptyFloat returns true if r has no area.
func RectEmptyFloat(r FRect) bool {
	return r.W <= 0 || r.H <= 0
}

func GetRectAndLineIntersection(rect Rect, x1, y1, x2, y2 *int32) bool {
	return sdlGetRectAndLineIntersection(&rect, x1, y1, x2, y2)
}

func GetRectAndLineIntersectionFloat(rect FRect, x1, y1, x2, y2 *float32) bool {
	return sdlGetRectAndLineIntersectionFloat(&rect, x1, y1, x2, y2)
}

func GetRectEnclosingPoints(points []Point, clip *Rect) (Rect, bool) {
	var result Rect
	var pointsPtr *Point
	if len(points) > 0 {
		pointsPtr = &points[0]
	}
	ret := sdlGetRectEnclosingPoints(pointsPtr, int32(len(points)), clip, &result)
	return result, ret
}

func GetRectEnclosingPointsFloat(points []FPoint, clip *FRect) (FRect, bool) {
	var result FRect
	var pointsPtr *FPoint
	if len(points) > 0 {
		pointsPtr = &points[0]
	}
	ret := sdlGetRectEnclosingPointsFloat(pointsPtr, int32(len(points)), clip, &result)
	return result, ret
}

func GetRectIntersection(a, b Rect) (Rect, bool) {
	var result Rect
	ret := sdlGetRectIntersection(&a, &b, &result)
	return result, ret
}

func GetRectIntersectionFloat(a, b FRect) (FRect, bool) {
	var result FRect
	ret := sdlGetRectIntersectionFloat(&a, &b, &result)
	return result, ret
}

func GetRectUnion(a, b Rect) (Rect, bool) {
	var result Rect
	ret := sdlGetRectUnion(&a, &b, &result)
	return result, ret
}

func GetRectUnionFloat(a, b FRect) (FRect, bool) {
	var result FRect
	ret := sdlGetRectUnionFloat(&a, &b, &result)
	return result, ret
}

func HasRectIntersection(a, b Rect) bool {
	return sdlHasRectIntersection(&a, &b)
}

func HasRectIntersectionFloat(a, b FRect) bool {
	return sdlHasRectIntersectionFloat(&a, &b)
}

func RectsEqualEpsilon(a, b FRect, epsilon float32) bool {
	return (float32(math.Abs(float64(a.X-b.X))) <= epsilon) &&
		(float32(math.Abs(float64(a.Y-b.Y))) <= epsilon) &&
		(float32(math.Abs(float64(a.W-b.W))) <= epsilon) &&
		(float32(math.Abs(float64(a.H-b.H))) <= epsilon)
}

func RectsEqualFloat(a, b FRect) bool {
	return RectsEqualEpsilon(a, b, FltEpsilon)
}

func RectsEqual(a, b Rect) bool {
	return a == b
}
