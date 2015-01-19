package main

import "fmt"

func main() {
    x := 1
    y := 2

    fmt.Println(x, y)

    swap(&x, &y)

    fmt.Println(x, y)
}

func swap(x *int, y *int) {
    temp := *y
    *y = *x
    *x = temp
}
