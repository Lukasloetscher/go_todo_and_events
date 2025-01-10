package config

import "testing"

func TestInitialise_Settings(t *testing.T) {
	_, err := Initialise_Settings()
	if err != nil {
		t.Error(err)
	}
}
