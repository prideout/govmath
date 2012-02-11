package vmath

import "fmt"

// Implements a 3x3 matrix type for 3D graphics.
type M3 struct {
    matrix [3 * 3]float32
}

// Create a 4x4 from the identity
func M3Identity() *M3 {
    m := new(M3)
    m.matrix = [3 * 3]float32{
        1, 0, 0,
        0, 1, 0,
        0, 0, 1}
    return m
}

// Create a 3x3 scale matrix
func M3Scale(x, y, z float32) *M3 {
    m := new(M3)
    m.matrix = [3 * 3]float32{
        x, 0, 0,
        0, y, 0,
        0, 0, z}
    return m
}

// Create the product of two 3x3 matrices
func (a *M3) Compose(b *M3) *M3 {
    m := new(M3)
    for x := 0; x < 9; x += 3 {
        y, z := x+1, x+2
        m.matrix[x] = a.matrix[x]*b.matrix[0] +
            a.matrix[y]*b.matrix[3] +
            a.matrix[z]*b.matrix[6]
        m.matrix[y] = a.matrix[x]*b.matrix[1] +
            a.matrix[y]*b.matrix[4] +
            a.matrix[z]*b.matrix[7]
        m.matrix[z] = a.matrix[x]*b.matrix[2] +
            a.matrix[y]*b.matrix[5] +
            a.matrix[z]*b.matrix[8]
    }
    return m
}

// Transform a vector and return the result
func (m *M3) Mul(v V3) V3 {
    x := v.X*m.matrix[0] +
        v.Y*m.matrix[3] +
        v.Z*m.matrix[6]
    y := v.X*m.matrix[1] +
        v.Y*m.matrix[4] +
        v.Z*m.matrix[7]
    z := v.X*m.matrix[2] +
        v.Y*m.matrix[5] +
        v.Z*m.matrix[8]
    return V3{x, y, z}
}

// Create a 3x3 for rotation about the X-axis
func M3RotateX(radians float32) *M3 {
    m := new(M3)
    s, c := sin(radians), cos(radians)
    m.matrix = [3 * 3]float32{
        1, 0, 0,
        0, c, -s,
        0, s, c}
    return m
}

// Create a 3x3 for rotation about the Y-axis
func M3RotateY(radians float32) *M3 {
    m := new(M3)
    s, c := sin(radians), cos(radians)
    m.matrix = [3 * 3]float32{
        c, 0, s,
        0, 1, 0,
        -s, 0, c}
    return m
}

// Create a 3x3 for rotation about the Z-axis
func M3RotateZ(radians float32) *M3 {
    m := new(M3)
    s, c := sin(radians), cos(radians)
    m.matrix = [3 * 3]float32{
        c, s, 0,
        -s, c, 0,
        0, 0, 1}
    return m
}

// Create a duplicate of self
func (m *M3) Clone() *M3 {
    n := new(M3)
    for i := 0; i < 3*3; i += 1 {
        n.matrix[i] = m.matrix[i]
    }
    return n
}

func (self *M3) Transpose() *M3 {
    n := new(M3)
    m := &self.matrix
    n.matrix = [...]float32{
        m[0], m[3], m[6],
        m[1], m[4], m[7],
        m[2], m[5], m[8],
    }
    return n
}

// Get string representation to appease fmt.Printf
func (m *M3) String() string {
    x := m.matrix
    return fmt.Sprintf("%f %f %f\n"+
        "%f %f %f\n"+
        "%f %f %f\n",
        x[0], x[1], x[2],
        x[3], x[4], x[5],
        x[6], x[7], x[8])
}

func (a *M3) Equivalent(b *M3, ε float32) bool {
    for i, f := range b.matrix {
        if abs(a.matrix[i]-f) > ε {
            return false
        }
    }
    return true
}
