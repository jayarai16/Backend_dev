package models

import (
	"testing"
)

func TestCalculateAge(t *testing.T) {
	// Test case: DOB 1990-05-10, current year 2023, age should be 33
	user := &UserWithAge{
		DOB: "1990-05-10",
	}
	user.CalculateAge()
	expected := 35 // Since current date is 2025-12-17, age is 35
	if user.Age != expected {
		t.Errorf("Expected age %d, got %d", expected, user.Age)
	}

	// Test case: DOB 2000-12-17, age should be 24
	user2 := &UserWithAge{
		DOB: "2000-12-17",
	}
	user2.CalculateAge()
	expected2 := 24
	if user2.Age != expected2 {
		t.Errorf("Expected age %d, got %d", expected2, user2.Age)
	}

	// Test case: Invalid DOB
	user3 := &UserWithAge{
		DOB: "invalid",
	}
	user3.CalculateAge()
	if user3.Age != 0 {
		t.Errorf("Expected age 0 for invalid DOB, got %d", user3.Age)
	}
}