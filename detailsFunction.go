package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to my site")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)
	return num1, num2
}

func addTwoNumbers(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayResult(sum int, name string) {
	fmt.Println("Hello", name)
	fmt.Println("The sum of the two numbers is:", sum)
}

func goodbyeMessage() {
	fmt.Println("Thank you for visiting my site")
	fmt.Println("Goodbye!")
}

func runDetailsFunction() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := addTwoNumbers(num1, num2)
	displayResult(sum, name)
	goodbyeMessage()
}

func main() {
	runDetailsFunction()
}
