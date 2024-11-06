package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExampleJSONSerialization(t *testing.T) {
	phone := "123-456-7890"
	example := User{
		Name:         "John Doe",
		hiddenField:  "This is hidden",
		Age:          0, // Test `omitempty` pour un champ entier avec valeur zéro
		Address:      "123 Main St",
		IgnoredField: "This will not appear",
		Email:        "", // Test `omitempty` pour un champ chaîne vide
		IsActive:     true,
		PhoneNumber:  &phone, // Test d'un champ pointeur non nul
	}

	data, err := json.Marshal(example)
	if err != nil {
		t.Fatalf("Erreur lors de la sérialisation JSON : %v", err)
	}

	jsonString := string(data)

	// Vérifie que le champ `name` est présent
	if !strings.Contains(jsonString, `"name":"John Doe"`) {
		t.Errorf("Le champ `name` est absent ou incorrect : %s", jsonString)
	}

	// Vérifie que le champ `age` est présent et vaut 0 (valeur zéro non ignorée dans cet exemple)
	if strings.Contains(jsonString, `"age":0`) {
		t.Errorf("Le champ `age` ne devrait pas être présent : %s", jsonString)
	}

	// Vérifie que le champ `address` est présent
	if !strings.Contains(jsonString, `"address":"123 Main St"`) {
		t.Errorf("Le champ `address` est absent ou incorrect : %s", jsonString)
	}

	// Vérifie que le champ `email_address` est absent en raison de `omitempty`
	if strings.Contains(jsonString, `"email_address"`) {
		t.Errorf("Le champ `email_address` ne devrait pas être présent : %s", jsonString)
	}

	// Vérifie que le champ `is_active` est présent et vaut true
	if !strings.Contains(jsonString, `"is_active":true`) {
		t.Errorf("Le champ `is_active` est absent ou incorrect : %s", jsonString)
	}

	// Vérifie que le champ `phone_number` est présent
	if !strings.Contains(jsonString, `"phone_number":"123-456-7890"`) {
		t.Errorf("Le champ `phone_number` est absent ou incorrect : %s", jsonString)
	}

	// Vérifie que le champ `IgnoredField` est absent en raison du tag "-"
	if strings.Contains(jsonString, `"IgnoredField"`) || strings.Contains(jsonString, `"ignored_field"`) {
		t.Errorf("Le champ `IgnoredField` ne devrait pas être présent : %s", jsonString)
	}

	// Vérifie que le champ non exporté `hiddenField` est absent
	if strings.Contains(jsonString, `"hiddenField"`) || strings.Contains(jsonString, `"hidden_field"`) {
		t.Errorf("Le champ `hiddenField` ne devrait pas être présent : %s", jsonString)
	}
}
