package main

import (
	"github.com/RickyNJ/dump/bin"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Position string
	Salary   int
	Person   Person
}

type Nested struct {
    words string
    Employee Employee
}

type Pokemon struct {
    Name string
    Type string
    weight int
}

type Trainer struct {
    Person Person
    Pokemon []Pokemon
    Champion bool
}

func main() {
	ricky := Person{Name: "Ricky", Age: 23}
	alice := Person{Name: "Alice", Age: 26}

	developer := Employee{Position: "developer", Salary: 12, Person: ricky}
	artist := Employee{Position: "artist", Salary: 12, Person: alice}

    teststruct := Nested{words: "hi there", Employee: developer}

	people := []Employee{developer, artist}

	csvbin := bin.NewBin("testnested.csv", Employee{})

	csvbin.Toss(developer)
	csvbin.Toss(people)


    b := bin.NewBin("nestedcsv.csv", Nested{})

    b.Toss(teststruct)
    b.Toss([]Nested{teststruct, teststruct, teststruct})

    pokemonbin := bin.NewBin("pokemontrainers.csv", Trainer{})

    mimikyu := Pokemon{Name: "Mimikyu", Type: "Ghost/Fairy", weight: 5}
    Gengar := Pokemon{Name: "Gengar", Type: "Ghost/Poison", weight: 100}

    trainerRicky := Trainer{Person: ricky,Pokemon: []Pokemon{Gengar, mimikyu},Champion: true}
    trainerAlice := Trainer{Person: alice,Pokemon: []Pokemon{Gengar, mimikyu},Champion: false}



    pokemonbin.Toss(trainerAlice)
    pokemonbin.Toss(trainerRicky)

    pokemonbin.Toss([]Trainer{trainerRicky, trainerRicky, trainerAlice})
    jsonbin := bin.NewBin("jsonbin.json", Person{})


    jsonbin.Toss(ricky)
}

