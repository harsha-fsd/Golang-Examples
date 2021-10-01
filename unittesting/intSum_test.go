package test

import (
	"fmt"
	"os"
	"testing"
)

func add(x int, y int) int {
	return x + y
}

//invoked only once but not before and after each test. Used for setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("Setup stuff goes here. ")
	exitValue := m.Run()
	fmt.Println("Cleanup stuff goes here.")
	os.Exit(exitValue)
}

func Test_add(t *testing.T) {
	sum := add(1, 3)
	if sum != 4 {
		t.Error("wrong sum")
	}

}
