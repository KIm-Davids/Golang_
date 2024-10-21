package main

import (
	"fmt"
	"regexp"
)

func printValuesFromOddIndex(array []int) {
	var newArray [10]int
	var counter int
	var result int

	for counter = 1; counter < len(array); counter += 2 {
		newArray[counter] = array[counter]
	}

	fmt.Print(newArray)

	for counter = 1; counter < len(array); counter += 2 {
		result += array[counter]
	}

	fmt.Println("\n", result)
}

func printValuesFromEvenIndex(array []int) {
	var newArray [10]int
	var counter int
	var result int

	for counter = 0; counter < len(array); counter += 2 {
		newArray[counter] = array[counter]
	}
	fmt.Print(newArray)

	for counter = 0; counter < len(array); counter += 2 {
		result += array[counter]
	}

	fmt.Println("\n", result)
}

func printValuesFromArrayHorizontally() {
	numbers, err := InputValuesIntoArray()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numbers)
}

func printValuesFromArrayVertically() {
	var counter int
	numbers, err := InputValuesIntoArray()

	if err != nil {
		fmt.Println(err)
	}

	for counter = 0; counter < len(numbers); counter++ {
		fmt.Println(numbers[counter])

	}
}

func InputValuesIntoArray() ([10]int, error) {
	var numbers int

	var counter int
	var array [10]int

	for counter = 0; counter < 10; counter++ {
		fmt.Println("Enter 10 numbers")
		fmt.Scan(&numbers)
		array[counter] = numbers
	}
	return array, nil

}

func printName() {
	var name interface{}
	fmt.Print("Enter a name")
	//name := fmt.Scanner(&name)
	fmt.Print("Hello", name)
}

func containsAllAlphabet(alphabet string) bool {
	regexCheck := `^[a-zA-Z]`
	check := regexp.MustCompile(regexCheck)

	if check.MatchString(alphabet) {
		return true
	}
	return false
}

func checkLetters(word string) bool {
	var array []string
	newArray := append(array, word)
	for index, values := range newArray {
		if values[index] == values[index+1] {
			return true
		}
	}
	return false
}

func main() {

	//array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//printValuesFromArrayVertically()
	//printValuesFromArrayHorizontally()
	//printValuesFromEvenIndex(array)
	//printValuesFromOddIndex(array)
	//fmt.Print(containsAllAlphabet("d"))
	fmt.Print(checkLetters("hheelloo" +
		""))

}
