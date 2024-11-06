package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	// Champ exporté qui sera inclus dans le JSON avec le même nom
	Name string `json:"name"`

	// Champ non exporté (non inclus dans le JSON car il commence par une minuscule)
	hiddenField string `json:"hidden_field"`

	// Champ avec omitempty : sera omis du JSON s'il est vide
	Age int `json:"age,omitempty"`

	// Champ avec un nom personnalisé
	Address string `json:"address,omitempty"`

	// Champ avec le tag "-" : sera toujours ignoré dans le JSON
	IgnoredField string `json:"-"`

	// Champ avec un nom personnalisé et omitempty
	Email string `json:"email_address,omitempty"`

	// Champ optionnel avec une valeur par défaut pour les booléens (true par défaut)
	IsActive bool `json:"is_active"`

	// Champ optionnel de type pointeur : sera omis s'il est nul
	PhoneNumber *string `json:"phone_number,omitempty"`
}

func main() {
	phone := "123-456-7890"
	example := User{
		Name:         "John Doe",
		hiddenField:  "This is hidden",
		Age:          0, // Champ vide, omis grâce à `omitempty`
		Address:      "123 Main St",
		IgnoredField: "This will not appear",
		Email:        "", // Champ vide, omis grâce à `omitempty`
		IsActive:     true,
		PhoneNumber:  &phone, // Inclus car non nul
	}

	// Sérialisation en JSON
	data, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation en JSON:", err)
		return
	}

	// Affichage du JSON
	fmt.Println("JSON sérialisé :")
	fmt.Println(string(data))
}
