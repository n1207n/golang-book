package main

import "fmt"

func main() {
    fmt.Println("FizzBuzz\n\nCreated by Silin\n")

    i := 1

    for i <= 100 {
        fmt.Print(i, " ")

        if i % 15 == 0 {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println("")
        }

        i++
    }
}
