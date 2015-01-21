package main

import (
    "fmt"
    "errors"
)

func main() {
    err := errors.New("Error message")

    fmt.Println(err)
}
