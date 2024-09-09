<pre>
 ____                                    
|  _ \ _   _ _ __ ___  _ __   __ _  ___  
| | | | | | | '_ ` _ \| '_ \ / _` |/ _ \ 
| |_| | |_| | | | | | | |_) | (_| | (_) |
|____/ \__,_|_| |_| |_| .__(_)__, |\___/ 
                      |_|    |___/       
</pre>
# Dump.go

## Introduction
Dump.go is a library written in pure Go and is the easiest way of writing struct instance into a file. <br>
It is built on the idea of a "Bin", an object instantiated by linking it to a struct. The bin object has a single Toss method which can be used to write single or multiple instances of the struct into a file. Currently dump.go supports the following file types:
- csv
- Microsoft Excel (xlsx, xlam, xlsm, xltm, xltx)
- json
  
## Basic Usage

### Installation

### Initialize bin
Initialize a bin with the NewBin(filename string, struct Struct) function. Using an unsupported filetype will result in an error.
> [!WARNING]
> Data in the struct when calling NewBin will not be written into the file, instead call NewBin using an empty struct and Toss the struct instance afterwards.
```
type Person struct {
  Name string
  Age int
}

b := bin.NewBin("people.csv", Person{})
```
#### Generated file: 
people.csv

| Name | Age |
| ------------- | ------------- |

### Toss struct instances to a bin
> [!NOTE]
> The Toss method supports arrays, slices and single instances of a struct. Different struct types will result in a panic.
```
person1 := Person{Name: "jake", Age: 12}
person2 := Person{Name: "jakey", Age: 24}

b.Toss(person1)
b.Toss([]Person{person1, person2})
```
#### Generated file
people.csv
| Name | Age |
| ------------- | ------------- |
| jake  | 12  |
| jake  | 12  |
| jakey  | 24  |

### Load bin
```
b := bin.LoadBin("people.csv")
```
