package main

import (
	"math/rand"
	"testing"
	"time"
)

// Génère une slice d'entiers aléatoires
func generateRandomIntSlice(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(max-min+1) + min
	}
	return slice
}

// Test aléatoire pour Find
func TestFindWithRandomData(t *testing.T) {
	slice := generateRandomIntSlice(20, 0, 10)
	value := rand.Intn(10)

	// Vérifie que Find retourne l'index correct ou -1
	index := Find(slice, value)
	found := false
	for i, v := range slice {
		if v == value {
			if index != i {
				t.Errorf("Find() = %d; want %d", index, i)
			}
			found = true
			break
		}
	}
	if !found && index != -1 {
		t.Errorf("Find() = %d; want -1", index)
	}
}

// Test aléatoire pour Sum
func TestSumWithRandomData(t *testing.T) {
	slice := generateRandomIntSlice(10, 1, 100)
	expected := 0
	for _, v := range slice {
		expected += v
	}

	result := Sum(slice)
	if result != expected {
		t.Errorf("Sum() = %d; want %d", result, expected)
	}
}

// Test aléatoire pour Pair.Swap
func TestPairSwapWithRandomData(t *testing.T) {
	first := rand.Intn(100)
	second := rand.Intn(100)
	pair := Pair[int]{First: first, Second: second}

	// Test Swap
	pair.Swap()

	if pair.First != second || pair.Second != first {
		t.Errorf("Pair.Swap() = {%d, %d}; want {%d, %d}", pair.First, pair.Second, second, first)
	}
}
