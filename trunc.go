/*Write a program which prompts the user to enter a floating point number and prints the integer which
is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.*/

package main

import "fmt"

func main() {

	fmt.Println("Enter a floating1 point number: ")
	
	var floating1 float32
	fmt.Scanln(&floating1)
	var x float32 = floating1
	var y int = int(x)
	fmt.Println(y)
	
}
