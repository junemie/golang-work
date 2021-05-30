package main

import "fmt"

func main() {
	var x int = 41
	var y string = "James Bond"
	var z bool = true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

/**
The compiler assigned values to the variables. What are these values called?
They are called zero value. var x int declares the variable "x" and is type int. It assigns the ZERO value type int to "x"
**/
