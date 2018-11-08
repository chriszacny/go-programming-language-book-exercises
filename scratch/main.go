package main

import "fmt"

func recursivePrintInts(current int, max int) {
	if current <= max {
		fmt.Println(current)
		current++
		recursivePrintInts(current, max)
	}
}

func main() {
	recursivePrintInts(1, 100)
}
