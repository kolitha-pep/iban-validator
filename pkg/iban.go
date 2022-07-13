package pkg

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	// Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
	// Move the four initial characters to the end of the string
	// Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
	// Interpret the string as a decimal integer and compute the remainder of that number on division by 97

	_, err := isValidIban("AD1200012030200359100100")
	fmt.Println(r)
	io.WriteString(w, err.Error())
}

func isValidIban(iban string) (isValid bool, err error) {
	// Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
	// Move the four initial characters to the end of the string
	// Replace each letter in the string with two digits, thereby expanding the string, where A or a = 10, B or b = 11, ..., Z or z = 35
	// Interpret the string as a decimal integer and compute the remainder of that number on division by 97

	// check for min iban length
	if len(iban) < 15 {
		return false, fmt.Errorf("IBAN: Invalid IBAN string <%s>", iban)
	}

	// convert to uppercase and remove spaces
	iban = strings.ToUpper(strings.Replace(iban, " ", "", -1))

	// extract iban country code
	ibcc := iban[0:2]

	// check if country code is valid
	cc, ok := CountryIBAN[ibcc]
	if !ok {
		return false, fmt.Errorf("IBAN: Invalid IBAN country code <%s>", ibcc)
	}

	// validate iban length based on the country
	if len(iban) != cc.Length {
		return false, fmt.Errorf("IBAN: Invalid IBAN length <%s>", iban)
	}

	// Move the four initial characters to the end of the string
	modifiedIban := iban[4:] + iban[:4]

	fmt.Println(cc, modifiedIban)

	// check if characters are A-z
	regx := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	var intStr string
	for _, c := range modifiedIban {
		if regx(string(c)) {
			c -= 55
			intStr += strconv.Itoa(int(c))
			continue
		}
		intStr += string(c)
	}

	n := new(big.Int)
	n, ok = n.SetString(intStr, 10)
	if !ok {
		return false, errors.New("SetString error")
	}

	divisor := new(big.Int).SetInt64(97)
	remainder := new(big.Int).Mod(n, divisor)

	if remainder.Int64() == 1 {
		return true, nil
	}
	return false, errors.New("iban has incorrect check digits")
}

// 00012030200359100100101312
// 00012030200359100100101312
