package main

import (
	"fmt"
)

func main() {

	/* This program demonstrates Modular Arithmatic for an value "Base" input, and two values added together */

	fmt.Print("\nModular Arithmatic\n")
	fmt.Println("This will demonstrate modular arithmatic with the base of your choice:")
	run := "no" //The variable "run" acts as the the While loop for the program
	for run == "no" || run == "n" || run == "N" || run == "NO" {

		var x, y, b, numBase, remain int // all of our math and inputs will be integer based.
		fmt.Print("What is the base value for calculation (example, 12 for a clock)?: ")
		fmt.Scanf("%d", &b) // user input for Base
		fmt.Print("Enter the first value of the calculation:")
		fmt.Scanf("%d", &x) // user input for first value to add.
		fmt.Print("Enter the second value of the calculation:")
		fmt.Scanf("%d", &y) // user input for second value to add.

		fmt.Printf("Our two numbers to add %d, %d where base is %d, \n", x, y, b) // display the Base and two values to add.

		numBase = ((x + y) / b) / 2 // calculate the number times through the Base.
		remain = (x + y) % b        // the remainder in the modulus arithmatic.

		fmt.Printf("\n%d is the Number of Times through the Base value, the value of the remainder is %d.\n\n", numBase, remain) // The output of the calculation.
		fmt.Print("Exit?")
		fmt.Scanf("%s", &run) // Prompt for input to do another calculation
		fmt.Println()
	}

} //End of program
