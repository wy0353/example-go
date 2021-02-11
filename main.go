package main

import (
	"fmt"
	"strings"
)


func lenAndUpper(word string) (int, string) {
	return len(word), strings.ToUpper(word)
}

func main() {
	len, _ := lenAndUpper("liner")

	fmt.Println(len)
}