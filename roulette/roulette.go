package roulette

import (
	"fmt"
	"math/rand"
	"time"
)

// Structure to represent the user's bankroll
type Bankroll struct {
	amount float64
}

// Enum to represent the bet type
type BetType int

const (
	RedOrBlack BetType = iota
	HighOrLow
	SpecificNumber
)

// Function to initialize the bankroll
func initBankroll() Bankroll {
	var amt float64
	fmt.Print("Enter your initial bankroll: $")
	fmt.Scan(&amt)
	return Bankroll{amount: amt}
}

// Function to prompt the user for the bet type
func promptBetType() BetType {
	fmt.Println("\nChoose a bet type:")
	fmt.Println("1. Bet on Red or Black")
	fmt.Println("2. Bet on High or Low")
	fmt.Println("3. Bet on a Specific Number")
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return BetType(choice)
}

// Function to get the user's bet amount
func getBetAmount(bankroll Bankroll) float64 {
	var amount float64
	for {
		fmt.Printf("Enter your bet amount (up to $%.2f): $", bankroll.amount)
		fmt.Scan(&amount)
		if amount <= bankroll.amount && amount > 0 {
			break
		}
		fmt.Println("Invalid bet amount. Please try again.")
	}
	return amount
}

// Function to generate a random roulette number
func spinRoulette() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(38)
}

// Function to determine the color of the roulette number
func getColor(number int) string {
	if number == 0 || number == 37 {
		return "Green"
	} else if number%2 == 0 {
		return "Red"
	} else {
		return "Black"
	}
}

// Function to simulate the roulette game
func playRoulette(bankroll Bankroll) {
	for bankroll.amount > 0 {
		fmt.Printf("\nYour current bankroll: $%.2f\n", bankroll.amount)
		betType := promptBetType()

		switch betType {
		case RedOrBlack +1:
			colorBet := getColorBet()
			amt := getBetAmount(bankroll)
			result := spinRoulette()
			color := getColor(result)

			fmt.Printf("The roulette landed on: %d (%s)\n", result, color)
			if color == colorBet {
				bankroll.amount += amt
				fmt.Printf("Congratulations! You won $%.2f\n", amt)
			} else {
				bankroll.amount -= amt
				fmt.Printf("Sorry, you lost $%.2f\n", amt)
			}

		case HighOrLow +1:
			highLowBet := getHighLowBet()
			amount := getBetAmount(bankroll)
			result := spinRoulette()

			fmt.Printf("The roulette landed on: %d\n", result)
			if (result >= 1 && result <= 18 && highLowBet == "low") || (result >= 19 && result <= 36 && highLowBet == "high") {
				bankroll.amount += amount
				fmt.Printf("Congratulations! You won $%.2f\n", amount)
			} else {
				bankroll.amount -= amount
				fmt.Printf("Sorry, you lost $%.2f\n", amount)
			}

		case SpecificNumber +1:
			numberBet := getNumberBet()
			amount := getBetAmount(bankroll)
			result := spinRoulette()

			fmt.Printf("The roulette landed on: %d\n", result)
			if result == numberBet {
				bankroll.amount += amount * 35
				fmt.Printf("Congratulations! You won $%.2f\n", amount*35)
			} else {
				bankroll.amount -= amount
				fmt.Printf("Sorry, you lost $%.2f\n", amount)
			}

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

	fmt.Println("\nGame over! You have run out of money.")
}

// Function to get the color bet (red or black)
func getColorBet() string {
	var colorBet string
	fmt.Print("Choose color (red or black): ")
	fmt.Scan(&colorBet)
	return colorBet
}

// Function to get the high/low bet
func getHighLowBet() string {
	var highLowBet string
	fmt.Print("Choose high or low: ")
	fmt.Scan(&highLowBet)

	if highLowBet == "high" || highLowBet == "low" {
		return highLowBet
	} else {
		for highLowBet != "high" && highLowBet != "low" {
			fmt.Println("Enter a valid choice")
			fmt.Scan(&highLowBet)
			if highLowBet == "high" || highLowBet == "low" {
				break
			}
		}
	}
	return highLowBet
}

// Function to get the specific number bet
func getNumberBet() int {
	var numberBet int
	fmt.Print("Choose a number (0-37): ")
	fmt.Scan(&numberBet)
	return numberBet
}

func Roulette() {
	bankroll := initBankroll()
	playRoulette(bankroll)
}
