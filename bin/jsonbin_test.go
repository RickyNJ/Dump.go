package bin

import (
	"fmt"
	"testing"
)


func TestJSONCreation(t *testing.T){
    err := createJSON("test.json", "Person")
    fmt.Print(err)
}
