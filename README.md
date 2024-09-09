 ____                                    
|  _ \ _   _ _ __ ___  _ __   __ _  ___  
| | | | | | | '_ ` _ \| '_ \ / _` |/ _ \ 
| |_| | |_| | | | | | | |_) | (_| | (_) |
|____/ \__,_|_| |_| |_| .__(_)__, |\___/ 
                      |_|    |___/       

# Dump.go

## Introcuction
Dump.go is a library written in pure Go and is the easiest way of writing struct instance into a file.
It is built on the idea of a "Bin", an object instantiated by linking it to a struct.

## Basic Usage

### Installation

### Initialize bin
When initializing a bin 
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
The Toss method supports arrays, slices and single instances of a struct.
Different struct types will result in a panic.
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
