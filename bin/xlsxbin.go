package bin

import "reflect"

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

func createXLSX(filename string, fields []string) error{
    return nil
}

