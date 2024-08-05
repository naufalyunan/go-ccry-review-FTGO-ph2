package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//create two channelss
	chOdd := make(chan int)
	chEven := make(chan int)
	chErr := make(chan error)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter an integer number: ")
	num, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: %v", err)
	}
	num = strings.TrimSpace(num)
	numInt, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid integer input, must be integer: %v", err)
	}

	go func(num int) {
		for i := 1; i <= num; i++ {
			if i <= 20 {
				if i%2 == 0 {
					chEven <- i
				} else {
					chOdd <- i
				}
			} else {
				chErr <- fmt.Errorf("error: number %d is greater than 20", i)
			}
		}
		close(chOdd)
		close(chEven)
	}(numInt)

	fmt.Println("")
	for i := 0; i < numInt; i++ {
		select {
		case oddNum, ok := <-chOdd:
			if ok {
				fmt.Printf("Received an odd number: %d\n", oddNum)
			}
		case evNum, ok := <-chEven:
			if ok {
				fmt.Printf("Received an even number: %d\n", evNum)
			}
		case errStr, ok := <-chErr:
			if ok {
				fmt.Println(errStr)
			}
		}

	}

}
