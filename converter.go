package main

import (
	"fmt"
	"time"

	"github.com/ervinaservinas/assessments/converter"
)

func main() {
    var decimal int
    fmt.Print("Decimal number you wish to convert: ")
    fmt.Scan(&decimal) // reading input value

    if decimal < 0 || decimal > 3999 {  // added simple check for input.
		fmt.Println("Error! Roman numbers should be from 0 to 3999.")
		return
	}

    time.Sleep(1 * time.Second) // one second delay for the answer
    fmt.Println("Result of Roman numeral is ", converter.RomanNumeral(decimal)) // printing roman numeral
}