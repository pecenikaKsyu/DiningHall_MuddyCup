package main

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
