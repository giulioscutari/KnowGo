package main

import "testing"

func AssertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Olly", "")
		want := "Hello, Olly!"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		AssertCorrectMessage(t, got, want)
		
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T){
		got := Hello("Marie", "French")
		want := "Bonjour, Marie!"
		AssertCorrectMessage(t, got, want)
	})
}