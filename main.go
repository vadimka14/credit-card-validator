package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func main() {
	var myCard string = "40007393398713"

	banks, err := loadBankData("bank.txt")
	if err != nil {
		log.Fatal("Не удалось загрузить банки: open banks.txt: no such file or directory")
	} else {
		fmt.Printf("Загружено банков: %d\n", len(banks))
	}

	fmt.Printf("Валиден по Луне: %t\n", LuhnCheck(myCard))

	bank := DetectBank(myCard, banks)
	if bank != nil {
		fmt.Printf("Банк: %s\n", bank.Name)
	} else {
		fmt.Println("Банк: не определён")
	}

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

func loadBankData(path string) ([]Bank, error) {
	var banks []Bank
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("неверный формат строки: %q", line)
		}
		banks = append(banks, Bank{
			Name:   strings.TrimSpace(parts[0]),
			Prefix: strings.TrimSpace(parts[1]),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, nil
}
