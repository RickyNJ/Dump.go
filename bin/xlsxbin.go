package bin

import (
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

type XLSXbin struct {
	StructType reflect.Type
    Fields []string
    FilePath string
}

func tossXLSX(input interface{}){
    return
}

func (bin *XLSXbin) Toss(input interface{}){
    return 
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

