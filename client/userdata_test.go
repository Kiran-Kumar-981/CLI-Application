package client

import (
	"os"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	input := "Surya\nKarna\nsurya.karna@example.com\n2\n"
	expectedFirstName := "Surya"
	expectedLastName := "Karna"
	expectedEmail := "surya.karna@example.com"
	expectedTickets := uint(2)

	// Use a pipe to simulate os.Stdin
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	// Write input to the pipe
	w.Write([]byte(input))
	w.Close()

	// Save the original os.Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Replace os.Stdin with the read end of the pipe
	os.Stdin = r

	firstName, lastName, email, userTickets := GetUserInput()

	if firstName != expectedFirstName || lastName != expectedLastName || email != expectedEmail || userTickets != expectedTickets {
		t.Errorf("expected (%v, %v, %v, %v), got (%v, %v, %v, %v)", expectedFirstName, expectedLastName, expectedEmail, expectedTickets, firstName, lastName, email, userTickets)
	}
}
