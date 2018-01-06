package config

import (
	"testing"
)

func TestIsTestUser(t *testing.T) {
	if !IsTestUser("user-fidplzi1ia2q") {
		t.Fatal("user-fidplzi1ia2q should be a test user")
	}
	if IsTestUser("user-fidplzi1ia2") {
		t.Fatal("user-fidplzi1ia2 should not be a test user")
	}
}
