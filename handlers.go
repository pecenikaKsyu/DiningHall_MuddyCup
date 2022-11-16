package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"bytes"
	"io"
	"log"
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


func GetFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonFoods, err := json.Marshal(menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonFoods)
}

func ServeOrder(w http.ResponseWriter, r *http.Request) {
	var cookedOrder Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cookedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if cookedOrder.WaiterId == -1 {
		OnlineCookedOrder = append(OnlineCookedOrder, cookedOrder)
	} else {
		latestDelivery := new(Delivery)
		var requestBody = make([]byte, r.ContentLength)
		r.Body.Read(requestBody)
		json.Unmarshal(requestBody, latestDelivery)
		deliveryChan <- latestDelivery
		fmt.Fprint(w, "OK")

	}

	defer r.Body.Close()
	jsonCookedOrder, _ := json.Marshal(cookedOrder)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCookedOrder)
}

func ReceiveOnlineOrder(w http.ResponseWriter, r *http.Request) {
	var receivedOrder OnlineReceivedOrder
	var responseOrder OnlineResponseOrder
	var kitchenInform KitchenInfo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receivedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	receivedOrder.Id = AiOrder.ID()

	//sending the order immediately to the kitchen
	kitchenInform = SendOnlineOrder(receivedOrder)

	//calculating estimating time
	estimatedTime := CalculateEstimatedTime(receivedOrder, kitchenInform)

	//responseOrder.RestaurantId = RestaurantId
	responseOrder.OrderId = receivedOrder.Id
	responseOrder.EstimatedWaitingTime = estimatedTime
	responseOrder.CreatedTime = receivedOrder.CreatedTime
	responseOrder.RegisteredTime = time.Now()

	defer r.Body.Close()
	jsonCookedOrder, _ := json.Marshal(responseOrder)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCookedOrder)
}

func CalculateEstimatedTime(currentOrder OnlineReceivedOrder, kitchenInform KitchenInfo) int {
	var A, B, C, D, E, F, estimatedTime int
	foods := Menu[] 

	for _, i := range currentOrder.Items {
		if foods[i-1].CookingApparatus == "" {
			A += foods[i-1].PreparationTime
		} else {
			C += foods[i-1].PreparationTime
		}
	}

	//Here i should use the variable from config and json config
	B = kitchenInform.CooksProfficiency
	D = kitchenInform.CookingApparatus
	E = kitchenInform.NrFoodsQueue
	F = len(foods)

	estimatedTime = (A/B + C/D) * (E + F) / F
	return estimatedTime

}

func SendOnlineOrder(receivedOrder OnlineReceivedOrder) KitchenInfo {
	reqBody, err := json.Marshal(receivedOrder)
	if err != nil {
		log.Fatal(err.Error())

	}
	//Sending the order to kitchen
	resp, err := http.Post(kitchenHost+":8000/onlineOrder", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal("Request Failed: %s", err.Error())
	}
	defer resp.Body.Close()

	//receiving info from kitchen regarding foodQueue, apparatus, cooks profficiency
	var kitchenInform KitchenInfo

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Response from Kitchen Failed: %s", err.Error())
		//return kitchenInform
	}
	_ = json.Unmarshal(body, &kitchenInform)

	log.Printf("---------The ONLINE order with id %d was SENT to Kitchen. Details: %+v\n", receivedOrder.Id, receivedOrder) // Unmarshal result
	return kitchenInform
}
