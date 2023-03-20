package roman

import (
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	testCases := []struct {
		decimal        int
		expected       string
		expectError    bool
	}{
		{-1, "", true},
		{-3, "", true},
		{-7, "", true},
		{-8, "", true},
		{1, "I", false},
		{2, "II", false},
		{3, "III", false},
		{4, "IV", false},
		{5, "V", false},
		{6, "VI", false},
		{7, "VII", false},
		{8, "VIII", false},
		{9, "IX", false},
		{10, "X", false},
		{14, "XIV", false},
		{15, "XV", false},
		{16, "XVI", false},
		{39, "XXXIX", false},
		{40, "XL", false},
		{49, "XLIX", false},
		{50, "L", false},
		{90, "XC", false},
		{99, "XCIX", false},
		{100, "C", false},
		{400, "CD", false},
		{500, "D", false},
		{900, "CM", false},
		{1000, "M", false},
		{4000, "", true},
		{1984, "MCMLXXXIV", false},
		{3999, "MMMCMXCIX", false},
		{0, "", true},
		{1000000, "", true},
		{4999, "", true},
		{-4000, "", true},
		{-1000000, "", true},
		{-4999, "", true},
	}

	for _, testCase := range testCases {
		actual, err := RomanNumeral(testCase.decimal)
		if (err != nil && !testCase.expectError) || (err == nil && testCase.expectError) {
			t.Errorf("Test case failed: expected error: %t, actual error: %v", testCase.expectError, err)
		} else if actual != testCase.expected {
			t.Errorf("Test case failed: input: %d, expected: %s, actual: %s", testCase.decimal, testCase.expected, actual)
		}
	}
}
