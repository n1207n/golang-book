package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, openError := ioutil.ReadFile("test.txt")

    if openError != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}
