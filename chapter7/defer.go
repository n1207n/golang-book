package main

import "fmt"

func main() {
    defer second()
    first()

    /* Another great example of using deter
    f, _ := os.Open(filename)
    defer f.Close()
    */
}

func first() {
    fmt.Println("First")
}

func second() {
    fmt.Println("Second")
}
