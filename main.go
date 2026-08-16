package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"},
		{Name: "Venus Express", Prefix: "6000"},
		{Name: "Saturn Ring", Prefix: "7000"},
		{Name: "Jupiter Trust", Prefix: "8000"},
	}
	fmt.Println(DetectBank("", banks))

	fmt.Println(LuhnCheck("79927393398713"))

}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if len(cardNumber) == 0 {
		return nil
	}
	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}

	}
	return nil
}

func LuhnCheck(cardNumber string) bool {
	var sumNumber int
	if len(cardNumber) == 0 {
		return false
	}

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if digit > 9 {
			return false
		}
		if i%2 == len(cardNumber)%2 {
			if digit*2 > 9 {
				digit = digit*2 - 9
			} else {
				digit = digit * 2
			}

		}
		sumNumber += digit
	}
	return sumNumber%10 == 0
}
