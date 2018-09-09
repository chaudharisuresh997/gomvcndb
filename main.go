package main

import (
	c "control"
	domain "domain"
	"fmt"
	"log"
	"net/http"
)

func main() {
	domain.CustomFun()
	router := c.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Println("Shiv")
}
