package bin

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
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

func getHigestRowXLSX(fileName string, structName string) int {
    f, err := excelize.OpenFile(fileName)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

    rows, err := f.GetRows(structName)
    if err != nil {
        panic(err)
    }

    return len(rows)
}

func loadCompatibilityXLSX(fileName string, fields []string, structName string) (bool, error) {
    f, err := excelize.OpenFile(fileName)
    if err != nil {
        return false, err
    }

    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

    rows, err := f.GetRows(structName)
    headers := rows[0]

    if !reflect.DeepEqual(headers, fields){
        return false, errors.New("the structfields and xlsx sheet headers are not the same")
    }

    return true, nil
}

func loadCompatibilityCSV(fileName string, fields []string) (bool, error) {
    f, err := os.Open(fileName)
    if err != nil {
        return false, err
    }

    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

    r := csv.NewReader(f)
    headers, err := r.Read()
    if err != nil {
        return false, err 
    }

    if !reflect.DeepEqual(fields, headers) {
        return false, err
    }
    
    return  true, nil
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

    case "xlsx":
        if ok, _ := loadCompatibilityXLSX(fileName, fields, structType.Name()); ok {
            return &XLSXbin{
                StructType: structType,
                SheetName: structType.Name(),
                Fields: fields,
                FilePath: fileName,
                Rows: getHigestRowXLSX(fileName, structType.Name()),
            }
        }
    case "json":
        fmt.Print("Loading json")
    }
    return nil
}

func NewBin[T any](fileName string, inputStruct T) Bin {
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
		err := createCSV(fileName, fields)
		if err != nil {
			log.Fatal(err)
		}
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
        return &XLSXbin{
            StructType: structType,
            SheetName: structType.Name(),
            Fields: fields,
            FilePath: fileName,
            Rows: 1,
        }
	}
	return nil 
}
