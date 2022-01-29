package main

import "fmt"

func main() {
	waterMl, milkMl, coffeeBeansGr := askForAvailableAmountsOfIngredients()
	requestedCupsCount := askNeededCupsCount()
	availableCupsCount := calculateAvailableCupsCount(waterMl, milkMl, coffeeBeansGr)
	respondToUser(requestedCupsCount, availableCupsCount)
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
	fmt.Scan(&cupsCount)
	return cupsCount
}

func askForAvailableAmountsOfIngredients() (int, int, int) {
	var water, milk, coffeeBeans int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&coffeeBeans)
	return water, milk, coffeeBeans
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

func calculateAvailableCupsCount(waterMl, milkMl, coffeeBeansGr int) int {
	const waterMlPerCup = 200
	const milkMlPerCup = 50
	const coffeeGrPerCup = 15
	availableWaterCups := waterMl / waterMlPerCup
	availableMilkCups := milkMl / milkMlPerCup
	availableCoffeCups := coffeeBeansGr / coffeeGrPerCup
	return minOfThree(availableWaterCups, availableMilkCups, availableCoffeCups)
}

func minOfThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func respondToUser(requestedCupsCount, availableCupsCount int) {
	switch {
	case requestedCupsCount == availableCupsCount:
		fmt.Println("Yes, I can make that amount of coffee")
	case requestedCupsCount > availableCupsCount:
		fmt.Printf("No, I can make only %d cups of coffee\n", availableCupsCount)
	case requestedCupsCount < availableCupsCount:
		fmt.Print("Yes, I can make that amount of coffee ")
		fmt.Printf("(and even %d more than that)\n", availableCupsCount - requestedCupsCount)
	}
}