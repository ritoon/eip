package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestEmailNotifierNotify(t *testing.T) {
	emailNotifier := EmailNotifier{EmailAddress: "test@example.com", Name: "Test User"}
	message := "Test Email Notification"

	// Construction de la sortie attendue
	expected := fmt.Sprintf("Envoi d'un email à %s (%s) : %s\n", emailNotifier.Name, emailNotifier.EmailAddress, message)

	// Capture de la sortie réelle en redirigeant vers un buffer temporaire
	output := captureOutput(func() {
		emailNotifier.Notify(message)
	})

	if output != expected {
		t.Errorf("EmailNotifier.Notify() = %q; want %q", output, expected)
	}
}

func TestSMSNotifierNotify(t *testing.T) {
	smsNotifier := SMSNotifier{PhoneNumber: "+1234567890", Carrier: "TestCarrier"}
	message := "Test SMS Notification"

	// Construction de la sortie attendue
	expected := fmt.Sprintf("Envoi d'un SMS à %s (%s) : %s\n", smsNotifier.PhoneNumber, smsNotifier.Carrier, message)

	// Capture de la sortie réelle en redirigeant vers un buffer temporaire
	output := captureOutput(func() {
		smsNotifier.Notify(message)
	})

	if output != expected {
		t.Errorf("SMSNotifier.Notify() = %q; want %q", output, expected)
	}
}

func TestSendNotification(t *testing.T) {
	message := "Generic Notification Message"

	emailNotifier := EmailNotifier{EmailAddress: "test@example.com", Name: "Test User"}
	smsNotifier := SMSNotifier{PhoneNumber: "+1234567890", Carrier: "TestCarrier"}

	// Test pour EmailNotifier
	expected := fmt.Sprintf("Envoi d'un email à %s (%s) : %s\n", emailNotifier.Name, emailNotifier.EmailAddress, message)
	output := captureOutput(func() {
		SendNotification(emailNotifier, message)
	})

	if output != expected {
		t.Errorf("SendNotification() with EmailNotifier = %q; want %q", output, expected)
	}

	// Test pour SMSNotifier
	expected = fmt.Sprintf("Envoi d'un SMS à %s (%s) : %s\n", smsNotifier.PhoneNumber, smsNotifier.Carrier, message)
	output = captureOutput(func() {
		SendNotification(smsNotifier, message)
	})

	if output != expected {
		t.Errorf("SendNotification() with SMSNotifier = %q; want %q", output, expected)
	}
}

// Fonction utilitaire pour capturer la sortie
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
