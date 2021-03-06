// Emacs tricks for utf-8:
//
//    C-x RET f utf-8 RET
//    M-x ucs-insert GREEK <tab>
//    M-x ucs-insert DEVANGARI <tab>

package vmath_test

import (
    "fmt"
    "math"
    "testing"
    . "vmath"
)

var ᴨ float32 = float32(math.Atan(1) * 4)
var ε float32 = 1e-4

func BenchmarkVectors(b *testing.B) {
    fmt.Println("No benchmarks yet...")
}

func TestV3andM3(t *testing.T) {

    i := V3{1, 0, 0}
    j := V3{0, 1, 0}
    k := V3{0, 0, 1}

    // Right-hand rule:
    îĵ := i.Cross(j)
    if !îĵ.Equivalent(k, ε) {
        t.Error("Cross product")
    }

    // Rotation about Z
    M := M3RotateZ(ᴨ / 2)
    v := M.Mul(i)
    if !v.Equivalent(j, ε) {
        t.Error("M3 Rotation about Z", v)
    }
    v = M.Mul(j)
    if !v.Equivalent(V3{-1, 0, 0}, ε) {
        t.Error("M3 Rotation about Z", v)
    }

    // Rotation about Y
    M = M3RotateY(ᴨ / 2)
    v = M.Mul(i)
    if !v.Equivalent(k, ε) {
        t.Error("M3 Rotation about Y: ", v)
    }

    // Rotation about X
    M = M3RotateX(-ᴨ / 2)
    v = M.Mul(j)
    if !v.Equivalent(k, ε) {
        t.Error("M3 Rotation about X: ", v)
    }
}

func TestT3andM4(t *testing.T) {

    // P3's are like V4's with w=1
    pi, vi := P3{0.5, 0, 0}, V4{0.5, 0, 0, 1}
    pj, vj := P3{0, 0.5, 0}, V4{0, 0.5, 0, 1}
    pk, vk := P3{0, 0, 0.5}, V4{0, 0, 0.5, 1}

    // Rotation about Z
    M, T := M4RotateZ(ᴨ/2), T3RotateZ(ᴨ/2)
    v, p := M.Mul(vi), T.Transform(pi)
    if !v.Equivalent(vj, ε) {
        t.Error("M4 rotation about Z: ", v)
    }
    if !p.Equivalent(pj, ε) {
        t.Error("P3 rotation about Z: ", p)
    }
    v, p = M.Mul(vj), T.Transform(pj)
    if !v.Equivalent(V4{-0.5, 0, 0, 1}, ε) {
        t.Error("M4 rotation about Z: ", v)
    }
    if !p.Equivalent(P3{-0.5, 0, 0}, ε) {
        t.Error("P4 rotation about Z: ", p)
    }

    // Rotation about Y
    M, T = M4RotateY(ᴨ/2), T3RotateY(ᴨ/2)
    v, p = M.Mul(vi), T.Transform(pi)
    if !v.Equivalent(vk, ε) {
        t.Error("M4 rotation about Y: ", v)
    }
    if !p.Equivalent(pk, ε) {
        t.Error("T3 rotation about Y: ", p)
    }

    // Rotation about X
    M, T = M4RotateX(-ᴨ/2), T3RotateX(-ᴨ/2)
    v, p = M.Mul(vj), T.Transform(pj)
    if !v.Equivalent(vk, ε) {
        t.Error("M4 rotation about X: ", v)
    }
    if !p.Equivalent(pk, ε) {
        t.Error("T3 rotation about X: ", p)
    }
}

// More of an API sandbox than a correctness test
func TestVectorsAndPoints(test *testing.T) {
    v := V3{0, 0, 1}
    p := P3{1, 0, 0}
    p = p.Add(v) // point + vector = point
    p = p.Sub(v) // point - vector = point
    t := T3Translate(0, 0, 1)
    m := M3RotateX(ᴨ)

    p = t.Transform(p) // 4x3 * 3x1 = 4x1 
    p = p.Transform(t) // 1x4 * 4x3 = 1x3

    var v3 V3 = V3{1, 0, 0}
    fmt.Printf("v3 is a %T\n", v3)

    var v4 V4 = t.Mul(v3) // 4x3 * 3x1 = 4x1
    v3 = v4.Transform(t)  // 1x4 * 4x3 = 1x3

    v = m.Mul(v)

    fmt.Println(v, p, v3, v4)
}

// Test transforms-of-transforms
func TestComposition(test *testing.T) {
    t := T3RotateZ(ᴨ / 4)
    p := P3{1, 0, 0}
    p1 := t.Transform(p)
    p2 := t.Transform(p1)
    if !p2.Equivalent(P3{0, 1, 0}, ε) {
        test.Error()
    }
    p2b := t.Compose(t).Transform(p)
    if !p2.Equivalent(p2b, ε) {
        test.Error()
    }

    m := M3RotateZ(ᴨ / 4)
    v := V3FromP3(p)
    v1 := m.Mul(v)
    v2 := m.Mul(v1)
    if !v2.Equivalent(V3{0, 1, 0}, ε) {
        test.Error()
    }
    v2b := m.Compose(m).Mul(v)
    if !v2.Equivalent(v2b, ε) {
        test.Error()
    }

    m4 := M4RotateZ(ᴨ / 4)
    v4 := V4FromP3(p)
    v41 := m4.Mul(v4)
    v42 := m4.Mul(v41)
    if !v42.Equivalent(V4{0, 1, 0, 1}, ε) {
        test.Error()
    }
    v42b := m4.Compose(m4).Mul(v4)
    if !v42.Equivalent(v42b, ε) {
        test.Error()
    }

    α := m4.GetUpperLeft()
    β := t.GetUpperLeft()
    if !α.Equivalent(β, ε) || !m.Equivalent(β, ε) {
        test.Error()
    }
}
