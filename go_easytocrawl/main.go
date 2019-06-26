package main

import (
	"fmt"
	// "reflect"
)

func main() {
	a := "我爱你"
	a_rune := []rune(a)
	for _, value := range a_rune {
		fmt.Println(string(value))
	}
}