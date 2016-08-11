package main

import "fmt"

var (
	author  = "@cryptogenesis"
	Version = "0.0.1"
)

const (
	horse    = "animal"
	computer = "electronic"
)

func main() {
	var greeting string = "Hello world!"
	var a, b, c int = 1, 2, 3
	fmt.Println(greeting, a, b, c)

	var name = "Evan Rose"
	var d, e, f = 1, 2.0, "three"
	fmt.Println(name, d, e, f)

	course := "Essential Go"
	x, y, z := 1, 2, 3
	fmt.Println(course, x, y, z)
}
