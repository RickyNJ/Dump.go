package bin

import (
    "strconv"
	"fmt"
	"reflect"
	"github.com/xuri/excelize/v2"
)

type XLSXbin struct {
	StructType reflect.Type
    SheetName string
    Fields []string
    FilePath string
    Rows int
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

func tossXLSX(bin *XLSXbin, f *excelize.File, input interface{}){
    inputStruct := structToArray(input)
    bin.Rows += 1
    for i := 0; i < len(inputStruct); i++ {
        cell := getColumn(i) + strconv.Itoa(bin.Rows)
        f.SetCellValue(bin.SheetName, cell,  inputStruct[i])
    }
    return
}

func (bin *XLSXbin) Toss(input interface{}){
    // open file, find lowest row without and input 
    f, err := excelize.OpenFile(bin.FilePath)
    if err != nil {
        fmt.Println(err)
    }

    t := reflect.TypeOf(input)
    switch t.Kind() {
    case reflect.Array, reflect.Slice:
        s := reflect.ValueOf(input)
        for i := 0; i < s.Len(); i++ {
            tossXLSX(bin, f, s.Index(i))
        }
    case reflect.Struct:
        tossXLSX(bin,f, input)
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

    f.SaveAs(filename)
    return err 

}

