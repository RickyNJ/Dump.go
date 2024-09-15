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
