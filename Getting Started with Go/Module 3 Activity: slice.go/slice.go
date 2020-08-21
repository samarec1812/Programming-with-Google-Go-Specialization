package main

import (
	"fmt"
	"sort"
)

func main () {
    arr := make([]int, 3);
    var n int = 0
    for {
		_, err := fmt.Scan(&n);
		if err != nil{
			break
		}
    	arr = append(arr, n)
		sort.Ints(arr)
		fmt.Print(arr)

	}
}