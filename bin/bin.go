package bin

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
    structType reflect.Type
    headers []string
    filePath string
    fileType string
}
 
func tossSliceCSV (w *csv.Writer, input interface{}) {
   return 
}

func tossCSV (bin *Bin, w *csv.Writer, input reflect.Value) {
    newLine := []string{}

    for _, v := range bin.headers {
        fmt.Println(input.FieldByName(v))
        newLine = append(newLine, input.FieldByName(v).String())
    }

    fmt.Println(newLine)


    w.Write(newLine)
    w.Flush()
}

func (bin *Bin) Toss (input interface{}) {
    f, err := os.Open(bin.filePath)
    if err != nil {
        panic("Couldnt open file")
    }

    v := reflect.TypeOf(input)
    b := bin.structType
    w := csv.NewWriter(f)

    if v == b {
        fmt.Print("the types are the same running tosscsv")
        tossCSV(bin, w, reflect.ValueOf(input))
    }
    if v.Kind() == reflect.Slice {
        fmt.Print("why is it running this ")
        if v.Elem() == b {
            tossSliceCSV(w, input)
        }
    }
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





func createFile(fileName string, headers []string) (*os.File, error){
	f, err := os.Create(fileName)
    w := csv.NewWriter(f)
	w.Write(headers)
	w.Flush()
    return f, err
}

func getStructFieldNames[T any](inputStruct T) []string {
	headers := []string{}
	structType := reflect.TypeOf(inputStruct)

    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        headers = append(headers, field.Name)
    }

	return headers
}

func NewBin[T any](fileName string, inputStruct T) *Bin {
    t := reflect.TypeOf(inputStruct)
    if t.Kind() != reflect.Struct {
        panic("input is not a struct")
    }

    structType := reflect.TypeOf(inputStruct)
    headers := getStructFieldNames(inputStruct)
    fileType := getFileType(fileName)

    if fileType == "csv" {
        binFile, err := createFile(fileName, headers) 
        if err != nil {
            log.Fatal(err)
        }
        binFile.Close()
    }

    return &Bin{
            structType: structType, 
            headers: headers, 
            filePath: fileName, 
            fileType: fileType,
        } 
}
