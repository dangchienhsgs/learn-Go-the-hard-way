package main

import (
	"fmt"
	"sync"
)

func gen(num ...int) <-chan int {
	chan_out := make(chan int)

	go func() {
		for _, n := range num {
			chan_out <- n
		}

		close(chan_out)
	}()

	return chan_out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}

		fmt.Println("Close sq chan")
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}


func main() {
	c := gen(2, 3, 4)


	c1 := sq(c)
	c2 := sq(c)

	for v := range merge(c1, c2) {
		fmt.Println(v)
	}

}
