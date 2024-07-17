package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harishkokcha91/GoToDO/router"
)

func main() {

	r := router.Router()
	fmt.Println("STarting the server on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
