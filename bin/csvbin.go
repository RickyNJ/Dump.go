package bin

import (
    "reflect"
    "os"
    "fmt"
    "encoding/csv"
)

type CSVBin struct {
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
