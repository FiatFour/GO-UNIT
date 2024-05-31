package main

import "testing"

func TestNumber(t *testing.T) {
	testCaseAdds := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCaseAdds {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			expectedResult := tc.expected

			if result != expectedResult {
				t.Errorf("Add(%d, %d) = %d is wrong, correct is %d", tc.a, tc.b, result, expectedResult)
			}
		})
	}

	testCaseFactorials := []struct {
		name     string
		num      int
		expected int
	}{
		{"Case 2", 2, 2},
		{"Case 5", 5, 120},
		{"Case -1", -1, 0},
	}

	for _, tc := range testCaseFactorials {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)
			expectedResult := tc.expected

			if result != expectedResult {
				t.Errorf("Factorial(%d) = %d is wrong, correct is %d", tc.num, result, expectedResult)
			}
		})
	}

	// result := Add(2, 3)
	// expectResult := 5
}
