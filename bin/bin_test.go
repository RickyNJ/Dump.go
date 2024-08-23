package bin

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type Person struct {
	Name string // hello
	Age  int    // will go fmt format this for me
}

type longstruct struct {
    Name string
    Age int
    Gender string
    Phonenumber int 
    Employed string
    Salary int
}

func initTests() (Person, []*Person) {
    return Person{Name: "Ricky", Age: 23},  []*Person{{Name: "Ricky", Age: 23}, {Name: 
    "Alice", Age: 26}}
}

func TestToss(t *testing.T) {
    b := NewBin("people.csv", Person{})
    
    if !b.Toss(Person{Name: "hi", Age: 21}) {
        t.Fail()
    }
}


func TestNewBinCreation(t *testing.T) {
    filename := "people.csv"
    defer os.Remove(filename)

    NewBin(filename, Person{})

    f, err  := os.Open(filename)
    if err != nil {
        fmt.Printf("%v is not created: %v",filename, err)
    }

    r := csv.NewReader(f)
    record, err := r.Read()
    if err != nil {
        fmt.Println(err)
    }

    if !reflect.DeepEqual(record, []string{"Name", "Age"}){
        t.Fatalf("csv content not correct")
    }

}


func TestPrintNamesPerson(t *testing.T) {
    p, _ := initTests()

	// test := Person{Name: "Ricky", Age: 23}
	want := []string{"Name", "Age"}

	ans := getStructFieldNames(p)
	if !reflect.DeepEqual(ans, want) {
		t.Fatalf("oh oh")
	}
}


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
// func TestPrintHeadersClasses(t *testing.T) {
// 	test := []*Classes{{Name: "A3", Count: 3, People: []*Person{{Name: "Ricky", Age: 23}, {Name: "Alice", Age: 26}}, Leader: &Person{Name: "Pekin", Age: 4}}}
// 	want := []string{"Name", "Count", "People", "Leader"}
//
// 	ans := getStructFieldNames(test)
//
// 	if !reflect.DeepEqual(ans, want) {
// 		t.Fatalf(" oh no")
// 	}
// }
//
// func TestFileCreation(t *testing.T) {
//     CreateFile("test.csv", []string{"ricky", "alice"})
//      
//     if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
//         t.Fatalf("file doesnt exist: %s", err)	
//     }
//
//     err := os.Remove("test.csv") 
//     if err != nil { 
//         t.Fatalf("Failed to remove the file: %v", err)
//     }
//
// }
