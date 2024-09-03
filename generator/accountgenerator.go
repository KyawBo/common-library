package generator

import (
	"fmt"
	"strconv"
)

const (
	initialAccountNumber = "000000000000"
	checkDigitLimit      = 9
	maxIncrementNumber   = 9000000000
)

func GenerateAccountNumberWithInitialAccount(prefixNumber int) (string, error) {
	return GenerateAccountNumber(prefixNumber, "")
}

func GenerateAccountNumber(prefixNumber int, accountNumberStr string) (string, error) {
	var digits []int

	if accountNumberStr == "" {
		accountNumberStr = initialAccountNumber
	}

	extractedNumberStr := accountNumberStr[1:11]
	incrementalNumber, err := strconv.Atoi(extractedNumberStr)
	if err != nil {
		return "", fmt.Errorf("failed to convert extracted number: %v", err)
	}
	incrementalNumber++
	for {
		digits = []int{prefixNumber}

		if incrementalNumber > maxIncrementNumber {
			return "", fmt.Errorf("incremental number exceeds maximum allowed value of %d", maxIncrementNumber)
		}

		// Convert incrementNumber to a string and append its digits to the digits slice
		formattedNumberStr := fmt.Sprintf("%010d", incrementalNumber)
		for _, char := range formattedNumberStr {
			digit := int(char - '0')
			digits = append(digits, digit)
		}

		// Calculate the check digit with digits slice
		checkDigit := calculateCheckDigit(digits)

		// If check digit is valid (0-9), break the loop
		if checkDigit <= checkDigitLimit {
			digits = append(digits, checkDigit)
			break
		}

		// Increment the incremental number and try again
		incrementalNumber++
	}

	var latestAccountNumber string
	for _, digit := range digits {
		latestAccountNumber += strconv.Itoa(digit)
	}

	return latestAccountNumber, nil
}

func calculateCheckDigit(accountNumber []int) int {
	// Weights for each digit
	weights := []int{12, 1, 10, 3, 8, 5, 6, 7, 4, 9, 2}

	sum := 0
	for i := 0; i < 11; i++ {
		sum += accountNumber[i] * weights[i]
	}

	checkDigit := sum % 11

	return checkDigit
}
