package main

import (

	"homework/task1"
	"homework/task2"
	"fmt"
	"homework/roulette"
)

func main() {
	fmt.Println("Some cganges for practice")
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Task 1")
		fmt.Println("2. Task 2")
		fmt.Println("3. Roulette")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter task num: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			task1.Task1()
		case 2:
			task2.Task2()
		case 3:
			roulette.Roulette()
		case 4:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
