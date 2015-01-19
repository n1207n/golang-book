package main

import "fmt"

func main() {
    // This stops the execution immediately
    // panic("PANIC")

    defer func() {
        str := recover()
        fmt.Println(str)
    }()

    panic("PANIC")
}
