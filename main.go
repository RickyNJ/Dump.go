package main

import (
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

    mimikyu := Pokemon{Name: "Mimikyu", Type: "Ghost/Fairy", weight: 5}
    Gengar := Pokemon{Name: "Gengar", Type: "Ghost/Poison", weight: 100}

    trainerRicky := Trainer{Person: ricky,Pokemon: Gengar, Champion: true}
    trainerAlice := Trainer{Person: alice,Pokemon: mimikyu ,Champion: false}


    csvbin := bin.NewBin("pokemontrainers.csv", Trainer{})
    csvbin.Toss(trainerAlice)
    csvbin.Toss(trainerRicky)
    csvbin.Toss([]Trainer{trainerRicky, trainerRicky, trainerAlice})

    xbin := bin.NewBin("Pokemon.xlsx", Trainer{})
    xbin.Toss(trainerAlice)

}

