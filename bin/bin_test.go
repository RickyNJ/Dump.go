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

type PersonStrings struct {
    Name string
    Age string
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
func TestToss(t *testing.T) {
    b := NewBin("people.csv", PersonStrings{})
    b.Toss(PersonStrings{Name: "hi", Age: "21"})
    
    file, err := os.Open("people.csv")
    if err != nil {
        t.Errorf("file doesnt exist %v", err)
    }

    r := csv.NewReader(file)

    lines, err := r.ReadAll()
    if err != nil {
        t.Errorf("Couldnt read file %v", err)
    }
    want := [][]string{{"Name", "Age"}, {"hi", "21"}}
    if !reflect.DeepEqual(lines, want){
        t.Fatalf("%v and %v are not equal", lines, want)
    }
}
type ints struct {
    age int
}
func TestTossInt(t *testing.T) {
    b := NewBin("ints.csv", ints{})
    b.Toss(ints{age: 32})
}


type floats struct {
    i float32
    ii float64
}

func TestTossFloat(t *testing.T) {
    b := NewBin("floats.csv", floats{})
    b.Toss(floats{i: 6.1, ii: 3.1 })
}
//
// func TestTossSlice(t *testing.T) {
//     b := NewBin("peoplestrings.csv", PersonStrings{})
//     ra := []PersonStrings{
//         {Name: "ricky", Age: "23"}, 
//         {Name: "alice", Age: "26"},
//         {Name: "klaj", Age: "223"},
//         {Name: "aleliafe", Age: "21"},
//         {Name: "kjndakjlfnace", Age: "516"},
//         {Name: "alicauyfeawfae", Age: "2345"},
//         {Name: "ale", Age: "23456"},
//     }
//     for i := 0; i < 100; i ++ {
//         b.Toss(ra)
//     }
//
//     b.Toss(PersonStrings{Name: "last single one", Age: "234"})
//
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
// func TestUnsupportedFileTypes (t *testing.T) {
//     if getFileType("hai.db") != "unsupported" {
//         t.Fatalf("not unsupported")
//     }
// }
//
// func TestNoFileTypes (t *testing.T) {
//     if getFileType("hai") != "please add filetype to the filename" {
//         t.Fatalf("No File Type failed")
//     }
// }
//
// func TestTooManyPeriodFileTypes (t *testing.T) {
//     if getFileType("hai.lasd.slkdf.json") != "json" {
//         t.Fatalf("Multiple periods failed")
//     }
// }
// func TestCSVFileTypes (t *testing.T) {
//     if getFileType("hai.csv") != "csv" {
//         t.Fatalf("csv failed")
//     }
// }
//
// func TestJSONFileTypes (t *testing.T) {
//     if getFileType("hai.json") != "json" {
//         t.Fatalf("json failed")
//     }
// }

