package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//A map maps keys to values
var y map[string]Vertex

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m)
	fmt.Println(&m)
	fmt.Printf("%p\n", &m)

	//The zero value of a map is nil. A nil map has no keys, nor can keys be added
	fmt.Println(y)
}
