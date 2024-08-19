package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)


type Bin struct {
    headers []string
    file *os.File
}

//  TODO:
//  Bin.Toss( data T | []T ) -> if reflect.TypeOf(data) == reflect.TypeOf(self.Data)
//  if data T -> func writeline[[]strings, os.File(?)] 
//  if data []T -> while len(data) != 0 -> writeline(data), data.pop?
//  find way of io.Writer to use streams

//  CAST EVERYTHING TO STRING
//  maybe think about ways to make modular so that other formats are possible?


func NewBin[T any](fileName string, inputStruct T) Bin {
    headers := GetHeaders(inputStruct)

    binFile, err := CreateFile(fileName, headers) 
    if err != nil {
        log.Fatal(err)
    }

    return Bin{headers: headers, file: binFile} 
}

func (b Bin) Toss(data interface{}){

}
    
func GetHeaders[T any](inputStruct T) []string {
	headers := []string{}
	structType := reflect.TypeOf(inputStruct)

    if structType.Kind() == reflect.Ptr {
        structType = structType.Elem()
    }

    if structType.Kind() != reflect.Struct {
        // TODO CREATE ERROR FOR NOT A STRUCT
        return headers
    }

    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        headers = append(headers, field.Name)
    }

	return headers
}

func CreateFile(fileName string, headers []string) (*os.File, error){
	f, err := os.Create(fileName)
    w := csv.NewWriter(f)
	w.Write(headers)
	w.Flush()
    return f, err
}


func WriteFile[T any](f *os.File, items []T ) (*os.File, error){
    return os.Create("f")
}



func DumpToCSV[T any](fileName string, items []T) bool {
	headers := GetHeaders(items)
	f, err  := CreateFile(fileName, headers)
    if err != nil {
        log.Fatalf("Error when creating the file %s", err)
    }
    f.Close()
	return true
}

func main() {
	fmt.Print("hi")
}
