package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed file.txt
var message string

func TestEmbed(t *testing.T) {

	// assert.Equal(t, message, "Hello World")
	fmt.Print(message)
	fmt.Print("Program Selesai")
}
