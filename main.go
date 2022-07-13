package main

import (
	"github.com/gin-gonic/gin"
	"iban-validator/handlers"
)

func main() {

	//http.HandleFunc("/validate", pkg.Validate)
	//
	//err := http.ListenAndServe(":3030", nil)
	//
	//if errors.Is(err, http.ErrServerClosed) {
	//	fmt.Printf("server closed\n")
	//} else if err != nil {
	//	fmt.Printf("error starting server: %s\n", err)
	//	os.Exit(1)
	//}
	r := gin.Default()
	r.GET("/validate", handlers.IBANValidator)
	r.Run()
}
