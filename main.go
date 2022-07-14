package main

import (
	"github.com/gin-gonic/gin"
	"iban-validator/handlers"
)

func main() {

	r := gin.Default()
	r.GET("/validate/iban", handlers.ValidateIBAN)
	r.Run()
}
