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



type CSVBin struct {
    structType reflect.Type
    fields []string
    filePath string
}

type JSONBin struct {
    structType reflect.Type
    fields []string
    filePath string
}

// type Bin struct {
//     structType reflect.Type
//     fields []string
//     filePath string
//     fileType string
// }


type Bin interface {
    Toss(input interface{})
}


func (bin JSONBin) Toss (input interface{}) {
    f, err := os.OpenFile(bin.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic("Couldnt open file")
    }
    f.Close()

    return
}


func tossCSV (bin *CSVBin, w *csv.Writer, input reflect.Value) { 
    newLine := []string{}
    for _, v := range bin.fields {
        value := input.FieldByName(v)
        var newValue string

        newValue = fmt.Sprint(value)
        newLine = append(newLine, newValue)
        // switch value.Kind(){
        // case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        //     newValue = strconv.Itoa(int(value.Int()))
        //
        // case reflect.String:
        //     newValue = value.String()
        // }
    }
    w.Write(newLine)
    w.Flush()
}


func (bin *CSVBin) Toss (input interface{}) {
    f, err := os.OpenFile(bin.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic("Couldnt open file")
    }

    w := csv.NewWriter(f)
    t := reflect.TypeOf(input)
    b := bin.structType

    switch reflect.ValueOf(input).Kind(){
    case reflect.Struct:
        if t == b {
            tossCSV(bin, w, reflect.ValueOf(input))
        }
    case reflect.Slice:
        s := reflect.ValueOf(input)
        for i := 0; i < s.Len(); i++ {
            tossCSV(bin, w, s.Index(i))
        }
    }
}


func createCSV(fileName string, fields []string) error{
	f, err := os.Create(fileName)
    w := csv.NewWriter(f)
    w.Write(fields)
    w.Flush()
    f.Close()
    return err
}

func createJSON(fileName string) error {
    f, err := os.Create(fileName)
    f.Close()
    return err
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


func getStructFieldNames[T any](inputStruct T) []string {
	fields := []string{}
	structType := reflect.TypeOf(inputStruct)

    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        fields = append(fields, field.Name)
    }

	return fields
}

func NewBin[T any](fileName string, inputStruct T) Bin {
    structType := reflect.TypeOf(inputStruct)
    fmt.Printf("Generating new Bin with filename: %v and input struct %v \n", fileName, structType)
    if structType.Kind() != reflect.Struct {
        panic("input is not a struct")
    }

    fields := getStructFieldNames(inputStruct)
    fileType := getFileType(fileName)
    fmt.Printf("extracted file type: %v, fields: %v \n", fileType, fields)
    switch fileType {
    case "csv":
        err := createCSV(fileName, fields) 
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("created file %v \n", fileName)
        return &CSVBin {
            structType: structType,
            fields: fields,
            filePath: fileName,
        }

    case "json":
        err := createJSON(fileName)
        if err != nil {
            log.Fatal(err)
        }
        return &JSONBin{
            structType: structType,
            fields: fields,
            filePath: fileName,
        }
    }

    return nil    
}
