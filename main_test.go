package main

import (
	// "os"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"testing"

)

type Person struct {
	Name string // hello
	Age  int    // will go fmt format this for me
}

type longstruct struct {
    Name string
    Age int
    Gender string
    Phonenumber int 
    Employed string
    Salary int
}

func initTests() (Person, []*Person) {
    return Person{Name: "Ricky", Age: 23},  []*Person{{Name: "Ricky", Age: 23}, {Name: 
    "Alice", Age: 26}}
}


func TestNewBinCreation(t *testing.T) {
    NewBin("people.csv", Person{})

    f, err  := os.Open("people.csv") 
    if err != nil {
        fmt.Printf("people.csv is not created: %v", err)
    }

    r := csv.NewReader(f)
    record, err := r.Read()
    if err != nil {
        fmt.Println(err)
    }

    if !reflect.DeepEqual(record, []string{"Name", "Age"}){
        t.Fatalf("csv content not correct")
    }

    os.Remove("people.csv")
}


func TestPrintNamesPerson(t *testing.T) {
    p, _ := initTests()

	// test := Person{Name: "Ricky", Age: 23}
	want := []string{"Name", "Age"}

	ans := getStructFieldNames(p)
	if !reflect.DeepEqual(ans, want) {
		t.Fatalf("oh oh")
	}
}

// func TestPrintHeadersClasses(t *testing.T) {
// 	test := []*Classes{{Name: "A3", Count: 3, People: []*Person{{Name: "Ricky", Age: 23}, {Name: "Alice", Age: 26}}, Leader: &Person{Name: "Pekin", Age: 4}}}
// 	want := []string{"Name", "Count", "People", "Leader"}
//
// 	ans := getStructFieldNames(test)
//
// 	if !reflect.DeepEqual(ans, want) {
// 		t.Fatalf(" oh no")
// 	}
// }
//
// func TestFileCreation(t *testing.T) {
//     CreateFile("test.csv", []string{"ricky", "alice"})
//      
//     if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
//         t.Fatalf("file doesnt exist: %s", err)	
//     }
//
//     err := os.Remove("test.csv") 
//     if err != nil { 
//         t.Fatalf("Failed to remove the file: %v", err)
//     }
//
// }
