package main

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "test@invalid"
	err := isValidEmail(validEmail)
	if err != nil {
		t.Errorf("Expected no error for valid email, got %v", err)
	}

	err = isValidEmail(invalidEmail)
	if err == nil {
		t.Errorf("Expected error for invalid email, got nil")
	}
}

func TestIsStrongPassword(t *testing.T) {
	validPassword := "StrongP@ssw0rd"
	invalidPassword := "weakpass"
	err := isStrongPassword(validPassword)
	if err != nil {
		t.Errorf("Expected no error for strong password, got %v", err)
	}

	err = isStrongPassword(invalidPassword)
	if err == nil {
		t.Errorf("Expected error for weak password, got nil")
	}
}

func TestContainsNumbersOrSymbols(t *testing.T) {
	containsNumbers := "password123"
	containsSymbols := "password@#$"
	validUsername := "username"
	if containsNumbersOrSymbols(containsNumbers) {
		t.Errorf("Expected false for username containing numbers, got true")
	}
	if containsNumbersOrSymbols(containsSymbols) {
		t.Errorf("Expected false for username containing symbols, got true")
	}
	if !containsNumbersOrSymbols(validUsername) {
		t.Errorf("Expected true for valid username, got false")
	}
}
