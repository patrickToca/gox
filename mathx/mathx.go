package mathx

import "math"

// AbsInt returns the absolute value of i.
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Round returns x rounded to the given unit.
// Tip: x is "arbitrary", maybe greater than 1.
// For example:
//     Round(0.363636, 0.001) // 0.364
//     Round(0.363636, 0.01)  // 0.36
//     Round(0.363636, 0.1)   // 0.4
//     Round(0.363636, 0.05)  // 0.35
//     Round(3.2, 1)          // 3
//     Round(32, 5)           // 30
//     Round(33, 5)           // 35
//     Round(32, 10)          // 30
//
// For details, see https://stackoverflow.com/a/39544897/1705598
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Near reports if 2 float64 numbers are "near" to each other.
// The caller is responsible to provide a sensible epsilon.
//
// "near" is defined as the following:
//     near := math.Abs(a - b) < eps
//
// Corner cases:
//  1. if a==b, result is true (eps will not be checked, may be NaN)
//  2. Inf is near to Inf (even if eps=NaN; consequence of 1.)
//  3. -Inf is near to -Inf (even if eps=NaN; consequence of 1.)
//  4. NaN is not near to anything (not even to NaN)
//  5. eps=Inf results in true (unless any of a or b is NaN)
func Near(a, b, eps float64) bool {
	// Quick check, also handles infinities:
	if a == b {
		return true
	}

	return math.Abs(a-b) < eps
}
