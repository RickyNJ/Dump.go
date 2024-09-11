package bin

import (
    "testing"
    "os"
    "encoding/csv"
    "fmt"
    "reflect"
)

func isEqual(input string, want [][]string) bool {
    f, err := os.Open(input)
    defer f.Close()
    if err != nil {
        fmt.Println(err)
    }

    r := csv.NewReader(f)
    record, err := r.ReadAll()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("testing for deepequal\ninput: %v \nwant: %v \n", record, want)
    return reflect.DeepEqual(record, want)
}

func TestCSVTossWithNil(t *testing.T) {
    b := NewBin("people.csv", Person{})
    b.Toss(Person{Age: 12})

    if !isEqual("people.csv", [][]string{{"Name", "Age"},{"", "12"}}) {
        panic("not equal")
    }

}
func TestNewCSVBinCreation(t *testing.T) {
	filename := "people.csv"
	NewBin(filename, Person{})
    if !isEqual("people.csv", [][]string{{"Name, Age"}}) {
        panic("not equal")
    }
}


func TestNewCSVBinCreationNestedStruct(t *testing.T) {
	filename := "job.csv"
	NewBin(filename, JobNested{})
    if !isEqual("job.csv",  [][]string{{"Job:Person:Name", "Job:Person:Age", "Job:Company", "Job:Salary", "Name"}}) {
        panic("not equal")
    }
}


func TestCSVToss(t *testing.T) {
	b := NewBin("people.csv", PersonStrings{})
	b.Toss(Person{Name: "hi", Age: 21})

	want := [][]string{{"Name", "Age"}, {"hi", "21"}}
    if !isEqual("people.csv", want) {
        panic("not equal")
    }
}


func TestCSVTossWithSlice(t *testing.T) {
	b := NewBin("peoples.csv", Person{})

	personslice := []Person{{Name: "Ricky", Age: 12}, {Name: "Alice", Age: 23}}
	personarray := [2]Person{{Name: "lkj", Age: 12}, {Name: "lkdsjAlice", Age: 23}}

	b.Toss(personslice)
	b.Toss(personarray)

    want := [][]string{{"Name", "Age"}, {"Ricky", "12"}, {"Alice", "23"}, {"lkj", "12"}, {"lkdsjAlice", "23"}}

    if !isEqual("peoples.csv", want) {
        panic("not equal")
    }
}

type arrstruct struct {
	i []string
}

func TestCSVTossArr(t *testing.T) {
	b := NewBin("arr.csv", arrstruct{})
	b.Toss(arrstruct{i: []string{"hi", "hello"}})
    
    want := [][]string{{"i"}, {"[hi hello]"}}
    if !isEqual("arr.csv", want) {
        panic("not equal")
    }
}

type Nested struct {
    Person Person
    Company string
}

func TestCSVTossNested(t *testing.T){
    // defer os.Remove("nested.csv")
    b := NewBin("nested.csv", Nested{}) 
    ricky := Nested{Person: Person{Name: "ricky", Age: 23}, Company: "vfz"}

    b.Toss(Nested{Person: Person{Name: "ricky", Age: 23}, Company: "vfz"})

    b.Toss([]Nested{ricky, ricky})

	file, err := os.Open("nested.csv")
	if err != nil {
		t.Errorf("file doesnt exist %v", err)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		t.Errorf("Couldnt read file %v", err)
	}

    
    want := [][]string{{"Person:Name", "Person:Age", "Company"}, {"ricky", "23", "vfz"}}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("%v and %v are not equal", lines, want)
	}
}

func BenchmarkTossPerLine(b *testing.B) {
    bin := NewBin("testsingle.csv", Person{})
    ricky := Person{Name: "Ricky", Age: 23}

    b.ResetTimer()

    for i := 0; i < 10000; i++ {
        bin.Toss(ricky)
    }
}

func BenchmarkTossAsArray(b *testing.B) {
    bin := NewBin("testingarray.csv", Person{})
    ricky := Person{Name: "Ricky", Age: 23}
    rickys := make([]Person, 10000)

    for i := 0; i < 10000; i++ {
        rickys[i] = ricky
    }

    b.ResetTimer()

    bin.Toss(rickys)
}
