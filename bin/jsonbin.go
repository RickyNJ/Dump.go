package bin

import (
    "reflect" 
    "os"
    "encoding/json"
)

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
