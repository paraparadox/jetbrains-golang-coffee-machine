package main

import "fmt"

func main() {
	determineNecessaryAmountOfIngredients()
}

func makingCoffeeProcess() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")
}

func determineNecessaryAmountOfIngredients() {
	neededCupsCount := askNeededCupsCount()
	fmt.Printf("For %d cups of coffee you will need:\n", neededCupsCount)
	fmt.Printf("%d ml of water\n", calculateRequiredAmountOfWater(neededCupsCount))
	fmt.Printf("%d ml of milk\n", calculateRequiredAmountOfMilk(neededCupsCount))
	fmt.Printf("%d g of coffee beans\n", calculateRequiredAmountOfCoffeeBeans(neededCupsCount))
}

func askNeededCupsCount() int {
	var cupsCount int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Print("> ")
	fmt.Scan(&cupsCount)
	return cupsCount
}

func calculateRequiredAmountOfWater(cupsCount int) int {
	const waterMlPerCup = 200
	return cupsCount * waterMlPerCup
}

func calculateRequiredAmountOfMilk(cupsCount int) int {
	const milkMlPerCup = 50
	return cupsCount * milkMlPerCup
}

func calculateRequiredAmountOfCoffeeBeans(cupsCount int) int {
	const coffeeGrPerCup = 15
	return cupsCount * coffeeGrPerCup
}