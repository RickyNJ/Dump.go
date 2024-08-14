package main

import (
 "fmt"
 "reflect"
) 




func GetHeaders [T any](items []T) []string {
    if items == nil {
        return []string{}
    }

    headers := []string{}

    val := reflect.ValueOf(items[0]).Elem()
    for i:=0; i<val.NumField();i++{
        headers = append(headers,val.Type().Field(i).Name )
    }

    return headers
}

func main() {
    fmt.Print("hi")
}


