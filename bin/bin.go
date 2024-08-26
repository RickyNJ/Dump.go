package bin

import (
	"encoding/csv"
	"encoding/json"

	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// total funcs for now
// NewBin, LoadBin, Toss

type Bin interface {
	Toss(input interface{})
}

type CSVBin struct {
	structType reflect.Type
	fields     []string
	filePath   string
}

type JSONBin struct {
	structType reflect.Type
	fields     []string
	filePath   string
}

func (bin *JSONBin) Toss(input interface{}) {
	f, err := os.OpenFile(bin.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Couldnt open file")
	}
	f.Close()
}

func tossCSV(bin *CSVBin, w *csv.Writer, input reflect.Value) {
	newLine := []string{}
	for _, v := range bin.fields {
		value := input.FieldByName(v)
		newValue := []string{}
		if value.Kind() == reflect.Struct {

		} else {
			newValue = append(newValue, fmt.Sprint(value))
		}
		newLine = append(newLine, newValue...)
	}
	w.Write(newLine)
}

func (bin *CSVBin) Toss(input interface{}) {
	f, err := os.OpenFile(bin.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Couldnt open file")
	}

	w := csv.NewWriter(f)
	t := reflect.TypeOf(input)
	b := bin.structType

	switch reflect.ValueOf(input).Kind() {
	case reflect.Struct:
		if t == b {
			tossCSV(bin, w, reflect.ValueOf(input))
		}
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(input)
		for i := 0; i < s.Len(); i++ {
			tossCSV(bin, w, s.Index(i))
		}
	}
	w.Flush()
}

func createCSV(fileName string, fields []string) error {
	f, err := os.Create(fileName)
	w := csv.NewWriter(f)
	w.Write(fields)
	w.Flush()
	f.Close()
	return err
}

func createJSON(fileName string, structname string) error {
	data := map[string][]interface{}{
		structname: {},
	}
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic("failed to marshall")
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(jsonData)
	f.Close()
	return err
}

func getFileType(filename string) string {
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

	var recursiveSearch func(parentStruct string, input interface{}) []string

	recursiveSearch = func(parentStruct string, input interface{}) []string {

		res := []string{}
		structType := reflect.TypeOf(input)
		structValue := reflect.ValueOf(input)

		for i := 0; i < structType.NumField(); i++ {
			field := structType.Field(i)
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
	fmt.Printf("Generating new Bin with filename: %v and input struct %v \n", fileName, structType)
	if structType.Kind() != reflect.Struct {
		panic("input is not a struct")
	}

	fileType := getFileType(fileName)
	fields := getStructFieldNames(inputStruct)

	fmt.Printf("extracted file type: %v, fields: %v \n", fileType, fields)
	switch fileType {
	case "csv":
		err := createCSV(fileName, fields)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created file %v \n", fileName)
		return &CSVBin{
			structType: structType,
			fields:     fields,
			filePath:   fileName,
		}

	case "json":
		err := createJSON(fileName, structType.Name())
		if err != nil {
			log.Fatal(err)
		}
		return &JSONBin{
			structType: structType,
			fields:     fields,
			filePath:   fileName,
		}
	}

	return nil
}
