(Hopefully) The easiest way to have persistent data in golang.

Initialize a bin using bin.NewBin(FILENAME, YOURSTRUCT{})

The goal is to support json, csv and excel files.
You can now use the bin.Toss() method to write single or struct slices/arrays to this bin

