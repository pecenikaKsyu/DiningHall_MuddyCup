package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func sendHandler(w http.ResponseWriter, r *http.Request) {
	order := Order{
		OrderId:    rand.Intn(6),
		TableId:    rand.Intn(6),
		WaiterId:   rand.Intn(2),
		Items:      getItems(),
		Priority:   rand.Intn(3),
		MaxWait:    rand.Intn(200) + 40,
		PickUpTime: time.Now().Unix(),
	}

	marshaledOrder, _ := json.Marshal(order)
	request, _ := http.NewRequest(http.MethodPost, kitchenHost+":8000/order", bytes.NewBuffer(marshaledOrder))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Fprintln(w, "ERROR:", err)
	} else {

		fmt.Fprintln(w, "Sent:"+string(marshaledOrder))

		var responseBuffer = make([]byte, response.ContentLength)

		response.Body.Read(responseBuffer)

		fmt.Fprintln(w, "Received:"+string(responseBuffer))
	}
}

func deliveryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
