package main

import "math/rand"

type Order struct {
	OrderId    int   `json:"order_id"`
	TableId    int   `json:"table_id"`
	WaiterId   int   `json:"waiter_id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max_wait"`
	PickUpTime int64 `json:"pick_up_time"`
}

func getItems() []int {
	var ret []int

	var itemNr = rand.Intn(10) + 1
	for i := 0; i < itemNr; i++ {
		ret = append(ret, rand.Intn(10))
	}
	return ret
}
