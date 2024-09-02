package bin

import (
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

type XLSXbin struct {
	StructType reflect.Type
    SheetName string
    Fields []string
    FilePath string
}
func getColumn(i int) string {
    var result string
    upper, lower := i / 26, i % 26
    lowerrune := rune(lower + 65)

    if upper != 0 {
        upperrune := rune(upper + 64) 
        result = string([]rune{upperrune, lowerrune})
    } else {
        result = string(lowerrune)
    }
    return result
}

func structToArray(input interface{}) []interface{} {
    v := reflect.ValueOf(input)
    values := make([]interface{}, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        values[i] = v.Field(i).Interface()
    }

    return values
}

func tossXLSX(f *excelize.StreamWriter, input interface{}){
    // inputstruct := structToArray(input)
    // w.SetRow("A2", inputstruct) 
    // fmt.Println(inputstruct...)
    // w.Flush()
    inputStruct := structToArray(input)
    for i:= 0; i < len(inputStruct); i++ {
        cell := getColumn(i) + "2" 
        value := inputStruct[i]
        sheet := reflect.TypeOf(input).Name()

        fmt.Printf("Sheet %v, setting value: %v, in cell: %v \n",sheet,  value, cell)
        err := f.SetCellValue(sheet, cell, value)
        if err != nil {
            fmt.Println(err)
        }
    }
    return
}

func (bin *XLSXbin) Toss(input interface{}){
    // open file, find lowest row without and input 
    f, err := excelize.OpenFile(bin.FilePath)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f.GetCellValue("Person", "A1"))

    sw, err := f.NewStreamWriter(bin.SheetName)
    if err != nil {
        fmt.Println(err)
    }




    t := reflect.TypeOf(input)
    switch t.Kind() {
    case reflect.Array, reflect.Slice:
        s := reflect.ValueOf(input)
        for i := 0; i < s.Len(); i++ {
            tossXLSX(f, s.Index(i))
        }
    case reflect.Struct:
        tossXLSX(f, input)
    } 
    f.SaveAs(bin.FilePath)
    return 
}

func createXLSX(filename string, structname string, fields []string) error{
    f := excelize.NewFile()

    err := f.SetSheetName("Sheet1", structname)
    if err != nil {
        fmt.Println(err)
    }

    for i := 0; i < len(fields); i++ {
        column := getColumn(i) + "1"
        fmt.Println(column, fields[i])
        f.SetCellValue(structname, column, fields[i])
    }
    saveerr := f.SaveAs(filename)

    return saveerr 
}

