package main

import "fmt"

func FilterEvenNumbers(numbers []int) []int {
	// À implémenter
	return nil
}

func SumSlice(numbers []int) int {
	// À implémenter
	return -1
}

func RemoveDuplicates(numbers []int) []int {
	// À implémenter
	return nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 9, 10}

	evenNumbers := FilterEvenNumbers(numbers)
	sum := SumSlice(numbers)
	uniqueNumbers := RemoveDuplicates(numbers)

	fmt.Println("Nombres pairs:", evenNumbers)
	fmt.Println("Somme des nombres:", sum)
	fmt.Println("Nombres uniques:", uniqueNumbers)
}
