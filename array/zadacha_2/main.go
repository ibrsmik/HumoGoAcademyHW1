package main

import "fmt"
func array(arr []int, n int) string { for _, num := range arr {
  if num == n {   return "YES"
  } }
 return "NO"
}

func main() {
	var N int
 fmt.Scanln(&N)
 arr := make([]int, N)

 for i := 0; i < N; i++ { 
	 fmt.Scan(&arr[i])
 }
var n int
fmt.Scanln(&n)
  result := array(arr, n)
 fmt.Println(result)
}