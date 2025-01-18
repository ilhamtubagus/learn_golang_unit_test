package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("John")
	if result != "Hello, John" {
		//t.Fail() failed test without aborting code execution afterwards
		//t.FailNow() failed test and stops further execution
		//t.Error() logs an error message and continues execution
		//t.Fatal() logs an error message and stops execution
		t.Error("Expected 'Hello, John', got", result)
	}
}
