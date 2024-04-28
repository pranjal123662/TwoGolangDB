package main

import (
	"TwoDB/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":7456", r))
}
