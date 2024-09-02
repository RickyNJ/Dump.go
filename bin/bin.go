package bin

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)


type Bin interface {
	Toss(input interface{})
}


func getFileType(filename string) string {
	filename_slice := strings.Split(filename, ".")

	if len(filename_slice) <= 1 {
		return "please add filetype to the filename"
	}

    switch filename_slice[len(filename_slice)-1] {
    case "csv":
        return "csv"
    case "json":
        return "json"
    case "xlsx":
        return "xlsx"
    }

	return "unsupported"
}

func getStructFieldNames[T any](inputStruct T) []string {

	var recursiveSearch func(parentStruct string, input interface{}) []string

	recursiveSearch = func(parentStruct string, input interface{}) []string {

		res := []string{}
		StructType := reflect.TypeOf(input)
		structValue := reflect.ValueOf(input)

		for i := 0; i < StructType.NumField(); i++ {
			field := StructType.Field(i)
			fieldValue := structValue.Field(i)

			var fieldName string

			if parentStruct == "" {
				fieldName = field.Name
			} else {
				fieldName = parentStruct + ":" + field.Name
			}

			if fieldValue.Kind() == reflect.Struct {
				nestedFieldNames := recursiveSearch(fieldName, fieldValue.Interface())
				res = append(res, nestedFieldNames...)
			} else {
				res = append(res, fieldName)
			}
			fmt.Printf("%v with type %v \n\n", field.Name, field.Type.Kind())
		}
		return res
	}

	return recursiveSearch("", inputStruct)
}

func NewBin[T any](fileName string, inputStruct T) Bin {
	structType := reflect.TypeOf(inputStruct)
	if structType.Kind() != reflect.Struct {
		panic("input is not a struct")
	}

	fileType := getFileType(fileName)
	fields := getStructFieldNames(inputStruct)
    fmt.Printf("bintype: %v, extracted fields: %v", fileType, fields)
    if len(fields) == 0 {
        panic("the struct has no fields")
    }

	switch fileType {
	case "csv":
		err := createCSV(fileName, fields)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created file %v \n", fileName)
		return &CSVBin{
			StructType: structType,
			Fields:     fields,
			FilePath:   fileName,
		}

	case "json":
		err := createJSON(fileName, structType.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created file %v \n", fileName)
		return &JSONBin{
			StructType: structType,
			Fields:     fields,
			FilePath:   fileName,
		}
    
    case "xlsx":
        err := createXLSX(fileName, structType.Name(), fields)
        if err != nil {
            log.Fatal(err)
        }
		fmt.Printf("created file %v \n", fileName)
        return &XLSXbin{
            StructType: structType,
            SheetName: structType.Name(),
            Fields: fields,
            FilePath: fileName,

        }
	}

	return nil 
}
