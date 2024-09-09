package bin

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

type CSVBin struct {
	StructType reflect.Type
	Fields     []string
	FilePath   string
}


func tossCSV(w *csv.Writer, input interface{}) {
	newLine := []string{}
	var scanInput func(inputStruct interface{})

	scanInput = func(inputStruct interface{}) {

		inputType := reflect.TypeOf(inputStruct)
		inputValue := reflect.ValueOf(inputStruct)

		if inputValue.Kind() == reflect.Ptr {
			inputValue = inputValue.Elem()
			inputType = inputType.Elem()
		}

		for i := 0; i < inputType.NumField(); i++ {
			fieldValue := inputValue.Field(i)

			if fieldValue.Kind() == reflect.Struct {
				scanInput(fieldValue.Interface())

			} else if fieldValue.Kind() == reflect.Ptr {
				if !fieldValue.IsNil() {
					elem := fmt.Sprint(fieldValue.Elem())
					newLine = append(newLine, elem)
				} else {
					newLine = append(newLine, "nil")
				}

			} else {
				newLine = append(newLine, fmt.Sprint(fieldValue))
			}
		}
	}

	scanInput(input)
	w.Write(newLine)
}

func (bin *CSVBin) Toss(input interface{}) {
	f, err := os.OpenFile(bin.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Couldnt open file")
	}

	w := csv.NewWriter(f)
	t := reflect.TypeOf(input)
	b := bin.StructType

	switch t.Kind() {
	case reflect.Struct:
		if t == b {
			tossCSV(w, input)
		}
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(input)
		for i := 0; i < s.Len(); i++ {
			tossCSV(w, s.Index(i).Interface())
		}
	}
	w.Flush()
}

func createCSV(fileName string, fields []string) error {
	f, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }

	w := csv.NewWriter(f)
    err := w.Write(fields) 
    if err != nil {
        panic(err)
    }
	w.Flush()
	f.Close()
	return err
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
