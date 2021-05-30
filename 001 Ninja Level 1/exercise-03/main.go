package main

import "fmt"

var x int = 41
var y string = "James Bond"
var z bool = true

func main() {
	// s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	s := fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(s)
}
