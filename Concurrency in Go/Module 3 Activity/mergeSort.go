package main

import (
   "fmt"
   "sort"
   "sync"
)

func sortSlice(arr []int, wg *sync.WaitGroup)  {
   sort.Ints(arr)
   fmt.Println(arr)
   wg.Done()

}

func merge(left []int,right []int) []int{
   //Merge two lists in ascending order.
   lst:=make([]int,0)
   for len(left) > 0 && len(right) > 0{
      if left[0] < right[0]{
         lst = append(lst,left[0])
         left = left[1:]
      }else{
         lst = append(lst,right[0])
         right = right[1:]
      }
   }
   if len(left) > 0{
      lst = append(lst,left...)
   }
   if  len(right) > 0{
      lst = append(lst,right...)
   }

   return lst
}

func createArr(arr []int) []int{
   var wg sync.WaitGroup
   wg.Add(4)
   go sortSlice(arr[:len(arr)/4], &wg)
   go sortSlice(arr[len(arr)/4:len(arr)/2], &wg)
   go sortSlice(arr[len(arr)/2: len(arr)/4 * 3], &wg)
   go sortSlice(arr[len(arr)/4*3:], &wg)
   wg.Wait()
   f1 :=  merge(arr[:len(arr)/4], arr[len(arr)/4:len(arr)/2])
   f2 :=  merge(arr[len(arr)/2: len(arr)/4 * 3], arr[len(arr)/4*3:])
   f3 := merge(f1, f2)
   return f3
}

func main() {
   var n int
   fmt.Println("Array size")
   _, _ = fmt.Scan(&n)
   arr := make([]int, n)
   for i := 0; i < n; i++ {
   	 _, _ = fmt.Scan(&arr[i])
   }
   f := createArr(arr)
   fmt.Print(f)

}
