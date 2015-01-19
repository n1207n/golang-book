package main

import "fmt"

func main() {
    fmt.Println("Feet -> Meter Converter\n\nCreated by Silin ;p\n")

    var (
        feet float64
        meter float64
    )

    fmt.Print("Enter the distance in ft: ")
    fmt.Scanf("%f", &feet)

    meter = convertFeetToMeter(feet)

    fmt.Printf("%1.2f ft is %1.2f m\n\n", feet, meter)
}

func convertFeetToMeter(value float64) float64 {
    return 0.3048 * value
}
