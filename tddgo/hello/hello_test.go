package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("LSJ", "english")
		want := "Hello, LSJ"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in chinese", func(t *testing.T) {
		got := Hello("LSJ", "Chinese")
		want := "你好， LSJ"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("LSJ", "French")
		want := "Bonjour, LSJ"

		assertCorrectMessage(t, got, want)
	})
}
