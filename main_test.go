package main

import (
	"math"
	"testing"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixedPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func Test_expectedCastsForGain(t *testing.T) {
	testCases := map[string]struct {
		currentSkill float64
		circle       float64
		expected     float64
	}{
		"80.0 skill, 7th circle": {80.0, 6.0, 5.064599},
		"80.0 skill, 8th circle": {80.0, 7.0, 5.0},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := toFixedPrecision(expectedCastsForGain(tc.currentSkill, tc.circle), 6)
			if actual != tc.expected {
				t.Errorf("%.10f != %.10f", actual, tc.expected)
			}
		})
	}
}
