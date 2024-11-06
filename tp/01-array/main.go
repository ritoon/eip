package main

import "fmt"

// 1. Déclarez un array de 10 entiers
var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 2. Fonction qui calcule la somme de tous les éléments d'un array
func SumArray(array [10]int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

// 3. Fonction qui trouve la valeur maximale d'un array
func FindMax(array [10]int) int {
	max := array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}

// 4. Fonction qui inverse les éléments d'un array en place
func ReverseArray(array *[10]int) {
	for i := 0; i < len(array)/2; i++ {
		// Échange des éléments
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
}

// 5. Fonction qui vérifie si un array est trié en ordre croissant
func IsArraySorted(array [10]int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Array initial:", arr)

	// Utilisez SumArray pour calculer la somme des éléments
	sum := SumArray(arr)
	fmt.Println("Somme de l'array:", sum)

	// Utilisez FindMax pour trouver la valeur maximale
	max := FindMax(arr)
	fmt.Println("Valeur maximale de l'array:", max)

	// Utilisez ReverseArray pour inverser les éléments de l'array
	ReverseArray(&arr)
	fmt.Println("Array après inversion:", arr)

	// Vérifiez si l'array est trié
	isSorted := IsArraySorted(arr)
	fmt.Println("L'array est trié:", isSorted)
}
