package main

import (
	"math/rand"
	"testing"
	"time"
)

// Génère un slice d'entiers aléatoires de longueur spécifiée
func generateRandomSlice(length, min, max int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, length)
	for i := range slice {
		slice[i] = rand.Intn(max-min+1) + min
	}
	return slice
}

func TestFilterEvenNumbersWithRandomData(t *testing.T) {
	// Génère un slice de nombres aléatoires
	randomNumbers := generateRandomSlice(20, 1, 100)
	result := FilterEvenNumbers(randomNumbers)

	// Vérifie que tous les nombres dans le résultat sont pairs
	for _, num := range result {
		if num%2 != 0 {
			t.Errorf("FilterEvenNumbers() returned an odd number: %d", num)
		}
	}
}

func TestSumSliceWithRandomData(t *testing.T) {
	randomNumbers := generateRandomSlice(10, 1, 10)
	expectedSum := 0
	for _, num := range randomNumbers {
		expectedSum += num
	}
	result := SumSlice(randomNumbers)

	if result != expectedSum {
		t.Errorf("SumSlice(%v) = %d; want %d", randomNumbers, result, expectedSum)
	}
}

func TestRemoveDuplicatesWithRandomData(t *testing.T) {
	randomNumbers := generateRandomSlice(15, 1, 10) // valeurs entre 1 et 10 pour augmenter les doublons
	result := RemoveDuplicates(randomNumbers)

	// Vérifie qu'il n'y a pas de doublons dans le résultat
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			t.Errorf("RemoveDuplicates() returned duplicate value: %d", num)
		}
		seen[num] = true
	}
}
