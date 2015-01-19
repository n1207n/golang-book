package main

import "fmt"

func main() {
    x := make([]float64, 5)
    fmt.Println(x)

    x = make([]float64, 5, 10)
    fmt.Println(x)

    array := [5]float64 {1,2,3,4,5}
    fmt.Println(array)

    x = array[0:5]
    fmt.Println(x)

    slice1 := []int {1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)

    slice2 = make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}
