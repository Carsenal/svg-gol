package main

import (
    "./gol"
    "fmt"
)

func main() {
    l := gol.NewLife(10, 8)
    l.Current.Set(4, 4, true)
    l.Current.Set(5, 4, true)
    l.Current.Set(6, 4, true)
    fmt.Printf("%s\n", l.ToString())
    l.Step()
    fmt.Printf("%s\n", l.ToString())
    l.Step()
    fmt.Printf("%s\n", l.ToString())
    l.Step()
    fmt.Printf("%s\n", l.ToString())
    l.Step()
    fmt.Printf("%s\n", l.ToString())
    l.Step()
    fmt.Printf("%s\n", l.ToString())
}

