package main

import "testing"

func TestGetHelloString(t *testing.T) {
	if getHelloString() != "Hello Go!" {
		t.Errorf("Hello string is not \"Hello Go!\": %s", getHelloString())
	}
}
