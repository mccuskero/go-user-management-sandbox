package entity

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	if _, err := NewUser("test@test.com", "", "", "", "", ""); err != nil {
		t.Errorf("Error creating user")
	}

	if _, err := NewUser("", "", "", "", "", ""); err == nil {
		t.Errorf("Should of created an error passing in empty email string")
	}
}
