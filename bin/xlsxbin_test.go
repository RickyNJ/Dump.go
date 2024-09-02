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

func TestGetColumn(t *testing.T) {
    var tests = []struct{
        a int
        want string
    }{
        {0, "A"},
        {25, "Z"},
    }

    for _, tt := range tests {
        testname := tt.want
        t.Run(testname, func(t *testing.T) {
            ans := getColumn(tt.a)
            if ans != tt.want {
                t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}


func TestXLSXBinCreation(t *testing.T) {
    NewBin("test.xlsx", Person{})
}

func TestXLSXBinSingleToss(t *testing.T) {
    b := NewBin("test.xlsx", Person{})
    b.Toss(Person{Name: "ricky", Age: 23})
}

func TestXLSXBinMultipleToss(t *testing.T) {
    b := NewBin("multipletest.xlsx", Person{})
    b.Toss(Person{Name: "ricky", Age: 23})
    b.Toss(Person{Name: "ricky", Age: 23})
}
