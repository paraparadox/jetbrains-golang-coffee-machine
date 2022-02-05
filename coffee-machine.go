package main

import "fmt"

func main() {
	money := 550
	water := 400
	milk := 540
	coffeeBeans := 120
	disposableCups := 9

	displaySupplies(water, milk, coffeeBeans, disposableCups, money)
	fmt.Println()
	action := askForAction()
	processTheAction(action, &water, &milk, &coffeeBeans, &disposableCups, &money)
	fmt.Println()
	displaySupplies(water, milk, coffeeBeans, disposableCups, money)
}

func displaySupplies(water, milk, coffeeBeans, disposableCups, money int) {
	fmt.Println("The coffee machine has:")
	fmt.Println(water, "of water")
	fmt.Println(milk, "of milk")
	fmt.Println(coffeeBeans, "of coffee beans")
	fmt.Println(disposableCups, "of disposable cups")
	fmt.Println(money, "of money")
}

func askForAction() string {
	var action string
	fmt.Print("Write action (buy, fill, take):\n> ")
	fmt.Scan(&action)
	return action
}

func processTheAction(action string, water, milk, coffeeBeans, disposableCups, money *int) {
	switch action {
	case "buy":
		processTheBuyAction(water, milk, coffeeBeans, disposableCups, money)
	case "fill":
		processTheFillAction(water, milk, coffeeBeans, disposableCups, money)
	case "take":
		processTheTakeAction(water, milk, coffeeBeans, disposableCups, money)
	}
}

func processTheBuyAction(water, milk, coffeeBeans, disposableCups, money *int) {
	var coffeeType int
	fmt.Print("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:\n> ")
	fmt.Scan(&coffeeType)
	switch coffeeType {
	case 1:
		sellEspresso(water, milk, coffeeBeans, money)
	case 2:
		sellLatte(water, milk, coffeeBeans, money)
	case 3:
		sellCappuccino(water, milk, coffeeBeans, money)
	}
	*disposableCups -= 1
}

func sellEspresso(water, milk, coffeeBeans, money *int) {
	waterPerCup := 250
	milkPerCup := 0
	coffeePerCup := 16
	pricePerCup := 4

	*water -= waterPerCup
	*milk -= milkPerCup
	*coffeeBeans -= coffeePerCup
	*money += pricePerCup
}

func sellLatte(water, milk, coffeeBeans, money *int) {
	waterPerCup := 350
	milkPerCup := 75
	coffeePerCup := 20
	pricePerCup := 7

	*water -= waterPerCup
	*milk -= milkPerCup
	*coffeeBeans -= coffeePerCup
	*money += pricePerCup
}

func sellCappuccino(water, milk, coffeeBeans, money *int) {
	waterPerCup := 200
	milkPerCup := 100
	coffeePerCup := 12
	pricePerCup := 6

	*water -= waterPerCup
	*milk -= milkPerCup
	*coffeeBeans -= coffeePerCup
	*money += pricePerCup
}

func processTheFillAction(water, milk, coffeeBeans, disposableCups, money *int) {

}

func processTheTakeAction(water, milk, coffeeBeans, disposableCups, money *int) {

}
