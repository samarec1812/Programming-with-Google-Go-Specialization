package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}
func BubleSort(arr []int){
     for i := 0; i < len(arr); i++ {
     	for j := 0; j < len(arr)-1; j++ {
     	       if (arr[j] > arr[j+1]){
     	       	Swap(&arr[j], &arr[j+1])
     	       	}
           }
     }
}

func main() {
      arr := make([]int, 0)
      var number int
      for i := 0; i < 10; i++ {
      	_, _ = fmt.Scan(&number)
      	arr = append(arr, number)
	  }
	  BubleSort(arr)
	  fmt.Print(arr)
}
