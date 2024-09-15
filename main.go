package main

import (
	"log"
	"time"

	"github.com/RickyNJ/dump/bin"
)

type Person struct {
	Name string
	Age  int
}

type Worker struct {
	Company string
	Person  Person
}

/*
For now this is just a quick demo on how you can use dump.go to save nested struct data into an excel file
*/

func main() {
    
	ricky := Person{Name: "Ricky", Age: 23}
	dev := Worker{Company: "vfz", Person: ricky}

	// alice := Person{Name: "Alice", Age: 26}
	// artist := Worker{Company: "Olsam", Person: alice}

	b := bin.NewBin("test.xlsx", Worker{})

    start := time.Now()
    for i := 0; i < 1000; i ++ {
        b.Toss(dev)
    }
    elapsed := time.Since(start)

    log.Printf("Took %s", elapsed)
}
