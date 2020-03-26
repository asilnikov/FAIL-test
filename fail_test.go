package main

import "testing"

func TestHelloWorld(t *testing.T) {
	iftest helloworld() != "fail test!!" {
		t.Fatal("Test fail")
	}
}
