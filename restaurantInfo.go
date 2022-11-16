package main

type RestaurantInfo struct {
	RestaurantId int                `json:"restaurant_id"`
	Name         string             `json:"name"`
	Address      string             `json:"address"`
	MenuItems    int                `json:"menu_items"`
	Menu         []Menu 			`json:"menu"`
	Rating       float32            `json:"rating"`
}