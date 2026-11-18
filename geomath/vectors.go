package geomath

import (
	"math"
)

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	X, Y, Z float64
}

type Point struct {
	Xw, Yw, Zw, Xp, Yp, Zp float64
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vec3) CrossXY(v2 Vec3) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func (v1 Vec3) CrossYZ(v2 Vec3) float64 {
	return v1.Y*v2.Z - v1.Z*v2.Y
}

func (v1 Vec3) CrossZX(v2 Vec3) float64 {
	return v1.Z*v2.X - v1.X*v2.Z
}

func (v Vec3) ProjectXY() Vec3 {
	return Vec3{X: v.X, Y: v.Y}
}

func (v Vec3) ProjectYZ() Vec3 {
	return Vec3{Y: v.Y, Z: v.Z}
}

func (v Vec3) ProjectZX() Vec3 {
	return Vec3{Z: v.Z, X: v.X}
}
