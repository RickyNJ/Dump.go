package bin

import (
	"fmt"
	"testing"
)

type TestPerson struct {
	Name string
	Age  int
}

type Employee struct {
	Company string
	Person  TestPerson
}

func TestStructToArray(t *testing.T) {
	ricky := Employee{Company: "vfz", Person: TestPerson{Name: "Ricky"}}
	response := structToArray(ricky)
	if response != nil {
		fmt.Print(response...)
	}

}

func TestXLSXCreation(t *testing.T) {
	err := createXLSX("test.xlsx", "people", []string{"Name", "Age"})
	if err != nil {
		panic(err)
	}
}

func TestGetColumn(t *testing.T) {
	var tests = []struct {
		a    int
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

func BenchmarkEXCELTossPerLine(b *testing.B) {
	bin := NewBin("testsingle.xlsx", Person{})
	ricky := Person{Name: "Ricky", Age: 23}

	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		bin.Toss(ricky)
	}
}

func BenchmarkEXCELTossAsArray(b *testing.B) {
	bin := NewBin("testingarray.xlsx", Person{})
	ricky := Person{Name: "Ricky", Age: 23}
	rickys := make([]Person, 100000)

	for i := 0; i < 100000; i++ {
		rickys[i] = ricky
	}

	b.ResetTimer()

	bin.Toss(rickys)
}
