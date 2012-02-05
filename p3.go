package vmath

import "fmt"

// https://bitbucket.org/prideout/pez-viewer/src/11899f6b6f02/vmath.h

type P3 struct {
    X, Y, Z float32
}

func P3New(x, y, z float32) P3 {
    p := new(P3)
    p.X = x
    p.Y = y
    p.Z = z
    return *p
}

func P3FromV3(v V3) P3 {
    p := new(P3)
    p.X = v.X
    p.Y = v.Y
    p.Z = v.Z
    return *p
}

func P3FromV4(v V4) P3 {
    p := new(P3)
    p.X = v.X
    p.Y = v.Y
    p.Z = v.Z
    return *p
}

func (a P3) Distance(b P3) float32 {
    return a.Sub(b).Length()
}

func (a P3) Add(b V3) P3 {
    return P3New(
        a.X+b.X,
        a.Y+b.Y,
        a.Z+b.Z)
}

func (a P3) Sub(b P3) V3 {
    return V3New(
        a.X-b.X,
        a.Y-b.Y,
        a.Z-b.Z)
}

func (p P3) Clone() P3 {
    return P3New(p.X, p.Y, p.Z)
}

func (a P3) Equivalent(b P3, ε float32) bool {
    return true &&
        abs(b.X-a.X) < ε &&
        abs(b.Y-a.Y) < ε &&
        abs(b.Z-a.Z) < ε
}

func (p P3) String() string {
    return fmt.Sprintf("(%g, %g, %g)", p.X, p.Y, p.Z)
}
