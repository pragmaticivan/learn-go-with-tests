package main

import "testing"

func TestHello(t *testing.T) {
	data := Hello("Ivan")
	expected := "Hello, Ivan"

	if data != expected {
		t.Errorf("got %q want %q", data, expected)
	}

}
