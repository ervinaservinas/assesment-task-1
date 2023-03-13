package main

import (
	"fmt"
	"time"
)


func RomanNumeral(decimal int) string {
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    romanValue := ""

	// for loop which goes though all lenght of symbols
    for i := 0; i < len(symbols); i++ {
        for decimal >= values[i] { // checking if decimal value is greater then values of "values" 
            decimal -= values[i]  // assigning the result to decimals (tried till it worked)
            romanValue += symbols[i]  // adding symbol as Roman number.
        }
    }
    return romanValue
}

func main() {
    var decimal int
    fmt.Print("Decimal number you wish to convert: ")
    fmt.Scan(&decimal) // reading input value
	time.Sleep(1 * time.Second) // one second delay for the answer
    fmt.Println("Result of Roman numeral is ", RomanNumeral(decimal)) // printing roman numeral
}