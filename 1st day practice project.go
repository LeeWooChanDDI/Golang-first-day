package main

import "fmt"

func main() {
	var a string
	var b string
	var c int

	fmt.Scanln(&a, &b, &c)
	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%d\n", c)
}
