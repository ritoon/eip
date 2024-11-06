package main

import (
	"fmt"
	"math"
	"testing"
)

// Test pour Add
func TestAdd(t *testing.T) {
	var i MyInt = 3
	result := i.Add(5)
	expected := 8
	if result != expected {
		t.Errorf("Add(5) = %d; want %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	var i MyInt = 3
	result := i.Sub(3)
	expected := 0
	if result != expected {
		t.Errorf("Add(5) = %d; want %d", result, expected)
	}
}

// Test pour Multiply
func FuzzMultiply(f *testing.F) {
	testparam := []int{1, math.MaxInt64, math.MinInt64}
	for _, tc := range testparam {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	var i MyInt = 0
	f.Fuzz(func(t *testing.T, orig int) {
		fmt.Println("orig", orig)
		res := i.Add(orig)
		res = i.Sub(res)

		if res != 0 {
			t.Errorf("Before: %v, after: %v with orig %v", 0, i, orig)
		}
	})
}

// Test pour Factorial
// func TestFactorial(t *testing.T) {
// 	result := Factorial(5)
// 	expected := 120
// 	if result != expected {
// 		t.Errorf("Factorial(5) = %d; want %d", result, expected)
// 	}
// }

// // Benchmark pour Add
// func BenchmarkAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(3, 5)
// 	}
// }

// // Benchmark pour Multiply
// func BenchmarkMultiply(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Multiply(3, 5)
// 	}
// }

// // Benchmark pour Factorial
// func BenchmarkFactorial(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Factorial(10)
// 	}
// }

// Fonction simple qui additionne deux entiers
func Add(a, b int) int {
	return a + b
}

func FuzzAdd(f *testing.F) {
	// Ajout de valeurs de départ pour le fuzzing
	f.Add(1, 2)
	f.Add(0, 0)
	f.Add(-1, 1)
	f.Add(100, -50)

	f.Fuzz(func(t *testing.T, a int, b int) {
		// On appelle la fonction Add
		result := Add(a, b)

		// Vérification que la fonction est cohérente avec certaines propriétés
		// Par exemple, le résultat devrait être symétrique : Add(a, b) == Add(b, a)
		if result != Add(b, a) {
			t.Errorf("Add(%d, %d) != Add(%d, %d)", a, b, b, a)
		}

		// Test d'une autre propriété : l'addition avec zéro ne change pas la valeur
		if a != 0 && result == a {
			t.Errorf("Add(%d, %d) devrait être différent de %d", a, b, a)
		}
	})
}
