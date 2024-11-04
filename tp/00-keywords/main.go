// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hallo Welt!")
// }

package main

import "fmt"

// Déclaration de constantes et de variables globales
const Pi = 3.14

var globalVar = "Hello"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return Pi * c.radius * c.radius
}

func main() {
	// Déclaration de variables locales
	var a int = 10
	var b int = 20

	// Boucle for avec if, else et continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
			continue
		} else {
			fmt.Println("Odd:", i)
		}
	}

	// Utilisation de switch, case et fallthrough
	switch a {
	case 10:
		fmt.Println("a is 10")
		fallthrough
	case 20:
		fmt.Println("a might also be 20")
	default:
		fmt.Println("a is something else")
	}

	// Utilisation de defer et func
	defer fmt.Println("This is deferred until the end of main")

	// Déclaration et utilisation d'une map et d'une struct
	myMap := make(map[string]int)
	myMap["key1"] = 100
	myMap["key2"] = 200

	circle := Circle{radius: 5}
	fmt.Println("Area of circle:", circle.Area())

	// Utilisation de chan, go, et select
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received from channel:", msg)
	default:
		fmt.Println("No message received")
	}

	// Utilisation de const, goto et label
	fmt.Println("Pi constant:", Pi)

	for i := 0; i < 3; i++ {
		if i == 1 {
			goto MyLabel
		}
		fmt.Println("Looping with goto:", i)
	}
MyLabel:

	// Utilisation de range
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Utilisation de type et interface avec struct et method
	var s Shape = circle
	fmt.Println("Area through interface:", s.Area())

	// Utilisation de return
	result := add(a, b)
	fmt.Println("Result of add function:", result)
}

// Fonction add pour démontrer func et return
func add(x, y int) int {
	return x + y
}
