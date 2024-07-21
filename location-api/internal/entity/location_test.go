package entity

import (
	"testing"
)

func TestLocation_Validate(t *testing.T) {
	locations := []struct {
		location      Location
		expectedError string
	}{
		{Location{1, "", 45.0, 90.0, 10.0, 180.0, 5.0, 100.0, 1629878400}, "deviceId cannot be empty"},
		{Location{2, "123", 91.0, 90.0, 10.0, 180.0, 5.0, 100.0, 1629878400}, "latitude must be between -90 and 90"},
		{Location{3, "123", -91.0, 90.0, 10.0, 180.0, 5.0, 100.0, 1629878400}, "latitude must be between -90 and 90"},
		{Location{4, "123", 90.0, 181.0, 10.0, 180.0, 5.0, 100.0, 1629878400}, "longitude must be between -180 and 180"},
		{Location{5, "123", 90.0, -181.0, 10.0, 180.0, 5.0, 100.0, 1629878400}, "longitude must be between -180 and 180"},
		{Location{6, "123", 90.0, 180.0, -10.0, 180.0, 5.0, 100.0, 1629878400}, "speed cannot be negative"},
		{Location{7, "123", 90.0, 180.0, 10.0, 180.0, -5.0, 100.0, 1629878400}, "accuracy cannot be negative"},
		{Location{8, "123", 90.0, 180.0, 10.0, 361.0, 5.0, 100.0, 1629878400}, "heading must be between 0 and 360 degrees"},
		{Location{9, "123", 90.0, 180.0, 10.0, 180.0, 5.0, -100.0, 1629878400}, "altitude cannot be negative"},
	}

	for _, l := range locations {
		err := l.location.validate()
		if err != nil && err.Error() != l.expectedError {
			t.Errorf("Validate() = %v, want valid = %v", err, l.expectedError)
		}
	}
}
