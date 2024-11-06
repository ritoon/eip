package main

import (
	"fmt"
	"math"
)

// 1. Struct Rectangle avec largeur et hauteur
type Rectangle struct {
	Width, Height float64
}

// 2. Méthode pour calculer l'aire du rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3. Méthode pour calculer le périmètre du rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 4. Struct Circle avec rayon
type Circle struct {
	Radius float64
}

// 5. Méthode pour calculer l'aire du cercle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Méthode pour calculer le périmètre du cercle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 6. Interface Shape avec méthodes Area et Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 7. Fonction pour afficher les informations d'un Shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Aire: %.2f\n", s.Area())
	fmt.Printf("Périmètre: %.2f\n", s.Perimeter())
}

func main() {
	// Exemple d'utilisation avec Rectangle
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Rectangle :")
	PrintShapeInfo(rect)

	// Exemple d'utilisation avec Circle
	circle := Circle{Radius: 4}
	fmt.Println("\nCercle :")
	PrintShapeInfo(circle)
}
