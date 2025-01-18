package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// TestSkip skips the current test if the current operating system is macOS
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Skipping test on macOS")
	}

	result := HelloWorld("John")
	assert.Equal(t, "Hello, John", result, "Expected 'Hello, John'")
}

// TestHelloWorld unit test with standard go library
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

// TestHelloWorldAssert unit test using Testify to perform assertions
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("John")
	assert.Equal(t, "Hello, John", result, "Expected 'Hello, John'")
}

// TestHelloWorldRequire unit test using Testify to perform assertions and stops execution if assertion fails
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("John")
	//invoke t.FailNow() if assertion fails
	require.Equal(t, "Hello, John", result, "Expected 'Hello, John'")
}
