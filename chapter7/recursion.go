package main

import "fmt"

func main() {
    var x uint = 10

    fmt.Println(factorial(x))
}

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x - 1)
}
