package pkg

type Country struct {
	Length int // length of IBAN code
}

var CountryIBAN = map[string]Country{
	"AD": {Length: 24},
	"AE": {Length: 23},
	"AL": {Length: 28},
	"AT": {Length: 20},
	"AZ": {Length: 28},
	"BA": {Length: 20},
	"BE": {Length: 16},
	"BG": {Length: 22},
	"BH": {Length: 22},
	"BR": {Length: 29},
	"CH": {Length: 21},
	"CR": {Length: 21},
	"CY": {Length: 28},
	"CZ": {Length: 24},
	"DE": {Length: 22},
	"DK": {Length: 18},
	"DO": {Length: 28},
	"EE": {Length: 20},
	"ES": {Length: 24},
	"FI": {Length: 18},
	"FO": {Length: 18},
	"FR": {Length: 27},
	"GB": {Length: 22},
	"GE": {Length: 22},
	"GI": {Length: 23},
	"GL": {Length: 18},
	"GR": {Length: 27},
	"GT": {Length: 28},
	"HR": {Length: 21},
	"HU": {Length: 28},
	"IE": {Length: 22},
	"IL": {Length: 23},
	"IS": {Length: 26},
	"IT": {Length: 27},
	"JO": {Length: 30},
	"KW": {Length: 30},
	"KZ": {Length: 20},
	"LB": {Length: 28},
	"LC": {Length: 32},
	"LI": {Length: 21},
	"LT": {Length: 20},
	"LU": {Length: 20},
	"LV": {Length: 21},
	"MC": {Length: 27},
	"MD": {Length: 24},
	"ME": {Length: 22},
	"MG": {Length: 27},
	"MK": {Length: 19},
	"MR": {Length: 27},
	"MT": {Length: 31},
	"MU": {Length: 30},
	"NL": {Length: 18},
	"NO": {Length: 15},
	"PK": {Length: 24},
	"PL": {Length: 28},
	"PS": {Length: 29},
	"PT": {Length: 25},
	"QA": {Length: 29},
	"RO": {Length: 24},
	"RS": {Length: 22},
	"SA": {Length: 24},
	"SC": {Length: 31},
	"SE": {Length: 24},
	"SI": {Length: 19},
	"SK": {Length: 24},
	"SM": {Length: 27},
	"ST": {Length: 25},
	"TL": {Length: 23},
	"TN": {Length: 24},
	"TR": {Length: 26},
	"UA": {Length: 29},
	"VG": {Length: 24},
	"XK": {Length: 20},
}
