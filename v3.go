package vmath

import (
    "fmt"
)

// https://bitbucket.org/prideout/pez-viewer/src/11899f6b6f02/vmath.h

type V3 struct {
    X, Y, Z float32
}

func V3New(x, y, z float32) V3 {
    v := new(V3)
    v.X = x
    v.Y = y
    v.Z = z
    return *v
}

func (a V3) Dot(b V3) float32 {
    return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func (a V3) Cross(b V3) V3 {
    return V3New(
        (a.Y*b.Z)-(a.Z*b.Y),
        (a.Z*b.X)-(a.X*b.Z),
        (a.X*b.Y)-(a.Y*b.X))
}

func (a V3) Add(b V3) V3 {
    return V3New(
        a.X+b.X,
        a.Y+b.Y,
        a.Z+b.Z)
}

func (a V3) Sub(b V3) V3 {
    return V3New(
        a.X-b.X,
        a.Y-b.Y,
        a.Z-b.Z)
}

func (v V3) Clone() V3 {
    return V3New(v.X, v.Y, v.Z)
}

func (v V3) Length() float32 {
    return sqrt(v.Dot(v))
}

func (a V3) Equivalent(b V3, ε float32) bool {
    return true &&
        abs(b.X-a.X) < ε &&
        abs(b.Y-a.Y) < ε &&
        abs(b.Z-a.Z) < ε
}

func (v V3) String() string {
    return fmt.Sprintf("(%g, %g, %g)", v.X, v.Y, v.Z)
}
