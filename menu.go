package main

// stabilirea structurii meniului
type Menu struct {
	id               int
	name             string
	prepTime         int
	complexity       int
	cookingApparatus string
}

// implementarea meniului
var pizza = Menu{1, "Pizza", 20, 2, "oven"}
var salad = Menu{2, "Salad", 10, 1, ""}
var zeama = Menu{3, "Zeama", 7, 1, "stove"}
var sashimi = Menu{4, "Scallop Sashimi with Meyer Lemon Confit", 32, 3, ""}
var duck = Menu{5, "Island Duck with Mulberry Mustard", 35, 3, "oven"}
var waffles = Menu{6, "Waffles", 10, 1, "stove"}
var aubergine = Menu{7, "Aubergine", 20, 2, "oven"}
var lasagna = Menu{8, "Lasagna", 30, 2, "oven"}
var burger = Menu{9, "Burger", 15, 1, "stove"}
var gyros = Menu{10, "Gyros", 15, 1, ""}
var kebab = Menu{11, "Kebab", 15, 1, ""}
var unagi = Menu{12, "Unagi Maki", 20, 2, ""}
var chick = Menu{13, "Tabacco Chicken", 30, 2, "oven"}

// totalizarea meniului
var menu = []Menu{pizza, pizza, salad, zeama, sashimi, duck, waffles, aubergine, lasagna, burger, gyros, kebab, unagi, chick}
