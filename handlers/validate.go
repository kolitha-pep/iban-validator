package handlers

import (
	"github.com/gin-gonic/gin"
	"iban-validator/pkg"
)

type iban struct {
	IBAN string `json:"iban" binding:"required"`
}

type ibanResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func (i ibanResponse) Error() string {
	return i.Message
}

func ValidateIBAN(c *gin.Context) {
	in := iban{}
	if err := c.BindJSON(&in); err != nil {
		responseObject(c, in, err)
		return
	}

	valid, err := pkg.Validate(in.IBAN)

	// if invalid, return error
	if !valid && err != nil {
		responseObject(c, ibanResponse{
			Valid:   false,
			Message: err.Error(),
		}, nil)
		return
	}

	// if valid, return success
	responseObject(c, ibanResponse{
		Valid:   true,
		Message: in.IBAN,
	}, nil)
	return
}
