package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

// Génère un array de n entiers aléatoires entre min et max
func generateRandomArray(n, min, max int) [10]int {
	var array [10]int
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(max-min+1) + min
	}
	return array
}

func TestSumArrayWithRandomData(t *testing.T) {
	array := generateRandomArray(10, 0, 100)
	expected := 0
	for _, v := range array {
		expected += v
	}

	result := SumArray(array)

	if result != expected {
		t.Errorf("SumArray() = %d; want %d", result, expected)
	}
}

func TestFindMaxWithRandomData(t *testing.T) {
	array := generateRandomArray(10, 0, 100)
	expected := array[0]
	for _, v := range array {
		if v > expected {
			expected = v
		}
	}

	result := FindMax(array)

	if result != expected {
		t.Errorf("FindMax() = %d; want %d", result, expected)
	}
}

func TestReverseArrayWithRandomData(t *testing.T) {
	array := generateRandomArray(10, 0, 100)
	originalArray := array
	expected := array

	// Crée l'inversion attendue
	for i := 0; i < len(expected)/2; i++ {
		expected[i], expected[len(expected)-1-i] = expected[len(expected)-1-i], expected[i]
	}

	ReverseArray(&array)

	if !reflect.DeepEqual(array, expected) {
		t.Errorf("ReverseArray() = %v; want %v", array, expected)
	}

	// Vérifie qu'un second appel renvoie l'array d'origine
	ReverseArray(&array)
	if !reflect.DeepEqual(array, originalArray) {
		t.Errorf("ReverseArray() did not restore original array, got %v, want %v", array, originalArray)
	}
}

func TestIsArraySortedWithRandomData(t *testing.T) {
	// Test avec un array trié
	array := generateRandomArray(10, 0, 100)
	sort.Ints(array[:]) // Trie l'array

	if !IsArraySorted(array) {
		t.Errorf("IsArraySorted() = false; want true for sorted array")
	}

	// Test avec un array aléatoire non trié
	unsortedArray := generateRandomArray(10, 0, 100)
	if IsArraySorted(unsortedArray) {
		t.Errorf("IsArraySorted() = true; want false for unsorted array")
	}
}
