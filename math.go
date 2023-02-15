package main

import (
	"fmt"
)

type Vec2i struct {
	X, Y int
}

type Vec2f struct {
	X, Y float64
}

func (v Vec2i) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func (v Vec2f) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", v.X, v.Y)
}
