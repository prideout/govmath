package vmath

import "math"

// TODO do something faster than this!

func cos(f float32) float32  { return float32(math.Cos(float64(f))) }
func sin(f float32) float32  { return float32(math.Sin(float64(f))) }
func sqrt(f float32) float32 { return float32(math.Sqrt(float64(f))) }
func abs(f float32) float32  { return float32(math.Abs(float64(f))) }
