package vmath

import "fmt"

// https://bitbucket.org/prideout/pez-viewer/src/11899f6b6f02/vmath.h

type V4 struct {
    X, Y, Z, W float32
}

func V4New(x, y, z, w float32) V4 {
    v := new(V4)
    v.X = x
    v.Y = y
    v.Z = z
    v.W = w
    return *v
}

func (a V4) Dot(b V4) float32 {
    return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z) + (a.W * b.W)
}

func (a V4) Add(b V4) V4 {
    return V4New(
        a.X+b.X,
        a.Y+b.Y,
        a.Z+b.Z,
        a.W+b.W)
}

func (a V4) Sub(b V4) V4 {
    return V4New(
        a.X-b.X,
        a.Y-b.Y,
        a.Z-b.Z,
        a.W-b.W)
}

func (v V4) Clone() V4 {
    return V4New(v.X, v.Y, v.Z, v.W)
}

func (v V4) Length() float32 {
    return sqrt(v.Dot(v))
}

func (v V4) Transform(t *T3) V3 {
    m := &t.matrix
    c0 := V4{m[0], m[3], m[6], m[9]}
    c1 := V4{m[1], m[4], m[7], m[10]}
    c2 := V4{m[2], m[5], m[8], m[11]}
    x := c0.Dot(v)
    y := c1.Dot(v)
    z := c2.Dot(v)
    return V3{x, y, z}
}

func (a V4) Equivalent(b V4, ε float32) bool {
    return true &&
        abs(b.X-a.X) < ε &&
        abs(b.Y-a.Y) < ε &&
        abs(b.Z-a.Z) < ε &&
        abs(b.W-a.W) < ε
}

func (v V4) String() string {
    return fmt.Sprintf("(%g, %g, %g, %g)", v.X, v.Y, v.Z, v.W)
}
