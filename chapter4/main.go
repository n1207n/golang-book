package main

import "fmt"

var globalScopeVariable string = "Hi there"

func main() {
    // var x string = "Hello World"
    // var x string;
    // x = "Hello World"

    var x string

    x = "first "
    fmt.Println(x)

    x = x + "second"
    fmt.Println(x)

    x = "first "
    fmt.Println(x)

    x += "second"
    fmt.Println(x)

    x = "Hello"

    var y string

    y = "Hello"

    fmt.Println(x == y)

    z := "Hello World"
    fmt.Println(z)

    v := 5
    fmt.Println(v)

    max := "Max"
    fmt.Println("My name is", max)

    fmt.Println(globalScopeVariable)

    const constantBaby = "CONST"
    fmt.Println(constantBaby)

    var (
        dog = "Dawgs"
        cat = "Katz"
        sleep = "Sleep"
    )

    fmt.Println(dog, cat, sleep)

    f()
}

func f() {
    fmt.Println(globalScopeVariable)
}
