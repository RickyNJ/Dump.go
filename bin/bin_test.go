package bin

import (
	"testing"
)

type Person struct {
	Name string // hello
	Age  int    // will go fmt format this for me
}

type PersonStrings struct {
	Name string
	Age  string
}

type Job struct {
	Person  Person
	Company string
	Salary  int
}

type JobNested struct {
    Job Job
    Name string
}


func TestJSONBinCreation(t *testing.T) {
	filename := "people.json"
	NewBin(filename, Person{})
}

func TestXLSXBinCreation(t *testing.T) {
    filename := "people.xlsx"
    NewBin(filename, JobNested{})
}


// }
//
// func TestStructWithInt(t *testing.T) {
//     b := NewBin("people.csv", Person{})
//
//     b.Toss(Person{Name: "Ricky", Age: 12})
//
//     personlist := []Person{
//         {Name: "ln", Age: 34},
//         {Name: "a;lskjdf", Age: 3244},
//         {Name: "fsdf", Age: 234},
//         {Name: "3254", Age: 1244},
//         {Name: "324n", Age: 32},
//         {Name: "ladsfklajshfn", Age: 4},
//     }
//
//     b.Toss(personlist)
//     b.Toss(personlist)
//
//     b.Toss(Person{Name: "Alice", Age: 26})
// }
//
//
// func TestPrintNamesPerson(t *testing.T) {
//     p, _ := initTests()
//
// 	// test := Person{Name: "Ricky", Age: 23}
// 	want := []string{"Name", "Age"}
//
// 	ans := getStructFieldNames(p)
// 	if !reflect.DeepEqual(ans, want) {
// 		t.Fatalf("oh oh")
// 	}
// }
//
//
func TestUnsupportedFileTypes (t *testing.T) {
    if getFileType("hai.db") != "unsupported" {
        t.Fatalf("not unsupported")
    }
}

func TestNoFileTypes (t *testing.T) {
    if getFileType("hai") != "please add filetype to the filename" {
        t.Fatalf("No File Type failed")
    }
}

func TestTooManyPeriodFileTypes (t *testing.T) {
    if getFileType("hai.lasd.slkdf.json") != "json" {
        t.Fatalf("Multiple periods failed")
    }
}
func TestCSVFileTypes (t *testing.T) {
    if getFileType("hai.csv") != "csv" {
        t.Fatalf("csv failed")
    }
}

func TestJSONFileTypes (t *testing.T) {
    if getFileType("hai.json") != "json" {
        t.Fatalf("json failed")
    }
}
