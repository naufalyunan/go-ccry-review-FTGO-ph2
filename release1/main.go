package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	letters := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
}

func main() {
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second)
}
