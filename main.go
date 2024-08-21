package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

//  TODO:
//  Bin.Toss( data T | []T ) -> if reflect.TypeOf(data) == reflect.TypeOf(self.Data)
//  if data T -> func writeline[[]strings, os.File(?)]
//  if data []T -> while len(data) != 0 -> writeline(data), data.pop?
//  ITERATOR?????????
//  find way of io.Writer to use streams

//  CAST EVERYTHING TO STRING
//  maybe think about ways to make modular so that other formats are possible?

// total funcs for now
// NewBin, Toss, getStructFieldNames, createFile,

type Bin struct {
    headers []string
    filepath string
    filetype string
}

func getFileType(filename string ) string{
    filename_slice := strings.Split(filename, ".")

    if len(filename_slice) <= 1 {
        return "please add filetype to the filename"
    }
    if filename_slice[len(filename_slice)-1] == "csv" {
        return "csv"
    }
    if filename_slice[len(filename_slice)-1] == "json" {
        return "json"
    }

    return "unsupported"
}

func NewBin[T any](filename string, inputStruct T) Bin {
    headers := getStructFieldNames(inputStruct)
    filetype := getFileType(filename)

    if filetype == "csv" {
        binFile, err := CreateFile(filename, headers) 
        if err != nil {
            log.Fatal(err)
        }
        binFile.Close()
    }


    return Bin{headers: headers, filepath: filename, filetype: filetype} 
}

func getStructFieldNames[T any](inputStruct T) []string {
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

func DumpToCSV[T any](fileName string, items []T) bool {
	headers := getStructFieldNames(items)
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
