package main

import (
	"fmt"
)

var printOut = fmt.Println
var input = fmt.Scanln

func main() {
	//Number 1
	var milesDriven int
	var gallonsUsed int
	var prompt string
	var result int

	for i := 0; i < 10; i++ {

		printOut("Enter miles driven")
		input(&milesDriven)

		printOut("Enter gallons used")
		input(&gallonsUsed)

		milesPerGallon := milesDriven / gallonsUsed

		array := [10]int{milesPerGallon}

		printOut("Would you like to continue ?")
		input(&prompt)
		if prompt == "yes" {
			index := 1
			for index = 0; index <= 100; index++ {
				result = array[index]
			}
			continue
		} else if prompt == "no" {
			printOut(result)
			break
		} else {
			break
		}

	}

	////Number 2
	//var accountNumber int
	//var beginningBalance int
	//var itemsCharged int
	//var creditForTheMonth int
	//var creditLimit int
	//
	//printOut("Enter account number ")
	//fmt.Scanln(&accountNumber)
	//
	//printOut("Enter balance at the beginning of the month")
	//fmt.Scanln(&beginningBalance)
	//
	//printOut("Enter items charged")
	//fmt.Scanln(&itemsCharged)
	//
	//printOut("Enter credit amount")
	//fmt.Scanln(&creditForTheMonth)
	//
	//printOut("Allowed credit limit: ")
	//fmt.Scanln(&creditLimit)
	//
	//var result = beginningBalance + itemsCharged - creditForTheMonth
	//
	//printOut("Account Number: ", accountNumber, "Balance =", result)
	//
	//if result > creditLimit {
	//	printOut("Credit limit exceeded")
	//}

	////Number 3
	//
	//var itemsSold float64
	//var promptUser string
	//var totalAmount float64
	//var salaryAmount float64 = 200
	//
	//i := 0
	//for i = 1; i < 5; i++ {
	//	printOut("Enter total amount of items sold")
	//	fmt.Scanln(&itemsSold)
	//
	//	printOut("Would you like to continue ?")
	//	fmt.Scanln(&promptUser)
	//
	//	if promptUser == "no" {
	//		totalAmount += itemsSold
	//		percentage := itemsSold * (9.0 / 100)
	//		result := salaryAmount + percentage
	//		printOut("items", i, "Earnings = ", result)
	//		break
	//	} else if promptUser == "yes" {
	//		continue
	//	} else {
	//		break
	//	}
	//
	//}

	////Number 4
	//var tax float64
	//for i := 1; i <= 3; i++ {
	//	printOut("Good morning, welcome to Our Tax Calculator Program")
	//	printOut("Please input your name")
	//	var name string
	//	fmt.Scanln(&name)
	//	printOut("Hello", name)
	//	printOut("Please Input your salary")
	//	var salary float64
	//	fmt.Scanln(&salary)
	//
	//	if salary <= 30000 {
	//		tax = 15.0 / 100 * salary
	//	} else if salary > 30000 {
	//		tax = 20.0 / 100 * salary
	//	}
	//	printOut("Your tax is", tax)
	//
	//	printOut("Would you like to continue ?")
	//	var cont string
	//	fmt.Scanln(&cont)
	//
	//	if cont == "no" {
	//		printOut("Thank you for using our app \n Hope to see you again")
	//		break
	//	} else if cont == "yes" {
	//		continue
	//	} else {
	//		break
	//	}
	//}

	//Number 5
	//var numbers int
	//var largestNumber int
	//
	//for i := 1; i <= 10; i++ {
	//	printOut("Enter number")
	//	fmt.Scanln(&numbers)
	//	var array [11]int
	//	array[i] = numbers
	//
	//	if array[i] > largestNumber {
	//		largestNumber = array[i]
	//	}
	//}
	//print(largestNumber)
}
