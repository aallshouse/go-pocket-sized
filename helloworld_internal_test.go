package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := "en"
	expectedGreeting := "Hello, World!"
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", greeting, expectedGreeting)
	}
}

func TestGreet_Spanish(t *testing.T) {
	lang := "es"
	expectedGreeting := "Hola, Mundo!"
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", greeting, expectedGreeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	expectedGreeting := "Bonjour, le monde!"
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", greeting, expectedGreeting)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := "akk"
	expectedGreeting := ""
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", greeting, expectedGreeting)
	}
}
