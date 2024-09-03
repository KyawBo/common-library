package generator

import (
	"fmt"
	"testing"
)

func TestGenerateAccountNumberSuccess(t *testing.T) {
	var tests = []struct {
		name          string
		prefixNumber  int
		accountNumber string
		want          string
	}{
		{"Generate second account", 3, "300000000015", "300000000027"},
		{"Generate third account", 3, "300000000027", "300000000039"},
		{"Generate fourth account", 3, "300000000039", "300000000040"},
		{"Generate account with recalculate", 3, "300000018058", "300000018071"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateAccountNumber(tt.prefixNumber, tt.accountNumber)
			if err != nil {
				t.Errorf("Got error while generating account number: %v", err)
			}
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
			if tt.accountNumber == result {
				t.Errorf("Got wrong account number, existing and latest account number are the same value")
			}
		})
	}

}

func TestGenerateAccountNumberError(t *testing.T) {
	var tests = []struct {
		name          string
		prefixNumber  int
		accountNumber string
		want          string
	}{
		{"Number of account exceeds limit", 3, "390000000015", ""},
		{"Failed to convert extracted number", 3, "3AV000000027", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateAccountNumber(tt.prefixNumber, tt.accountNumber)
			fmt.Println(err)
			if err == nil {
				t.Errorf("Should get error while generating account number: %v", err)
			}
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}

		})
	}

}

func TestGenerateAccountNumberWithInitialAccountSuccess(t *testing.T) {
	prefixNumber := 3
	expectedResult := "300000000015"

	t.Run("Generate first ever account", func(t *testing.T) {
		result, err := GenerateAccountNumberWithInitialAccount(prefixNumber)
		if err != nil {
			t.Errorf("Got error while generating account number: %v", err)
		}
		if result != expectedResult {
			t.Errorf("got %s, want %s", result, expectedResult)
		}
	})
}
