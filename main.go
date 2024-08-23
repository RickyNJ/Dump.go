package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"

	"github.com/RickyNJ/dump/bin"
)

type Person struct {
    Name string
    Age int
}

func main() {
    b := bin.NewBin("people.csv", Person{})
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Name:")
        namescanner := scanner.Scan()
        if !namescanner {
            return
        }
        name := scanner.Text()

        fmt.Print("Age:")
        agescanner := scanner.Scan()
        if !agescanner {
            return
        }
        age := scanner.Text()
        ageint, err := strconv.Atoi(age)
        if err != nil {
            return
        }
        b.Toss(Person{Name: name, Age: ageint})
    }
} 
