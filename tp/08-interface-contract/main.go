package main

import (
	"fmt"
)

// 1. Interface Notifier avec la méthode Notify
type Notifier interface {
	Notify(message string)
}

// 2. Struct EmailNotifier avec les champs EmailAddress et Name
type EmailNotifier struct {
	EmailAddress string
	Name         string
}

// 3. Implémentation de la méthode Notify pour EmailNotifier
func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Envoi d'un email à %s (%s) : %s\n", e.Name, e.EmailAddress, message)
}

// 4. Struct SMSNotifier avec les champs PhoneNumber et Carrier
type SMSNotifier struct {
	PhoneNumber string
	Carrier     string
}

// 5. Implémentation de la méthode Notify pour SMSNotifier
func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Envoi d'un SMS à %s (%s) : %s\n", s.PhoneNumber, s.Carrier, message)
}

// 6. Fonction SendNotification qui utilise l'interface Notifier
func SendNotification(n Notifier, message string) {
	n.Notify(message)
}

func PrintNotificationInfo(n Notifier) {
	// Affiche les informations spécifiques en fonction du type sous-jacent
	switch notif := n.(type) {
	case EmailNotifier:
		fmt.Printf("Notif email %v et nom %v\n", notif.EmailAddress, notif.Name)
	case SMSNotifier:
		fmt.Printf("Notif numéro de tel %v et propriétaire %v\n", notif.PhoneNumber, notif.Carrier)
	default:
		fmt.Println("Type de notification inconnue")
	}
}

func main() {
	// Création d'une liste de notifications avec différents types de Notifier
	notifications := []Notifier{
		EmailNotifier{EmailAddress: "alice@example.com", Name: "Alice"},
		SMSNotifier{PhoneNumber: "+1234567890", Carrier: "Carrier1"},
		EmailNotifier{EmailAddress: "bob@example.com", Name: "Bob"},
		SMSNotifier{PhoneNumber: "+0987654321", Carrier: "Carrier2"},
	}

	// Envoi de notifications en utilisant SendNotification
	for _, notifier := range notifications {
		SendNotification(notifier, "Vous avez une nouvelle notification !")
	}
}
