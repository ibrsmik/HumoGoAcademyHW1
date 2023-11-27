package main
import ( "fmt"
 "math")
func  array(arr []int, n int) int {
 closest := arr[0]
  c := int(math.Abs(float64(arr[0] - n)))
 for _, num := range arr {
  b := int(math.Abs(float64(num - n)))
    if b < c {
   c = b   
   closest = num
  } }
 return closest
}
func main() {
 var N int
 fmt.Scanln(&N)
arr := make([]int, N)
 for i := 0; i < N; i++ {  fmt.Scan(&arr[i])
 }

 var n int 
 fmt.Scanln(&n)


 result := array(arr, n)
 fmt.Println(result)
}


// package main
// import ( "fmt"
//  "math")
// func main() {
//   var N int
//  arr := make([]int, N)
//  for i := 0; i < N; i++ { 
// 	 fmt.Scan(&arr[i])
//  }

//  var x int
//   fmt.Scan(&x)

// c := arr[1]
// min := int(math.Abs(float64(arr[0] - x)))

//  for _, num := range arr { 
// 	d := int(math.Abs(float64(num - x)))
//   if d < min {   
// 	min = d
//     c = num  }
//  }
//  fmt.Println(c)
// }