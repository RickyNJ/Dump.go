package bin

import (
	"encoding/csv"
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
    structname string
    headers []string
    filepath string
    filetype string
}

func tossCSV (input interface{}) {
    return
}

func (bin *Bin) Toss (input interface{}) {
    if reflect.TypeOf(input).Name() != bin.structname {
        return 
    }

    if bin.filetype == "csv" {
        tossCSV(input) 
    }
    return 
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

func NewBin[T any](filename string, inputStruct T) *Bin {

    v := reflect.ValueOf(inputStruct)

    for v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {
        panic("input is not a struct")
    }

    structname := reflect.TypeOf(inputStruct).Name()
    headers := getStructFieldNames(inputStruct)
    filetype := getFileType(filename)

    if filetype == "csv" {
        binFile, err := CreateFile(filename, headers) 
        if err != nil {
            log.Fatal(err)
        }
        binFile.Close()
    }

    return &Bin{
            structname: structname, 
            headers: headers, 
            filepath: filename, 
            filetype: filetype,
        } 
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

