package main

import (
	"fmt"
	"net/http"
	"os"
)

var kitchenHost = "http://localhost"

func main() {

	args := os.Args

	if len(args) > 1 {
		//Set the docker internal host
		kitchenHost = args[1]
	}

	fmt.Println("Dining hall is up and running!")

	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/delivery", deliveryHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
