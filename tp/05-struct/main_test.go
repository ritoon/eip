package main

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	expected := 15.0
	result := rect.Area()

	if result != expected {
		t.Errorf("Rectangle Area() = %.2f; want %.2f", result, expected)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	expected := 16.0
	result := rect.Perimeter()

	if result != expected {
		t.Errorf("Rectangle Perimeter() = %.2f; want %.2f", result, expected)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 4}
	expected := math.Pi * 4 * 4 // π * r^2
	result := circle.Area()

	if result != expected {
		t.Errorf("Circle Area() = %.2f; want %.2f", result, expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 4}
	expected := 2 * math.Pi * 4 // 2πr
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Circle Perimeter() = %.2f; want %.2f", result, expected)
	}
}

func TestPrintShapeInfo(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	// Cette fonction ne renvoie rien, nous testons simplement qu'elle s'exécute sans erreur
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
}
