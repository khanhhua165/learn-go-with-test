package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Khanh", "English")
		want := "Hello, Khanh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to World", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Khanh", "Spanish")
		want := "Hola, Khanh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Khanh", "French")
		want := "Bonjour, Khanh"
		assertCorrectMessage(t, got, want)
	})
}
