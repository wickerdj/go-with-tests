package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Elodine in Spanish ", func(t *testing.T) {
		got := Hello("Elodine", "Spanish")
		want := "Hola, Elodine"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Peter in French ", func(t *testing.T) {
		got := Hello("Peter", "French")
		want := "Bonjour, Peter"
		assertCorrectMessage(t, got, want)
	})

}
