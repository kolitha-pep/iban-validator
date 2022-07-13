package handlers

import (
	"github.com/gin-gonic/gin"
	"iban-validator/pkg"
)

type iban struct {
	IBAN string `json:"iban"`
}

func IBANValidator(c *gin.Context) {
	in := iban{}
	if err := c.BindJSON(&in); err != nil {
		responseObject(c, in, err)
		return
	}

	_, err := pkg.Validate(in.IBAN)

	if err != nil {
		responseObject(c, in, err)
		return
	}
	responseObject(c, "valid IBAN", nil)
	return
}
