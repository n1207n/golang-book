package main

import (
    "fmt"
)

type Rectangle struct {
    x1, x2, y1, y2 float64
}

func main() {
    r := Rectangle {0, 0, 10, 10}

    fmt.Println(r.area())
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)

    return l * w
}
