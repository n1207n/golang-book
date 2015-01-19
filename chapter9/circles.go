package main

import (
    "fmt"
    "math"
)

type Circle struct {
    // x float64
    // y float64
    // r float64
    // OR
    x, y, r float64
}

func main() {
    var c Circle

    // OR

    newCircle := new(Circle)

    // At this point, c and newCircle are initialized to zero
    // All its member variable are set to zero

    c = Circle {x: 0, y: 0, r: 5}
    fmt.Println(newCircle.x, newCircle.y, newCircle.r)

    fmt.Println(c.x, c.y, c.r)

    *newCircle = Circle {10, 5, 5}

    fmt.Println(newCircle.x, newCircle.y, newCircle.r)

    fmt.Println(circleArea(c))
    fmt.Println(circleAreaWithPointer(newCircle))

    fmt.Println(c.area())
}

func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}

func circleAreaWithPointer(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

// Receiver method for struct
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}
