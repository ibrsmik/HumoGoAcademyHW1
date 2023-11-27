package main

import (
	"fmt"
)
func mainn(arr []int, n int) int {count :=0
	for _, num:= range arr{if num == n{
		count++
	}
	}
	return count 
}
func main()  {
	var N int 
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i< N; i++{
		fmt.Scan(&arr[i])
	}
	var n int
	fmt.Scan(&n)
	result := mainn(arr, n)
	fmt.Println(result)
}

// package main
// import ( "fmt")
// func countOccurrences(arr []int, x int) int { count := 0
//  for _, num := range arr {  if num == x {
//    count++  }
//  } return count
// }
// func main() {

//  fmt.Scanln(&N)

//  arr := make([]int, N) for i := 0; i < N; i++ {
//   fmt.Scan(&arr[i]) }

//  var x int fmt.Scanln(&x)

//  result := countOccurrences(arr, x)
 
// }