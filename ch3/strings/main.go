package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basenameSimpler("a/b/c.go"))
	fmt.Println(basenameSimpler("c.d.go"))
	fmt.Println(basenameSimpler("abc"))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basenameSimpler(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
