package piscine

type food struct {
	preptime int
	name     string
}

func menu(item *food) {
	if item.name == "burger" {
		item.preptime = 15
	} else if item.name == "chips" {
		item.preptime = 10
	} else if item.name == "nuggets" {
		item.preptime = 12
	} else {
		item.preptime = 404
	}
}

func FoodDeliveryTime(order string) int {
	food := &food{}
	food.name = order
	menu(food)

	return food.preptime
}
