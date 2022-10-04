package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var deliveryChan = make(chan *Delivery)

func deliveryHandler(w http.ResponseWriter, r *http.Request) {
	latestDelivery := new(Delivery)
	var requestBody = make([]byte, r.ContentLength)
	r.Body.Read(requestBody)
	json.Unmarshal(requestBody, latestDelivery)
	deliveryChan <- latestDelivery
	fmt.Fprint(w, "OK")
}
