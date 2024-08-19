package main

import (
	// "os"
	"reflect"
	"testing"
)

type Person struct {
	Name string // hello
	Age  int    // will go fmt format this for me
}

func initTests() (Person, []*Person) {
    return Person{Name: "Ricky", Age: 23},  []*Person{{Name: "Ricky", Age: 23}, {Name: 
    "Alice", Age: 26}}
}


func TestNewBinCreation(t *testing.T) {
    NewBin("people.csv", Person{})
}

func TestPrintNamesPerson(t *testing.T) {
    p, _ := initTests()

	// test := Person{Name: "Ricky", Age: 23}
	want := []string{"Name", "Age"}

	ans := GetHeaders(p)
	if !reflect.DeepEqual(ans, want) {
		t.Fatalf("oh oh")
	}
}

// func TestPrintHeadersClasses(t *testing.T) {
// 	test := []*Classes{{Name: "A3", Count: 3, People: []*Person{{Name: "Ricky", Age: 23}, {Name: "Alice", Age: 26}}, Leader: &Person{Name: "Pekin", Age: 4}}}
// 	want := []string{"Name", "Count", "People", "Leader"}
//
// 	ans := GetHeaders(test)
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
