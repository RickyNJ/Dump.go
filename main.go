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

type Worker struct {
    Company string
    Person Person
}

func main() {
	ricky := Person{Name: "Ricky", Age: 23}
	alice := Person{Name: "Alice", Age: 26}

    dev := Worker{Company: "vfz", Person: ricky}
    artist := Worker{Company: "Olsam", Person: alice}


    b := bin.LoadBin("data/demo.xlsx", Worker{})
    

    b.Toss(dev)
    b.Toss(artist)

    for i := 0; i < 100; i ++ {
        b.Toss([]Worker{dev, artist})
    }

    for i := 0; i < 100; i++ {
        b.Toss(dev)
    }
}
