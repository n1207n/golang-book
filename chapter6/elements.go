package main

import "fmt"

func main() {
    elements := make(map[string]string)

    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements)

    fmt.Println("Un")

    name, ok := elements["Un"]
    fmt.Println(name, ok)

    if name, ok := elements["He"]; ok {
        fmt.Println(name, ok)
    }

    anotherElementMap := map[string]string {
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }

    fmt.Println(anotherElementMap)

    elementWithStateMap := map[string]map[string]string {
        "H": map[string]string {
            "name": "Hydrogen",
            "state": "gas",
        },
    }

    if el, ok := elementWithStateMap["H"]; ok {
        fmt.Println(el["name"], el["state"])
    }
}
