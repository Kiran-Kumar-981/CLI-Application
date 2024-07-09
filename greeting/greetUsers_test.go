package greeting

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGreetUsers(t *testing.T) {
	// Save the original stdout
	oldStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Variables for the test
	comedian := "Zakir-khan"
	conferenceName := "Stand-up comedy"
	remainingTickets := uint(200)
	count := uint(50)

	// Call the function
	GreetUsers(comedian, conferenceName, remainingTickets, count)

	// Restore the original stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Expected output
	expected := "Zakir-khan welcomes you to the show\nWelcome to Stand-up comedy booking application\nWe have total of 250 tickets and 200 are still available.\nGet your tickets here to attend\n"

	// Compare the output
	if buf.String() != expected {
		t.Errorf("expected %v, got %v", expected, buf.String())
	}
}
