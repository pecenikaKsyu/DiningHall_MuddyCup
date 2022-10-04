package main

import (
	"math/rand"
	"sync"
	"time"
)

var tableId = 0

// definim structura unui obiect Masa
type Table struct {
	ID           int  // indicele
	readyToOrder bool // nivelul de pregatire
}

// functia initializarii mesei
func newTable() Table {
	var ret Table
	ret.ID = tableId
	tableId += 1
	ret.readyToOrder = true
	return ret
}

// structura listei de mese
type TableList struct {
	tableArr []Table
	mx       sync.Mutex // pentru a putea bloca accesul la thread pana se elibereaza o masa
}

// cream o lista de mese
func newTableList() TableList {
	var tableArr []Table
	for i := 0; i < TableNum; i++ {
		tableArr = append(tableArr, newTable())
	}
	var ret TableList
	ret.tableArr = tableArr
	return ret
}

var orderId = 0

// generarea comenzilor
func (t Table) generateOrder(waiter Waiter) *Order {
	ret := new(Order)
	ret.OrderId = orderId
	orderId += 1
	ret.TableId = t.ID
	ret.Items = getItems()
	ret.Priority = rand.Intn(5)

	//alegem timpul maxim de preparare in comanda
	itemMaxPrep := 0
	for _, item := range ret.Items {
		if menu[item].prepTime > itemMaxPrep {
			itemMaxPrep = menu[item].prepTime
		}
	}
	// setam timpul maxim de asteptare
	ret.MaxWait = int(float32(itemMaxPrep) * 1.4)
	// inregistram momentul primirii comenzii
	ret.PickUpTime = time.Now().Unix()
	return ret
}

// functia de primire a listei de mancaruri
func getItems() []int {
	var ret []int

	var itemNr = rand.Intn(10) + 1
	for i := 0; i < itemNr; i++ {
		ret = append(ret, rand.Intn(13)+1)
	}
	return ret
}
