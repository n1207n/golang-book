package main

import (
    "fmt"
    "os"
)

func main() {
    dir, openError := os.Open(".")

    if openError != nil {
        return
    }

    defer dir.Close()

    fileInfos, readDirectoryError := dir.Readdir(-1)

    if readDirectoryError != nil {
        return
    }

    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}
