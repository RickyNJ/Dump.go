package bin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)


type Bin interface {
	Toss(input interface{})
}

type OptFunc func(*Opts)

type Opts struct {
    timestamp bool
}

func defaultOpts() Opts {
    return Opts{
        timestamp: true,
    }
}

func NewBin[T any](fileName string, inputStruct T, opts ...OptFunc) Bin {
    options := defaultOpts()

	structType := reflect.TypeOf(inputStruct)
	if structType.Kind() != reflect.Struct {
		panic("input is not a struct")
	}

	fields := getStructFieldNames(inputStruct)
    if len(fields) == 0 {
        panic("the struct has no fields")
    }

    if options.timestamp == true {
        fields = append([]string{"timestamp"}, fields...)
    }

	switch getFileType(fileName) {
	case "csv":
		err := createCSV(fileName, fields)
		if err != nil {
			log.Fatal(err)
		}
		return &CSVBin{
			StructType: structType,
			Fields:     fields,
			FilePath:   fileName,
            Options: options,
		}
    
    case "xlsx", "xlam", "xlsm", "xltm", "xltx":
        err := createXLSX(fileName, structType.Name(), fields)
        if err != nil {
            log.Fatal(err)
        }
        return &XLSXbin{
            StructType: structType,
            SheetName: structType.Name(),
            Fields: fields,
            FilePath: fileName,
            Rows: 1,
            Options: options,
        }
	}
	return nil 
}


func getFileType(filename string) string {
	filename_slice := strings.Split(filename, ".")

	if len(filename_slice) <= 1 {
		return "please add filetype to the filename"
	}
    fileType := filename_slice[len(filename_slice)-1]

    switch fileType {
    case "csv", "xlsx", "xlsm", "xlam", "xltm", "xltx":
        return fileType
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


func LoadBin[T any](fileName string, inputStruct T) Bin {
    if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
        panic("bin does not exist")
    }

    structType := reflect.TypeOf(inputStruct)
    if structType.Kind() != reflect.Struct {
        panic("input is not a struct")
    }

    fields := getStructFieldNames(inputStruct) 
    if len(fields) == 0 {
        panic("the struct has no fields")
    }

    switch getFileType(fileName) {
    case "csv":
        if ok, _ := loadCompatibilityCSV(fileName, fields); ok {
            return &CSVBin{
                StructType: structType,
                Fields:     fields,
                FilePath:   fileName,
            }
        }

    case "xlsx", "xlam", "xlsm", "xltm", "xltx":
        if ok, _ := loadCompatibilityXLSX(fileName, fields, structType.Name()); ok {
            return &XLSXbin{
                StructType: structType,
                SheetName: structType.Name(),
                Fields: fields,
                FilePath: fileName,
                Rows: getHigestRowXLSX(fileName, structType.Name()),
            }
        }
    }
    return nil
}
