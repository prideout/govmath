package vmath

import "fmt"

// Implements a 4x3 matrix type for 3D graphics. (four rows, three columns)
// This is useful for affine transformations only.
type T3 struct {
    matrix [4 * 3]float32
}

// Create a 4x3 from the identity
func T3Identity() *T3 {
    m := new(T3)
    m.matrix = [4 * 3]float32{
        1, 0, 0,
        0, 1, 0,
        0, 0, 1,
        0, 0, 0}
    return m
}

// Create a 4x3 translation matrix
func T3Translate(x, y, z float32) *T3 {
    m := new(T3)
    m.matrix = [4 * 3]float32{
        1, 0, 0,
        0, 1, 0,
        0, 0, 1,
        x, y, z}
    return m
}

// Create a 4x3 scale matrix
func T3Scale(x, y, z float32) *T3 {
    m := new(T3)
    m.matrix = [4 * 3]float32{
        x, 0, 0,
        0, y, 0,
        0, 0, z,
        0, 0, 0}
    return m
}

// Transform a P3, which is like V4 with w=1
// Another way of thinking about it: 1x4 * 4x3 = 1x3
func (matrix *T3) Transform(p P3) P3 {
    m := &matrix.matrix
    c0 := V3New(m[0], m[3], m[6])
    c1 := V3New(m[1], m[4], m[7])
    c2 := V3New(m[2], m[5], m[8])
    v := V3New(p.X, p.Y, p.Z)
    x := c0.Dot(v) + m[9]
    y := c1.Dot(v) + m[10]
    z := c2.Dot(v) + m[11]
    return P3New(x, y, z)
}

func (matrix *T3) Mul(v V3) V4 {
    m := &matrix.matrix
    c0 := V3New(m[0], m[1], m[2])
    c1 := V3New(m[3], m[4], m[5])
    c2 := V3New(m[6], m[7], m[8])
    c3 := V3New(m[9], m[10], m[11])
    x := c0.Dot(v)
    y := c1.Dot(v)
    z := c2.Dot(v)
    w := c3.Dot(v)
    return V4New(x, y, z, w)
}

// Compose two 4x3 matrices
// Technically this isn't matrix multiplication.
func (a *T3) Compose(b *T3) *T3 {
    m := new(T3)
    for x := 0; x < 12; x += 3 {
        y, z := x+1, x+2
        m.matrix[x] = 0 +
            a.matrix[x]*b.matrix[0] +
            a.matrix[y]*b.matrix[3] +
            a.matrix[z]*b.matrix[6] +
            b.matrix[9]
        m.matrix[y] = 0 +
            a.matrix[x]*b.matrix[1] +
            a.matrix[y]*b.matrix[4] +
            a.matrix[z]*b.matrix[7] +
            b.matrix[10]
        m.matrix[z] = 0 +
            a.matrix[x]*b.matrix[2] +
            a.matrix[y]*b.matrix[5] +
            a.matrix[z]*b.matrix[8] +
            b.matrix[11]
    }
    return m
}

// Create a 4x3 for rotation about the X-axis
func T3RotateX(radians float32) *T3 {
    m := new(T3)
    s, c := sin(radians), cos(radians)
    m.matrix = [4 * 3]float32{
        1, 0, 0,
        0, c, -s,
        0, s, c,
        0, 0, 0}
    return m
}

// Create a 4x3 for rotation about the Y-axis
func T3RotateY(radians float32) *T3 {
    m := new(T3)
    s, c := sin(radians), cos(radians)
    m.matrix = [4 * 3]float32{
        c, 0, s, 
        0, 1, 0,
        -s, 0, c,
        0, 0, 0}
    return m
}

// Create a 4x3 for rotation about the Z-axis
func T3RotateZ(radians float32) *T3 {
    m := new(T3)
    s, c := sin(radians), cos(radians)
    m.matrix = [4 * 3]float32{
        c, s, 0,
        -s, c, 0,
        0, 0, 1,
        0, 0, 0}
    return m
}

// Create a duplicate of self
func (m *T3) Clone() *T3 {
    n := new(T3)
    for i := 0; i < 4*3; i++ {
        n.matrix[i] = m.matrix[i]
    }
    return n
}

// Return a M3 object for the upper-left portion
func (m *T3) GetUpperLeft() *M3 {
    n := new(M3)
    x := &m.matrix
    n.matrix = [3 * 3]float32{ // this would be easier with slices...
        x[0], x[1], x[2],
        x[3], x[4], x[5],
        x[6], x[7], x[8]}
    return n
}

// Return last row of matrix
func (m *T3) GetTranslation() V3 {
    x := &m.matrix
    return V3New(x[9], x[10], x[11])
}

// Get string representation to appease fmt.Printf
func (m *T3) String() string {
    x := &m.matrix
    return fmt.Sprintf("%f %f %f\n"+
        "%f %f %f\n"+
        "%f %f %f\n"+
        "%f %f %f\n",
        x[0], x[1], x[2],
		x[3], x[4], x[5],
		x[6], x[7], x[8],
		x[9], x[10], x[11])
}
