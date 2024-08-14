package main

import (
	"os"
	"reflect"
	"testing"
)

type Person struct {
	Name string // hello
	Age  int    // will go fmt format this for me
}

type Classes struct {
	Name   string
	Count  int
	People []*Person
	Leader *Person
}

func TestPrintNamesPerson(t *testing.T) {
	test := []*Person{{Name: "Ricky", Age: 23}, {Name: "Alice", Age: 26}}
	want := []string{"Name", "Age"}

	ans := GetHeaders(test)
	if !reflect.DeepEqual(ans, want) {
		t.Fatalf("oh oh")
	}
}

func TestPrintHeadersClasses(t *testing.T) {
	test := []*Classes{{Name: "A3", Count: 3, People: []*Person{{Name: "Ricky", Age: 23}, {Name: "Alice", Age: 26}}, Leader: &Person{Name: "Pekin", Age: 4}}}
	want := []string{"Name", "Count", "People", "Leader"}

	ans := GetHeaders(test)

	if !reflect.DeepEqual(ans, want) {
		t.Fatalf(" oh no")
	}
}

func TestFileCreation(t *testing.T) {
    CreateFile("test")
     
    if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
        t.Fatalf("file doesnt exist: %s", err)	
    }

    err := os.Remove("test.csv") 
    if err != nil { 
        t.Fatalf("Failed to remove the file: %v", err)
    }

}
