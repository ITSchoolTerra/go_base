package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("09File/file.png"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
	file, err := os.ReadFile("09File/file.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}
