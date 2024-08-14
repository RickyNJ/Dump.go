package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func GetHeaders[T any](items []T) []string {
	if items == nil {
		return []string{}
	}

	headers := []string{}

	val := reflect.ValueOf(items[0]).Elem()
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Name)
	}

	return headers
}

func CreateFile(fileName string) (*os.File, error){
	fileName = fileName + ".csv"
	f, err := os.Create(fileName)
	return f, err
}

func DumpToCSV[T any](fileName string, items []T) bool {

	f, err  := CreateFile(fileName)
    if err != nil {
        log.Fatalf("Error when creating the file %s", err)
    }

	headers := GetHeaders(items)

	w := csv.NewWriter(f)
	w.Write(headers)
	w.Flush()

	return true
}

func main() {
	fmt.Print("hi")
}
