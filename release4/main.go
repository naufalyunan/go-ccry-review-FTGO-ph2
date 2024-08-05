package main

import (
	"fmt"
	"time"
)

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	start := time.Now()
	ch := make(chan int, 5)
	go produce(ch)
	consume(ch)
	duration := time.Since(start)
	fmt.Println("Running time: ", duration)
}

//if we scale the function into 1000000 iteration,
//the time needed to run the function is 5.1479277s
//so, using buffered array make the execution time faster
