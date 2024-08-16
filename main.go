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

func CreateFile(fileName string, headers []string) (*os.File, error){
	fileName = fileName + ".csv"
	f, err := os.Create(fileName)	
    w := csv.NewWriter(f)
	w.Write(headers)
	w.Flush()

    return f, err
}

func WriteFile[T any](f *os.File, items []T ) (*os.File, error){
    
}



func DumpToCSV[T any](fileName string, items []T) bool {
	headers := GetHeaders(items)

	f, err  := CreateFile(fileName, headers)
    if err != nil {
        log.Fatalf("Error when creating the file %s", err)
    }

    f.Close()


	return true
}

func main() {
	fmt.Print("hi")
}
