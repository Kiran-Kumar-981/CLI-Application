package Ticketsent

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestSendTicket(t *testing.T) {
	// Save the original stdout
	oldStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Variables for the test
	userTickets := uint(2)
	firstName := "Surya"
	lastName := "Karna"
	email := "surya.karna@example.com"

	// Call the function
	go SendTicket(userTickets, firstName, lastName, email)

	// Sleep to ensure the function has time to write to stdout
	time.Sleep(3 * time.Second)

	// Restore the original stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Expected output
	expected := "#################\nSending ticket:\n 2 tickets for Surya Karna \nto email address surya.karna@example.com\n#################\nyour show will be held on Monday(dd-mm-yyyy)\n\t\t please come 15-min early for ticket checking\n"

	// Compare the output
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
