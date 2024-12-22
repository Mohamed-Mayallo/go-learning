package main

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	w, h float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.w + r.h)
}

func (r Rectangle) Perimeter() float64 {
	return r.h*2 + r.w*2
}

func (r Rectangle) Area() float64 {
	return r.h * r.w
}

type Circle struct {
	r float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}
