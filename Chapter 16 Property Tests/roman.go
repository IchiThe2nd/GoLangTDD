package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

//Adding user input to get away from direct copying for a second.

func main() {
	for i := 10; i > 0; i-- {
		number := "23"
		//reader := bufio.NewReader(os.Stdin)
		if i == 1 {
			fmt.Println("ENOUGH!")

		}
		fmt.Println("Enter Arabic Number..")
		reader := bufio.NewReader(os.Stdin)

		number, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		number = strings.TrimSpace(number)
		if number == "" {
			return
		}

		{
			newint, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ConvertToRoman(newint))
		}
	}
}
