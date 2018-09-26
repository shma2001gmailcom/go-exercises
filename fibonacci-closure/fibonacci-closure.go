package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev int = 0
	var prevPrev int = 0
	var tmp = 0
	return func() int {
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
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
