package testmain

import (
	"fmt"
	"testing"
)

func TestRomanNumeral(t *testing.T) {
    assert.Equal(t, "I", romanValue.RomanNumeral(1))
    assert.Equal(t, "IV", romanValue.RomanNumeral(4))
    assert.Equal(t, "IX", romanValue.RomanNumeral(9))
    assert.Equal(t, "XIV", romanValue.RomanNumeral(14))
    assert.Equal(t, "XXXIX", romanValue.RomanNumeral(39))
    assert.Equal(t, "XLIV", romanValue.RomanNumeral(44))
    assert.Equal(t, "XCIX", romanValue.RomanNumeral(99))
    assert.Equal(t, "CDXLIV", romanValue.RomanNumeral(444))
    assert.Equal(t, "CMXCIX", romanValue.RomanNumeral(999))
    assert.Equal(t, "MMXXI", romanValue.RomanNumeral(2021))
    
    fmt.Println("Tests done")
}
