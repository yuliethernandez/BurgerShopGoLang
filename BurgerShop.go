package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Order struct {
	burger []Burger
	drinks []Drink
	sides  []Side
	combos []Combo
}

type Combo struct {
	price  float64 //price: The price of the combo is the price of the three items with a 20% discount
	burger Burger
	side   Side
	drink  Drink
}

type Side struct {
	name  string
	price float64
}

type Drink struct {
	typeDrink string
	price     float64
}

type Burger struct {
	withBun     bool
	ingredients string
	price       float64
}

func main() {

	//Take an order from a user, including a burger, a drink, a side, or a combo.
	fmt.Print("Welcome to Burger Shop")
	take_order_from_user()

}

func exitProgram(option string) bool {
	if option == "exit" || option == "quit" {
		return true
	}
	return false
}

func take_order_from_user() {
	orderCustomer := Order{}
	var opt, option string
	var totalOrder, totalSide, totalCombo, totalDrinks, totalBurgers float64
	for {
		fmt.Println("Choose the order: \n1. Burger \n2. Drink \n3. Side \n4. Combo")
		fmt.Scan(&opt)
		if exitProgram(opt) {
			break
		}
		switch opt {
		case "1":
			{
				var quant string
				fmt.Print("How many burgers do you would like to order?")
				fmt.Scan(&quant)
				quantity, _ := strconv.Atoi(quant)
				for i := 0; quantity > 0; i++ {
					a := user_input_burger()
					orderCustomer.burger = append(orderCustomer.burger, a)
					quantity--
				}
			}
		case "2":
			{
				var quant string
				fmt.Print("How many drinks do you would like to order?")
				fmt.Scan(&quant)
				quantity, _ := strconv.Atoi(quant)
				for i := 0; quantity > 0; i++ {
					a := user_input_drink()
					orderCustomer.drinks = append(orderCustomer.drinks, a)
					quantity--
				}
			}
		case "3":
			{
				var quant string
				fmt.Print("How many sides do you would like to order?")
				fmt.Scan(&quant)
				quantity, _ := strconv.Atoi(quant)
				for i := 0; quantity > 0; i++ {
					a := user_input_side()
					orderCustomer.sides = append(orderCustomer.sides, a)
					quantity--

				}
			}
		case "4":
			{
				var quant string
				fmt.Print("How many combos do you would like to order?")
				fmt.Scan(&quant)
				quantity, _ := strconv.Atoi(quant)
				for i := 0; quantity > 0; i++ {
					a := user_input_combo()
					orderCustomer.combos = append(orderCustomer.combos, a)
					quantity--

				}
			}
		}

		fmt.Println("Do you would like to order something else? Y/N: ")
		fmt.Scan(&option)
		if strings.ToLower(option) == "n" {
			break
		}
	}

	//Display the completed order to the user.
	//orderEmpty := Order{}
	if orderCustomer.burger != nil || orderCustomer.combos != nil || orderCustomer.sides != nil || orderCustomer.drinks != nil {
		sides := orderCustomer.sides
		for _, val := range sides {
			totalSide += val.price
		}
		//var totalCombo float64
		combos := orderCustomer.combos
		for _, val := range combos {
			totalCombo += val.price
		}
		//var totalDrinks float64
		drinks := orderCustomer.drinks
		for _, val := range drinks {
			totalDrinks += val.price
		}
		//var totalBurgers float64
		burgers := orderCustomer.burger
		for _, val := range burgers {
			totalBurgers += val.price
		}
		totalOrder = totalBurgers + totalCombo + totalDrinks + totalSide

		fmt.Println("\n---------------------------------------*** Your order ***")
		if len(orderCustomer.burger) > 0 {
			fmt.Println("Burgers: ", len(orderCustomer.burger), "  Total Price: $", totalBurgers)
		}
		if len(orderCustomer.sides) > 0 {
			fmt.Println("Sides: ", len(orderCustomer.sides), "  Total Price: $", totalSide)
		}
		if len(orderCustomer.combos) > 0 {
			fmt.Println("Combos: ", len(orderCustomer.combos), "  Total Price: $", totalCombo)
		}
		if len(orderCustomer.drinks) > 0 {
			fmt.Println("Drinks: ", len(orderCustomer.drinks), "  Total Price: $", totalDrinks)
		}
		fmt.Println("\n--------------------------------------- \nTotal cost of the Order: $", totalOrder, "\nThank you for your order! \n\n---------------------------------------")

		var confirmation string
		fmt.Print("Do you wanna confirm the order? y/n")
		fmt.Scan(&confirmation)
		if strings.ToLower(confirmation) == "y" {
			fmt.Println("")
			fmt.Println("***Your order will be ready soon*** \n**Thank you for choose BurgerShop!**")
		}
	}
}

/*func (o Order) IsEmpty() bool {
	return reflect.DeepEqual(o, Order{})
}*/

func user_input_burger() (burger Burger) { //Customize selected items, including toppings for the burger.
	burger = Burger{}
	//bun
	var bun, opt string
	fmt.Print("Do you would like your burger with bun or not? Y/N? ")
	fmt.Scan(&bun)
	if exitProgram(opt) {
		return burger
	}
	if strings.ToLower(bun) == "y" {
		burger.withBun = true
	} else {
		burger.withBun = false
	}
	//Price
	if burger.withBun {
		burger.price = 5
	} else {
		burger.price = 4
	}
	//Ingredients
	var ingredients string
	fmt.Println("Choose the ingredients separate by space \n 1. Onion \n 2. Tomatoes \n 3. Lettuce \n 4. Mayonaise \n 5. Ketchup")
	fmt.Scan(&ingredients)
	//var ingredients []string = strings.Split(ing, " ")
	burger.ingredients = ingredients
	return burger
}

func user_input_drink() (drink Drink) {
	drink = Drink{}
	var d string
	fmt.Println("Drinks: \n1. Coke Zero zugar\n2. Pepsi \n3. Coke Cola \n4. Water \nChoose the number of your drink")
	fmt.Scan(&d)
	switch d {
	case "1":
		{
			drink := Drink{"Coke Zero zugar", 2}
			return drink
		}
	case "2":
		{
			drink := Drink{"Pepsi", 2}
			return drink
		}
	case "3":
		{
			drink := Drink{"Coke Cola", 2}
			return drink
		}
	case "4":
		{
			drink := Drink{"Water", 1}
			return drink
		}
	}
	return drink
}

func user_input_side() (side Side) {
	side = Side{}
	var d string
	fmt.Println("Sides: \n1. Frites\n2. Onions Rings\n3. Salad \nChoose the number of your side")
	fmt.Scan(&d)
	switch d {
	case "1":
		{
			side := Side{"Frites", 1}
			return side
		}
	case "2":
		{
			side := Side{"Onions Rings", 2}
			return side
		}
	case "3":
		{
			side := Side{"Salad", 2}
			return side
		}
	}
	return side
}

func user_input_combo() (combo Combo) {
	combo = Combo{}

	fmt.Println("The combo contains: 1 Burger, 1 Drink and 1 Side")

	fmt.Println("Choose your Burger: ")
	burger := user_input_burger()
	fmt.Println("Choose your Drink: ")
	drink := user_input_drink()
	fmt.Println("Choose your Side: ")
	side := user_input_side()

	combo.burger = burger
	combo.side = side
	combo.drink = drink

	//price: The price of the combo is the price of the three items with a 20% discount
	totalPriceCombo := burger.price + drink.price + side.price
	discount := totalPriceCombo * 0.20
	totalPriceCombo -= discount
	combo.price = totalPriceCombo

	return combo
}
