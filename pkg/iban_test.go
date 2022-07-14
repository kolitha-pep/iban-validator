package pkg

import "testing"

var validIBANNumbers = []struct {
	number string
}{
	{"LU28 0019 4006 4475 0000"},
	{"AL35202111090000000001234567"},
	{"AD1400080001001234567890"},
	{"AT483200000012345864"},
	{"AZ77VTBA00000000001234567890"},
	{"DO22ACAU00000000000123456789"},
	{"FO9264600123456789"},
	{"JO71CBJO0000000000001234567890"},
	{"RS35105008123123123173"},
}

var invalidIBANNumbers = []struct {
	number string
}{
	{"LU12 3456 7890 1234 5678"},
	{"世界 3456 7890 1234 5678"},
	{"RS12 3456 7890 1234 5678"},
	{"!@#$% 1234 5678"},
}

func TestValidIBAN(t *testing.T) {
	for _, ibanTestNumber := range validIBANNumbers {
		valid, err := Validate(ibanTestNumber.number)
		if err != nil || !valid {
			t.Error("No object was created!")
			t.Log(err)
		}
	}
}

func TestInvalidIBAN(t *testing.T) {
	for _, ibanTestNumber := range invalidIBANNumbers {
		_, err := Validate(ibanTestNumber.number)
		if err == nil {
			t.Error("No error was thrown for an invalid IBAN number!")
		}
	}
}
