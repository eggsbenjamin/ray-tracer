package game

type Ray struct {
	origin *Point
	angle  float64
	length float64
}

func NewRay(o *Point, a, l float64) *Ray {
	return &Ray{
		origin: o,
		angle:  a,
		length: l,
	}
}

func (r *Ray) XIntersects() []*Point {
	return GetXIntersects(r.origin, r.angle, r.length)
}

func (r *Ray) YIntersects() []*Point {
	return GetYIntersects(r.origin, r.angle, r.length)
}
