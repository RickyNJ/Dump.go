package bin

import (
    "testing"
    "os"
    "encoding/csv"
    "fmt"
    "reflect"
)


func TestNewCSVBinCreation(t *testing.T) {
	filename := "people.csv"
	defer os.Remove(filename)

	NewBin(filename, Person{})

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%v is not created: %v", filename, err)
	}

	r := csv.NewReader(f)
	record, err := r.Read()
	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(record, []string{"Name", "Age"}) {
		t.Fatalf("csv content not correct")
	}

}


func TestNewCSVBinCreationNestedStruct(t *testing.T) {
	defer os.Remove("job.csv")
	filename := "job.csv"

	NewBin(filename, JobNested{})

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%v is not created: %v", filename, err)
	}

	r := csv.NewReader(f)
	record, err := r.Read()
	if err != nil {
		fmt.Println(err)
	}

    if !reflect.DeepEqual(record, []string{"Job:Person:Name", "Job:Person:Age", "Job:Company", "Job:Salary", "Name"}) {
		t.Fatalf("csv content not correct")
	}
}


func TestCSVToss(t *testing.T) {
	b := NewBin("people.csv", PersonStrings{})
	b.Toss(Person{Name: "hi", Age: 21})

	file, err := os.Open("peoples.csv")
	if err != nil {
		t.Errorf("file doesnt exist %v", err)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		t.Errorf("Couldnt read file %v", err)
	}
	want := [][]string{{"Name", "Age"}, {"hi", "21"}}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("%v and %v are not equal", lines, want)
	}
}

func TestCSVTossWithSlice(t *testing.T) {
    defer os.Remove("peoples.csv")
	b := NewBin("peoples.csv", Person{})

	personslice := []Person{{Name: "Ricky", Age: 12}, {Name: "Alice", Age: 23}}
	personarray := [2]Person{{Name: "lkj", Age: 12}, {Name: "lkdsjAlice", Age: 23}}

	b.Toss(personslice)
	b.Toss(personarray)
}

type arrstruct struct {
	i []string
}

func TestCSVTossArr(t *testing.T) {
    defer os.Remove("arr.csv")

	b := NewBin("arr.csv", arrstruct{})
	b.Toss(arrstruct{i: []string{"hi", "hello"}})
}

type Nested struct {
    Person Person
    Company string
}

func TestCSVTossNested(t *testing.T){
    // defer os.Remove("nested.csv")
    b := NewBin("nested.csv", Nested{}) 

    b.Toss(Nested{Person: Person{Name: "ricky", Age: 23}, Company: "vfz"})

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
