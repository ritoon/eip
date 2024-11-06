package main

import (
	"testing"
)

// Test pour la conversion Celsius vers Fahrenheit
func TestToFahrenheit(t *testing.T) {
	temp := Temperature(0)
	expected := Temperature(32)
	result := temp.ToFahrenheit()

	if result != expected {
		t.Errorf("ToFahrenheit() = %.2f; want %.2f", result, expected)
	}

	temp = Temperature(100)
	expected = Temperature(212)
	result = temp.ToFahrenheit()

	if result != expected {
		t.Errorf("ToFahrenheit() = %.2f; want %.2f", result, expected)
	}
}

// Test pour la conversion Fahrenheit vers Celsius
func TestToCelsius(t *testing.T) {
	temp := Temperature(32)
	expected := Temperature(0)
	result := temp.ToCelsius()

	if result != expected {
		t.Errorf("ToCelsius() = %.2f; want %.2f", result, expected)
	}

	temp = Temperature(212)
	expected = Temperature(100)
	result = temp.ToCelsius()

	if result != expected {
		t.Errorf("ToCelsius() = %.2f; want %.2f", result, expected)
	}
}

// Test pour la méthode Describe de Temperature
func TestTemperatureDescribe(t *testing.T) {
	temp := Temperature(25)
	expected := "Temperature: 25.00°C (77.00°F)"
	result := temp.Describe()

	if result != expected {
		t.Errorf("Describe() = %s; want %s", result, expected)
	}
}

// Test pour la méthode Describe de Person
func TestPersonDescribe(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	expected := "Person: Alice, Age: 30"
	result := person.Describe()

	if result != expected {
		t.Errorf("Describe() = %s; want %s", result, expected)
	}
}

// Test pour la conversion float64 en int
func TestConvertToInt(t *testing.T) {
	value := 45.67
	expected := 45
	result := ConvertToInt(value)

	if result != expected {
		t.Errorf("ConvertToInt() = %d; want %d", result, expected)
	}

	value = -12.34
	expected = -12
	result = ConvertToInt(value)

	if result != expected {
		t.Errorf("ConvertToInt() = %d; want %d", result, expected)
	}
}
