package main

import (
    "fmt"
    "os"
)

func main() {
    file, openError := os.Open("test.txt")

    if openError != nil {
        return
    }

    defer file.Close()

    stat, statError := file.Stat()

    if statError != nil {
        return
    }

    bs := make([]byte, stat.Size())
    _, readError := file.Read(bs)

    if readError != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}
