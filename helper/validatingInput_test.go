package helper

import "testing"

func TestValidateUserInput(t *testing.T) {
	tests := []struct {
		firstName        string
		lastName         string
		email            string
		userTickets      uint
		remainingTickets uint
		expectedName     bool
		expectedEmail    bool
		expectedTickets  bool
	}{
		{"Surya", "Karna", "surya.karna@example.com", 2, 250, true, true, true},
		{"S", "K", "surya.karna@example.com", 2, 250, false, true, true},
		{"Surya", "Karna", "suryakarnaexample.com", 2, 250, true, false, true},
		{"Surya", "Karna", "surya.karna@example.com", 0, 250, true, true, false},
		{"Surya", "Karna", "surya.karna@example.com", 2, 1, true, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.firstName, func(t *testing.T) {
			isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(tt.firstName, tt.lastName, tt.email, tt.userTickets, tt.remainingTickets)
			if isValidName != tt.expectedName {
				t.Errorf("expected %v, got %v", tt.expectedName, isValidName)
			}
			if isValidEmail != tt.expectedEmail {
				t.Errorf("expected %v, got %v", tt.expectedEmail, isValidEmail)
			}
			if isValidTicketNumber != tt.expectedTickets {
				t.Errorf("expected %v, got %v", tt.expectedTickets, isValidTicketNumber)
			}
		})
	}
}
