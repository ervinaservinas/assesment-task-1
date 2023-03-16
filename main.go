package main

import (
	"fmt"
	"time"

	roman "github.com/ervinaservinas/assessments/roman"
)

func main() {
    var decimal int
    fmt.Print("Decimal number you wish to convert: ")
    fmt.Scan(&decimal) // reading input value

   romanValue, err := roman.RomanNumeral(decimal) // checking if the imput is good
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    time.Sleep(1 * time.Second)
    fmt.Println("Result of Roman numeral is", romanValue)
}