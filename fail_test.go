package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "fail test!!" {
		t.Fatal("Test fail")
	}
}
