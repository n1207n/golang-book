package main

import (
    "os"
)

func main() {
    file, createError := os.Create("new.txt")

    if createError != nil {
        return
    }

    defer file.Close()

    file.WriteString("TEST\n")
}
