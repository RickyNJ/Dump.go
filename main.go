package main
import (
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

/*
	For now this is just a quick demo on how you can use dump.go to save nested struct data into an excel file
*/

func main() {
	ricky := Person{Name: "Ricky", Age: 23}
	alice := Person{Name: "Alice", Age: 26}

    dev := Worker{Company: "vfz", Person: ricky}
    artist := Worker{Company: "Olsam", Person: alice}


    b := bin.LoadBin("data/demo.xlsx", Worker{})
    

    b.Toss(dev)
    b.Toss(artist)

    for i := 0; i < 10; i ++ {
        b.Toss([]Worker{dev, artist})
    }
}
