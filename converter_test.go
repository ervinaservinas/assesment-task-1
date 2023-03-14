package main

import (
	"testing"

	"github.com/ervinaservinas/assessments/converter"
)

func TestRomanNumeral(t *testing.T) {
	testCases := []struct {
		name     string
		decimal  int
		expected string
	}{
        {"Negative numbers should return an empty string", -1, ""},
        {"Numbers with 3 should return error", -3, ""},
		{"Numbers with 7 should return error", -7, ""},
		{"Numbers with 8 should return error", -8, ""},
		{"1  to I", 1, "I"},
		{"4  to IV", 4, "IV"},
		{"9  to IX", 9, "IX"},
		{"10  to X", 10, "X"},
		{"40  to XL", 40, "XL"},
		{"50  to L", 50, "L"},
		{"90  to XC", 90, "XC"},
		{"100  to C", 100, "C"},
		{"400  to CD", 400, "CD"},
		{"500  to D", 500, "D"},
		{"900  to CM", 900, "CM"},
		{"1000  to M", 1000, "M"},
		{"1984  to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 to MMMCMXCIX", 3999, "MMMCMXCIX"},
	}
	for _, testCase := range testCases {  // taking range from testCases
		t.Run(testCase.name, func(t *testing.T) {  // testing each test range
			actual := converter.RomanNumeral(testCase.decimal)
			if actual != testCase.expected {
				t.Errorf("Expected Roman numeral for %d to be %s, but got %s", testCase.decimal, testCase.expected, actual)
			}
		})
	}
}