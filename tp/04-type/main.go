package main

import (
	"fmt"
)

// 1. Créer un type personnalisé Temperature basé sur float64
type Temperature float64

// 2. Méthode pour convertir Celsius en Fahrenheit
func (t Temperature) ToFahrenheit() Temperature {
	return t*9/5 + 32
}

// Méthode pour convertir Fahrenheit en Celsius
func (t Temperature) ToCelsius() Temperature {
	return (t - 32) * 5 / 9
}

// 3. Créer une struct Person avec Name et Age
type Person struct {
	Name string
	Age  int
}

// 4. Définir une interface Describer
type Describer interface {
	Describe() string
}

// Implémentation de Describe pour Person
func (p Person) Describe() string {
	return fmt.Sprintf("Person: %s, Age: %d", p.Name, p.Age)
}

// Implémentation de Describe pour Temperature
func (t Temperature) Describe() string {
	return fmt.Sprintf("Temperature: %.2f°C (%.2f°F)", t, t.ToFahrenheit())
}

// 5. Fonction pour afficher la description d'un Describer
func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}

// 6. Fonction pour convertir un float64 en int
func ConvertToInt(value float64) int {
	return int(value)
}

func main() {
	// Exemple d'utilisation de Temperature
	temp := Temperature(25.0)
	fmt.Println("En Fahrenheit:", temp.ToFahrenheit())
	fmt.Println("En Celsius:", temp.ToCelsius())
	PrintDescription(temp)

	// Exemple d'utilisation de Person
	person := Person{Name: "Alice", Age: 30}
	PrintDescription(person)

	// Exemple de conversion float64 en int
	value := 45.67
	fmt.Printf("Valeur convertie: %d\n", ConvertToInt(value))
}
