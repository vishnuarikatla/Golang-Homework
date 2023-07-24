package task1

import (
	"fmt"
)

func Task1() {
	var a, b int
	fmt.Println("Welcome to task 1")
	fmt.Println("Enter first number :")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Printf("You entered: %d\n", a)
	fmt.Println("Enter second number :")
	_, err2 := fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error reading input:", err2)
		return
	}
	fmt.Printf("You entered: %d\n", b)
	add(a, b)

	fmt.Println("Subtraction : ", subtract(a, b))

	c := divide(a, b)
	fmt.Println("Division : ", c)

	d := multiply(a,b)
	fmt.Printf("Multiplication : %d ", d)

	fmt.Println("Remainder", remainder(a, b))

	relational(a, b)
	Logical(a, b)

}

//Arithmetic operators using 4 types of functions

func add(x, y int) {
	fmt.Println("Addition : ", x+y)
}
func subtract(x, y int) int {
	return x - y
}
func divide(x, y int) int {
	if y == 0 {
		fmt.Println("Cannot divide by zero")
		return -1
	}
	return x / y
}
func multiply(x, y int) int {
	return x * y
}
func remainder(x, y int) (r int) {
	r = x % y
	return
}

//Relational Operators using "if"

func relational(x, y int) {
	fmt.Println("\nRelational Operators:")
	if x > y {
		fmt.Println("a is greater than b")
	}

	if x < y {
		fmt.Println("a is less than b")
	}

	if x == y {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is not equal to b")
	}
}

//Logical operators using switch

func Logical(x, y int) {
	switch {
	case x > 0 && y > 0:
		fmt.Println("Both x and y are positive")
	case x > 0 || y > 0:
		fmt.Println("Either x or y is positive")
	case !(x > 0):
		fmt.Println("x is not positive")
	case !(y > 0):
		fmt.Println("y is not positive")
	default:
		fmt.Println("None of the conditions matched")
	}
}
