package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

var waiterID = 0

var avg = 0.00
var summ = 0
var counter = 0

type Waiter struct {
	ID int
}

// cream o lista de chelneri
func newWaiterList() []Waiter {
	var ret []Waiter
	for i := 0; i < WaiterNum; i++ {
		ret = append(ret, Waiter{waiterID})
		waiterID += 1
	}
	return ret
}

// cum lucreaza un chelner
func (w Waiter) work() {
	for {
		tableList.mx.Lock() // la determinarea accesibiltatii meselor au acces chelnerii cate unul pentru a nu incurca accesibilitatea acestora
		for i := range tableList.tableArr {
			table := tableList.tableArr[i]
			if table.readyToOrder == true { //daca e accesibil catre comanda
				tableList.tableArr[i].readyToOrder = false // schimbam valoarea accesibilitatii
				order := table.generateOrder(w)            // generam random o comanda
				sendOrder(order)                           //trimitem comanda
				// afisam comanda
				fmt.Println("Order", order.OrderId, "was sent from table:", table.ID)

				break
			}
		}
		tableList.mx.Unlock() // deschidem accesul catre urmatoarea masa

		select {
		case delivery := <-deliveryChan: //transmitem comanda realizata
			tableList.mx.Lock()
			tableList.tableArr[delivery.TableID].readyToOrder = true // modificam accesibilitatea
			now := time.Now().Unix()

			// calculam recenzia comenzii
			rating := 0
			maxWaitF := float64(delivery.MaxWait)
			timeWaitedF := float64(now - delivery.PickUpTime)

			if maxWaitF*1.4 >= timeWaitedF {
				rating += 1
			}
			if maxWaitF*1.3 >= timeWaitedF {
				rating += 1
			}
			if maxWaitF*1.2 >= timeWaitedF {
				rating += 1
			}
			if maxWaitF*1.1 >= timeWaitedF {
				rating += 1
			}
			if maxWaitF >= timeWaitedF {
				rating += 1
			}

			summ = summ + rating
			counter += 1
			avg = float64(summ/counter) + 0.5
			fmt.Println("Average ", math.Ceil(avg))
			fmt.Println(math.Ceil(avg))
			fmt.Println(maxWaitF, " ", timeWaitedF)
			fmt.Println("Order", delivery.OrderId, " | Rating: ", rating)

			tableList.mx.Unlock()
		default:
			break
		}
		time.Sleep(time.Second)

	}
}

// trimiterea comenzii
func sendOrder(orderTicket *Order) bool {
	requestBody, marshallErr := json.Marshal(orderTicket)
	if marshallErr != nil {
		log.Fatalln(marshallErr)
	}

	request, newRequestError := http.NewRequest(http.MethodPost, kitchenHost+":8000"+"/order", bytes.NewBuffer(requestBody))
	if newRequestError != nil {
		fmt.Println("Could not create new request. Error:", newRequestError)
		log.Fatal(newRequestError)
	} else {
		response, doError := http.DefaultClient.Do(request)
		fmt.Println("Sending order to kitchen attempt")
		if doError != nil {
			fmt.Println("ERROR Sending request. ERR:", doError)
			log.Fatal(doError)
		}
		var responseBody = make([]byte, response.ContentLength)
		response.Body.Read(responseBody)
		fmt.Println("Response: ", string(responseBody))
		if string(responseBody) != "OK" {
			return false
		}
		return true
	}
	return true
}
