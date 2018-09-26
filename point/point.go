package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	m["me"] = Vertex{
		5, -5,
	}
	fmt.Println(m)
	delete(m, "me")
	fmt.Println(m)
}
