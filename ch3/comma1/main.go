package main

import (
	"bytes"
	"fmt"
)

func commaNonRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer

	var numOfThreePairs int
	if n%3 == 0 {
		numOfThreePairs = n / 3
	} else {
		numOfThreePairs = (n / 3) + 1
	}

	bytesToRead := n % 3
	if bytesToRead == 0 {
		bytesToRead = 3
	}
	currentPosition := 0
	for i := 0; i < numOfThreePairs; i++ {
		for j := 0; j < bytesToRead; j++ {
			buf.WriteByte(s[currentPosition])
			currentPosition++
		}
		if i < numOfThreePairs-1 {
			buf.WriteByte(',')
		}
		bytesToRead = 3
	}

	return buf.String()
}

func main() {
	fmt.Println(commaNonRecursive("96537"))   // Should print: 96,537
	fmt.Println(commaNonRecursive("9653"))    // Should print: 9,653
	fmt.Println(commaNonRecursive("965"))     // Should print: 965
	fmt.Println(commaNonRecursive("965371"))  // Should print: 965,371
	fmt.Println(commaNonRecursive("9653712")) // Should print: 9,653,712
}
