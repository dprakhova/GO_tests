package main

import "testing"

func TestCalculateTax(t *testing.T) {
	testCases := []struct {
		salary            float64
		expectedTax       float64
		expectedRemainder float64
	}{
		{1000, 50, 950}, {1500, 90, 1410}, {3000, 210, 2790}, {5000, 410, 4590}, {7000, 810, 6190}}
	for _, tc := range testCases {
		tax, remainder := calculateTax(tc.salary)
		if tax != tc.expectedTax || remainder != tc.expectedRemainder {
			t.Errorf("calculateTax(%f) = (%f, %f), expected (%f, %f)", tc.salary, tax, remainder, tc.expectedTax, tc.expectedRemainder)
		}
	}
}
