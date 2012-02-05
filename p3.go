package vmath

import "fmt"

type P3 struct {
    X, Y, Z float32
}

func P3FromV3(v V3) P3 {
    return P3{v.X, v.Y, v.Z}
}

func P3FromV4(v V4) P3 {
    return P3{v.X, v.Y, v.Z}
}

func (a P3) Distance(b P3) float32 {
    return V3FromP3(a.Sub(V3FromP3(b))).Length()
}

func (a P3) Add(b V3) P3 {
    return P3{
        a.X + b.X,
        a.Y + b.Y,
        a.Z + b.Z}
}

func (a P3) Sub(b V3) P3 {
    return P3{
        a.X - b.X,
        a.Y - b.Y,
        a.Z - b.Z}
}

func (p P3) Clone() P3 {
    return P3{p.X, p.Y, p.Z}
}

func (p P3) Transform(t *T3) P3 {
    m := &t.matrix
    c0 := V3{m[0], m[3], m[6]}
    c1 := V3{m[1], m[4], m[7]}
    c2 := V3{m[2], m[5], m[8]}
    v := V3FromP3(p)
    x := c0.Dot(v) + m[9]
    y := c1.Dot(v) + m[10]
    z := c2.Dot(v) + m[11]
    return P3{x, y, z}
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
