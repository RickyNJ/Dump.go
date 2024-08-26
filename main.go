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

func main() {
	ricky := Person{Name: "Ricky", Age: 23}
	alice := Person{Name: "Alice", Age: 25}

	developer := Employee{Position: "developer", Salary: 12, Person: ricky}
	artist := Employee{Position: "artist", Salary: 12, Person: alice}

    teststruct := Nested{words: "hi there", Employee: developer}

	people := []Employee{developer, artist}

	csvbin := bin.NewBin("testnested.csv", Employee{})

	csvbin.Toss(ricky)

	csvbin.Toss(people)

    b := bin.NewBin("nestedcsv.csv", Nested{})

    b.Toss(teststruct)


    b.Toss([]Nested{teststruct, teststruct, teststruct})
}
