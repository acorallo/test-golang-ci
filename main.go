package main

import "fmt"

func main() {
	fmt.Println("Hello World !!")
	result := Sum(2, 3)
	fmt.Println("result:", result)
}

func Sum(a, b int) int {
	return a + b
}
