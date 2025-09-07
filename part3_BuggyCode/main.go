package main

import (
	"fmt"
)

//The issue:
// The main function stops before the goroutines start running.
// The goroutines need time to start, so they may begin executing with an unexpected value of I variable.

func main() {
	const N = 5
	done := make([]chan struct{}, N)

	for i := 0; i < N; i++ {
		done[i] = make(chan struct{})
		go func(n int) {
			if n > 0 {
				<-done[n-1]
			}
			fmt.Println(n)
			close(done[n])
		}(i)
	}
	<-done[N-1]
}
