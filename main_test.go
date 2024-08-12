package main

import (
    "testing"
    "reflect"
)


func TestPrintNames(t *testing.T) {
    test := &Person{Name: "Ricky", Age: 23}
    ans := Printnames(test)
    want := []string{"Name", "Age"}

    if !reflect.DeepEqual(ans, want) {
        t.Fatalf("oh oh")
    }
}
