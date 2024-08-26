package bin

import (
    "reflect" 
    "os"
    "encoding/json"
)

type JSONBin struct {
	StructType reflect.Type
	Fields     []string
	FilePath   string
}


func (bin *JSONBin) Toss(input interface{}) {
	f, err := os.OpenFile(bin.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
