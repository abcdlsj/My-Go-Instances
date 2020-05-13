package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lsj")
	want := "Hello, Lsj"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
