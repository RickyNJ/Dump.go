package main

import (

    "fmt"
	// "github.com/RickyNJ/dump/bin"
	"github.com/xuri/excelize/v2"
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
	// ricky := Person{Name: "Ricky", Age: 23}
	// alice := Person{Name: "Alice", Age: 26}

    // mimikyu := Pokemon{Name: "Mimikyu", Type: "Ghost/Fairy", weight: 5}
    // Gengar := Pokemon{Name: "Gengar", Type: "Ghost/Poison", weight: 100}
    //
    // trainerRicky := Trainer{Person: ricky,Pokemon: Gengar, Champion: true}
    // trainerAlice := Trainer{Person: alice,Pokemon: mimikyu ,Champion: false}
    //
    //
    // csvbin := bin.NewBin("pokemontrainers.csv", Trainer{})
    // csvbin.Toss(trainerAlice)
    // csvbin.Toss(trainerRicky)
    // csvbin.Toss([]Trainer{trainerRicky, trainerRicky, trainerAlice})

    // xbin := bin.NewBin("Pokemon.xlsx", Person{})
    // xbin.Toss(ricky)


    f := excelize.NewFile()
    err := f.SetSheetName("Sheet1", "Person")
    if err != nil {
        fmt.Println(err)
    }

    f.SetCellValue("Person", "A1", "Name")
    f.SetCellValue("Person", "B1", "Age")

    f.SaveAs("person.xlsx")
    f.Close()

    ff, err := excelize.OpenFile("person.xlsx") 
    if err != nil {
        fmt.Println(err)
    }

    ff.SetCellValue("Person", "A2", "Ricky")
    ff.SetCellValue("Person", "B2", 12)

    // sw, err := ff.NewStreamWriter("Person")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // input := []interface{}{"hi", 12}
    // 
    // sw.SetRow("A2", input)
    // sw.Flush()

    ff.SaveAs("person.xlsx")
    ff.Close()
    
}

