package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="-2"
	fmt.Println(Checker(s))
}

func Checker(s string) int {
	part := strings.Split(s , "-")
	return len(part)
}