package task2

import "fmt"

//Defining a astructure to perform arithmetic operations

var a, b int 

type Operation struct{
	x, y int
}

// Methods

func (op Operation) Add() int{
	return op.x + op.y
}
func (op Operation) Subtract() int{
	return op.x - op.y
}
func (op Operation) Multiply() int{
	return op.x * op.y
}
func (op Operation) Divide() float64{
	if op.y != 0{
		return float64(op.x) / float64(op.y)
} 
    return 0

}

type Relational_Operation interface{

	LessThan() bool
	EqualTo() bool
	GreaterThan() bool


}

//Methods to perform Relational operation

func (op Operation) LessThan() bool{
	return op.x < op.y
}

func (op Operation) EqualTo() bool{
	return op.x == op.y
}
	
func (op Operation) GreaterThan() bool{
	return op.x > op.y
}


func Task2(){
	fmt.Println("Welcome to task 2")
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
	Structures()
	Arrays()
	Slices()
	Pointers()

}

func Structures(){
	op := Operation{a, b}

	fmt.Printf("Addition : %d\n", op.Add())
	fmt.Printf("Subtraction : %d\n", op.Subtract())
	fmt.Printf("Multiplication : %d\n", op.Multiply())
	fmt.Printf("Division : %.2f\n", op.Divide())
    
	// INTERFACE
	fmt.Printf("Less Than: %t\n", op.LessThan())
	fmt.Printf("Equal To: %t\n", op.EqualTo())
	fmt.Printf("Greater Than: %t\n", op.GreaterThan())

}

func Arrays(){
	numArray := [2]int{a, b}
	fmt.Println("Array Example:")
	fmt.Printf("Addition: %d\n", numArray[0]+numArray[1])
	fmt.Printf("Subtraction: %d\n", numArray[0]-numArray[1])
	fmt.Printf("Multiplication: %d\n", numArray[0]*numArray[1])
	fmt.Printf("Division: %.2f\n", float64(numArray[0])/float64(numArray[1]))
}

func Slices(){
	numSlice := []int{a, b}
	fmt.Println("\nSlice Example:")
	fmt.Printf("Addition: %d\n", numSlice[0]+numSlice[1])
	fmt.Printf("Subtraction: %d\n", numSlice[0]-numSlice[1])
	fmt.Printf("Multiplication: %d\n", numSlice[0]*numSlice[1])
	fmt.Printf("Division: %.2f\n", float64(numSlice[0])/float64(numSlice[1]))

}

func Pointers(){
	
	ptrA, ptrB := &a, &b
	fmt.Println("\nPointer Example:")
	fmt.Printf("Addition: %d\n", *ptrA+*ptrB)
	fmt.Printf("Subtraction: %d\n", *ptrA-*ptrB)
	fmt.Printf("Multiplication: %d\n", *ptrA**ptrB)
	fmt.Printf("Division: %.2f\n", float64(*ptrA)/float64(*ptrB))
}