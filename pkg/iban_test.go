package pkg

import "testing"

var validIBANTestNumbers = []struct {
	number string
}{
	{"LU28 0019 4006 4475 0000"},
	{"LU28 0019 4006 4475 0000"},
	{"LU28 0019 4006 4475 0000"},
	{"LU28 0019 4006 4475 0000"},
	{"LU28 0019 4006 4475 0000"},
}

var invalidIBANTestNumbers = []struct {
	number string
}{
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
	{"LU12 3456 7890 1234 5678"},
}

func TestValidIBAN(t *testing.T) {
	for _, ibanTestNumber := range validIBANTestNumbers {
		valid, err := Validate(ibanTestNumber.number)
		if err != nil || !valid {
			t.Error("No object was created!")
			t.Log(err)
		}
	}
}

func TestInvalidIBAN(t *testing.T) {
	for _, ibanTestNumber := range invalidIBANTestNumbers {
		_, err := Validate(ibanTestNumber.number)
		if err == nil {
			t.Error("No error was thrown for an invalid IBAN number!")
		}
	}
}
