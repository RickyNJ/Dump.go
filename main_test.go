package main

import (
    "testing"
    "reflect"
)

type Person struct {
    Name string
    Age int
}

type Classes struct {
    Name string
    Count int 
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
