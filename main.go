package main

import (
	// "fmt"
	// "github.com/xuri/excelize/v2"

    "github.com/RickyNJ/dump/bin"
)

type Person struct {
	Name string
	Age  int
}

type Pokemon struct {
    Name string
    Type string
    weight int
}

type Trainer struct {
    Person Person
    Pokemon Pokemon
    Champion bool
}

func main() {
	ricky := Person{Name: "Ricky", Age: 23}
	alice := Person{Name: "Alice", Age: 26}

    // b := bin.NewBin("person.csv", Person{})
    //
    // b.Toss(ricky)
    // b.Toss(alice)
    //
    // c := bin.LoadBin("person.csv", Person{})
    //
    // c.Toss(ricky)


    b := bin.LoadBin("person.xlsx", Person{})

    b.Toss(ricky)
    b.Toss(alice)

}

