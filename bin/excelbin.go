package bin

import "reflect"

type Excelbin struct {
	StructType reflect.Type
    Fields []string
    FilePath string
}

func tossExcel(input interface{}){
    return
}

func (bin *Excelbin) Toss(input interface{}){
    return 
}

func createExcel(filename string, fields []string) error{
    return nil
}

