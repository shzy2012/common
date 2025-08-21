package tools

import "math"

func Round(num float64, digits int) float64 {
	multiplier := math.Pow(10, float64(digits))
	return math.Round(num*multiplier) / multiplier
}
