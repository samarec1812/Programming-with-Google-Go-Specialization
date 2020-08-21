package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var a string
	_, _ = fmt.Scan(&a)
	a = strings.ToLower(a)

	arrRune := []rune(a)
	if arrRune[0] == 'i' && arrRune[utf8.RuneCountInString(a)-1] == 'n' && strings.Contains(a,"a"){
	   	fmt.Print("Found")
    } else {
    	fmt.Print("Not Found")
	}
}