package converter

func RomanNumeral(decimal int) string {

	numerals := []struct { // I found struct
		Symbol string
		Value  int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	romanValue := ""
	for _, numeral := range numerals { // as from documentation understoon _, is blank identifier
		for decimal >= numeral.Value {
			decimal -= numeral.Value
			romanValue += numeral.Symbol
		}
	}
	return romanValue
}
