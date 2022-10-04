package main

import (
	"fmt"
	"net/http"
	"os"
)

// Stabilim numarul de chelneri in simulare
const WaiterNum = 4

// Stabilim numarul de mese in simulare
const TableNum = 10

var kitchenHost = "http://localhost"

// Crearea listelor pentru chelneri si mese
var waiterList = newWaiterList()
var tableList = newTableList()

func main() {
	args := os.Args
	if len(args) > 1 {
		//Stablim docker internal host
		kitchenHost = args[1]
	}

	fmt.Println("Dining hall is up and running!")

	// Activam chelnerii la lucru
	for _, waiter := range waiterList {
		go waiter.work()
	}
	// Unde e transmisa informatia
	http.HandleFunc("/delivery", deliveryHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
