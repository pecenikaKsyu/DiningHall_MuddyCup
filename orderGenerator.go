package main

import "time"

// Definim structura comenzii
type Order struct {
	OrderId    int   `json:"order_id"`     // codul identificator
	TableId    int   `json:"table_id"`     // masa de la care a fost preluata comanda
	WaiterId   int   `json:"waiter_id"`    // indexul chelnerului
	Items      []int `json:"items"`        // continutul comenzii
	Priority   int   `json:"priority"`     // nivelul d prioritate
	MaxWait    int   `json:"max_wait"`     // timpul maxim de asteptare a comenzii
	PickUpTime int64 `json:"pick_up_time"` // momentul primirii comenzii
}

type OnlineReceivedOrder struct {
	Id          int       `json:"id,omitempty"`
	Items       []int     `json:"items"`
	Priority    int       `json:"priority"`
	MaxWait     float32   `json:"max_wait"`
	CreatedTime time.Time `json:"created_time"`
}

type OnlineResponseOrder struct {
	RestaurantId         int       `json:"restaurant_id"`
	OrderId              int       `json:"order_id"`
	EstimatedWaitingTime int       `json:"estimated_waiting_time"`
	CreatedTime          time.Time `json:"created_time"`
	RegisteredTime       time.Time `json:"registered_time"`
}