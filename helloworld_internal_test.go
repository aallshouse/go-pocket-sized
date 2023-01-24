package main

import "testing"

func TestGreet(t *testing.T) {
	expectedGreeting := "Hello, World!"
	greeting := greet()

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", greeting, expectedGreeting)
	}
}
