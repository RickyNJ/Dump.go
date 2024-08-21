package main

import (
    "testing"
)

func TestUnsupportedFileTypes (t *testing.T) {
    if getFileType("hai.db") != "unsupported" {
        t.Fatalf("not unsupported")
    }
}

func TestNoFileTypes (t *testing.T) {
    if getFileType("hai") != "please add filetype to the filename" {
        t.Fatalf("No File Type failed")
    }
}

func TestTooManyPeriodFileTypes (t *testing.T) {
    if getFileType("hai.lasd.slkdf.json") != "json" {
        t.Fatalf("Multiple periods failed")
    }
}
func TestCSVFileTypes (t *testing.T) {
    if getFileType("hai.csv") != "csv" {
        t.Fatalf("csv failed")
    }
}

func TestJSONFileTypes (t *testing.T) {
    if getFileType("hai.json") != "json" {
        t.Fatalf("json failed")
    }
}
