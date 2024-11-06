package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 1. Fonction générique Find qui retourne l'index d'une valeur dans une slice
func Find[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// 2. Fonction générique Sum qui calcule la somme des éléments d'une slice numérique
func Sum[T constraints.Ordered](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 3. Struct générique Pair avec deux valeurs de type T
type Pair[T any] struct {
	First, Second T
}

// 4. Méthode Swap pour Pair qui échange les valeurs First et Second
func (p *Pair[T]) Swap() {
	p.First, p.Second = p.Second, p.First
}

func main() {
	// Exemple d'utilisation de Find avec une slice de chaînes
	strings := []string{"apple", "banana", "cherry"}
	fmt.Println("Index de 'banana':", Find(strings, "banana")) // Devrait afficher 1
	fmt.Println("Index de 'orange':", Find(strings, "orange")) // Devrait afficher -1

	// Exemple d'utilisation de Sum avec une slice d'entiers
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("Somme des entiers:", Sum(ints)) // Devrait afficher 15

	// Exemple d'utilisation de Sum avec une slice de float64
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println("Somme des floats:", Sum(floats)) // Devrait afficher 6.6

	// Exemple d'utilisation de Pair avec des valeurs de type string
	p := Pair[string]{First: "Hello", Second: "World"}
	fmt.Println("Avant Swap:", p)
	p.Swap()
	fmt.Println("Après Swap:", p)

	// Exemple d'utilisation de Pair avec des valeurs de type int
	p2 := Pair[int]{First: 1, Second: 2}
	fmt.Println("Avant Swap:", p2)
	p2.Swap()
	fmt.Println("Après Swap:", p2)
}
