package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() uint64 {
	var prev uint64 = 0
	var prevPrev uint64 = 0
	var tmp uint64 = 0
	return func() uint64 {
		if prev == 0 {
			prev = 1
			return prev
		}
		tmp = prev
		prev = prev + prevPrev
		prevPrev = tmp
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 93; i++ {
		fmt.Println(f())
	}
}
