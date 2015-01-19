package main

import "fmt"

func main() {
    fmt.Println(add(1, 2, 3))

    // We can pass the slice as variadic value
    xs := []int {1, 2, 3}
    fmt.Println(add(xs...))

    // Closure
    x := 0

    increment := func() int {
        x++
        return x
    }

    fmt.Println(increment())
    fmt.Println(increment())
}

func add(args ... int) int {
    total := 0
    for _, v := range args {
        total += v
    }

    return total
}
