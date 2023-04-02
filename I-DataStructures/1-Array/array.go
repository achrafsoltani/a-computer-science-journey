package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	var fruits [5]string
	fruits[0] = "Apples"
	fruits[1] = "Bananas"
	fruits[2] = "Oranges"
	fruits[3] = "Peaches"
	fruits[4] = "Berries"
	fmt.Println("Before: " + fruits[0])
	fruits[0] = "Pineapples"
	fmt.Println("After: " + fruits[0])
}
