package roman

import (
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	testCases := []struct {
		name     string
		decimal  int
		expected string
	}{
        {"Negative numbers", -1, ""},
        {"Negative numbers", -3, ""},
		{"Negative numbers", -7, ""},
		{"Negative numbers", -8, ""},
		{"1  to I", 1, "I"},
		{"2 to II", 2, "II"},
		{"3 to III", 3, "III"},
		{"4  to IV", 4, "IV"},
		{"5 to V", 5, "V"},
		{"6 to VI", 6, "VI"},
		{"7 to VII", 7, "VII"},
		{"8 to VIII", 8, "VIII"},
		{"9  to IX", 9, "IX"},
		{"10  to X", 10, "X"},
		{"14 to XIV", 14, "XIV"},
		{"15 to XV", 15, "XV"},
		{"16 to XVI", 16, "XVI"},
		{"39 to XXXIX", 39, "XXXIX"},
		{"40  to XL", 40, "XL"},
		{"49 to XLIX", 49, "XLIX"},
		{"50  to L", 50, "L"},
		{"90  to XC", 90, "XC"},
		{"99 to XCIX", 99, "XCIX"},
		{"100  to C", 100, "C"},
		{"400  to CD", 400, "CD"},
		{"500  to D", 500, "D"},
		{"900  to CM", 900, "CM"},
		{"1000  to M", 1000, "M"},
		{"4000 to empty string", 4000, ""},
		{"1984  to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"0 to empty string", 0, ""},
		{"1000000 to empty string", 1000000, ""},
		{"4999 to empty string", 4999, ""},
		{"-4000 to empty string", -4000, ""},
		{"-1000000 to empty string", -1000000, ""},
		{"-4999 to empty string", -4999, ""},
	}

	for _, testCase := range testCases {
		actual, _ := RomanNumeral(testCase.decimal)
		if actual != testCase.expected {
			t.Errorf("Test case %s failed: expected %s but got %s", testCase.name, testCase.expected, actual)
		}
	}
}