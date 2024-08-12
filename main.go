package main

import (
 "fmt"
 "testing"
 "reflect"
) 



type Person struct {
    Name string
    Age int
}

func Printnames(*Person) []string {
    return []string{"wrong"}
}

func main() {
    fmt.Println("Hi there")
}


func test_Printnames(t *testing.T){
    a := &Person{Name: "Ricky", Age: 24}

    ans := Printnames(a)
    want := []string{"Name", "Age"}

    if reflect.DeepEqual(ans, want){
        t.Error("didnt work")
    }
}
