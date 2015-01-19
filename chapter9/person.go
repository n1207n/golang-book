package main

import "fmt"

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}

func main() {
    p := Person {Name: "Silin Na"}
    p.Talk()

    a := new(Android)
    a.Person.Talk()
    a.Talk()
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
