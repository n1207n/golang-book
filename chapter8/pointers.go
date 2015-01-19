package main

import "fmt"

func main() {
    x := 5
    zero(&x)

    fmt.Println(x)

    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr)
}

// func zero(x int) {
//     // x is passed by value, thus we need to access its pointer to mutate
//     x = 0
// }

func zero(xPtr *int) {
    *xPtr = 0
}

func one(xPtr *int) {
    *xPtr = 1
}
