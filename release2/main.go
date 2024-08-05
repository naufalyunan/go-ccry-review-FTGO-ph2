package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	letters := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}
