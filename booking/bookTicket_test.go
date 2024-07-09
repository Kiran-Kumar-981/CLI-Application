package booking

import (
	"testing"
)

func TestBookTicket(t *testing.T) {
	conferenceName := "Stand-up comedy"
	remainingTickets := uint(250)
	firstName := "surya"
	lastName := "karna"
	email := "suryakarna@example.com"
	userTickets := uint(2)

	BookTicket(userTickets, firstName, lastName, email, conferenceName)

	if len(bookings) != 1 {
		t.Errorf("expected %v bookings, got %v", 1, len(bookings))
	}

	if remainingTickets != 248 {
		t.Errorf("expected %v remaining tickets, got %v", 248, remainingTickets)
	}
}

func TestGetFirstNames(t *testing.T) {
	firstNames := GetFirstNames()

	if len(firstNames) != 1 || firstNames[0] != "surya" {
		t.Errorf("expected %v, got %v", []string{"surya"}, firstNames)
	}
}
