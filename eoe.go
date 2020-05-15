package main

import (
	"fmt"
	"os"
)

func main() {
	src, err := os.Open("filename")
	// defer src.Close()

	if err != nil {
		return
	}

	// defer src.Close()

	src.WriteString("hello")

	// defer src.Close()

	
}