package bin

import (
    "testing"
)

func TestXLSXCreation(t *testing.T){
    err := createXLSX("test.xlsx", "people", []string{"Name", "Age"}) 
    if err != nil {
        panic(err)
    }
}
