# IBAN API

IBAN number is a unique identification number for an bank account. It is used in international payments.
It is a combination of country code and account number. The country code is a two-letter code that identifies the country.
The account number is a unique number that identifies the account.

## <b>Algorithm </b>
1. Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid.
2. Move the four initial characters to the end of the string.
3. Replace each letter in the string with two digits, thereby expanding the string, where A or a = 10, B or b = 11, ..., Z or z = 35.
4. Interpret the string as a decimal integer and compute the remainder of that number on division by 97

If the remainder is 1, the check digit test is passed and the IBAN might be valid.

## <b>Build</b>

### <b>Makefile</b>
Compiling for Windows, Linux and Darwin at once.
```
make compile
```
Build binary based on local OS.

```
make build

# build and run binary
make build_then_run
```

## <b>Tests</b>

Run the unit tests using make.

```
make test
```

## <b>Usage</b>

```
GET 0.0.0.0:8080/validate/iban
{
    "iban": "AD1200012030200359100100"
}
```
Header
```
Content-Type: application/json
```

Success Response - 200
```
{
    "valid": true,
    "message": "AD1200012030200359100100"
}
```
```
{
    "valid": false,
    "message": "IBAN: Invalid IBAN country code <LK>"
}
```

Error Response - 400
```
{
    "message": "Key: 'iban.IBAN' Error:Field validation for 'IBAN' failed on the 'required' tag",
    "code": 0,
    "level": 0
}
```
