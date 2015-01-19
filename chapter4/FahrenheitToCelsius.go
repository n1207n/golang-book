package main

import "fmt"

func main() {
    fmt.Println("Fahrenheit -> Celsius Converter\n\nCreated by Silin ;p\n")

    var (
        fahrenheit float64
        celsius float64
    )

    fmt.Print("Enter the temperature in Fahrenheit: ")
    fmt.Scanf("%f", &fahrenheit)

    celsius = convertFahrenheitToCelsius(fahrenheit)
    fmt.Printf("The %1.1f F is %1.1f C\n\n", fahrenheit, celsius)
}

func convertFahrenheitToCelsius(value float64) float64 {
    return (value - 32) * 5/9
}
