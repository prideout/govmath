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
    fmt.Println("No benchmarks yet.")
}

func TestV3andM3(t *testing.T) {

    i := V3New(1, 0, 0)
    j := V3New(0, 1, 0)
    k := V3New(0, 0, 1)

    // Right-hand rule:
    îĵ := i.Cross(j)
    if !îĵ.Equivalent(k, ε) {
        t.Error("Cross product")
    }

    // Rotation about Z
    M := M3RotateZ(ᴨ / 2)
    v := M.MulV3(i)
    if !v.Equivalent(j, ε) {
        t.Error("M3 Rotation about Z", v)
    }
    v = M.MulV3(j)
    if !v.Equivalent(V3New(-1, 0, 0), ε) {
        t.Error("M3 Rotation about Z", v)
    }

    // Rotation about Y
    M = M3RotateY(ᴨ / 2)
    v = M.MulV3(i)
    if !v.Equivalent(k, ε) {
        t.Error("M3 Rotation about Y: ", v)
    }

    // Rotation about X
    M = M3RotateX(-ᴨ / 2)
    v = M.MulV3(j)
    if !v.Equivalent(k, ε) {
        t.Error("M3 Rotation about X: ", v)
    }
}

func TestT3andM4(t *testing.T) {

    // P3's are like V4's with w=1
    pi, vi := P3New(0.5, 0, 0), V4New(0.5, 0, 0, 1)
    pj, vj := P3New(0, 0.5, 0), V4New(0, 0.5, 0, 1)
    pk, vk := P3New(0, 0, 0.5), V4New(0, 0, 0.5, 1)

    // Rotation about Z
    M, T := M4RotateZ(ᴨ/2), T3RotateZ(ᴨ/2)
    v, p := M.MulV4(vi), T.MulP3(pi)
    if !v.Equivalent(vj, ε) {
        t.Error("M4 rotation about Z: ", v)
    }
    if !p.Equivalent(pj, ε) {
        t.Error("P3 rotation about Z: ", p)
    }
    v, p = M.MulV4(vj), T.MulP3(pj)
    if !v.Equivalent(V4New(-0.5, 0, 0, 1), ε) {
        t.Error("M4 rotation about Z: ", v)
    }
    if !p.Equivalent(P3New(-0.5, 0, 0), ε) {
        t.Error("P4 rotation about Z: ", p)
    }

    // Rotation about Y
    M, T = M4RotateY(ᴨ/2), T3RotateY(ᴨ/2)
    v, p = M.MulV4(vi), T.MulP3(pi)
    if !v.Equivalent(vk, ε) {
        t.Error("M4 rotation about Y: ", v)
    }
    if !p.Equivalent(pk, ε) {
        t.Error("T3 rotation about Y: ", p)
    }

    // Rotation about X
    M, T = M4RotateX(-ᴨ / 2), T3RotateX(-ᴨ / 2)
    v, p = M.MulV4(vj), T.MulP3(pj)
    if !v.Equivalent(vk, ε) {
        t.Error("M4 rotation about X: ", v)
    }
    if !p.Equivalent(pk, ε) {
        t.Error("T3 rotation about X: ", p)
    }
}

// Test transforms-of-transforms
func TestComposition(t *testing.T) {

}