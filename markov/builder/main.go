package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

//A name is the name of an Ngram
type Name string

//A gram
type NGram struct {
	Name
	Length, Count int64
}

//An NGDB is a map of NGrams indexed by their name
type NGDB map[Name]NGram

var file = flag.String("file", "text.txt", "The path to the file from which you will build your markov database.")

func main() {
	flag.Parse()

	f, err := ioutil.ReadFile(*file)
	if err != nil {
		errors.New(err.Error())
	}

	db, lenF := NGDB{}, len(f)

	for pos, _ := range f {
		if (pos + 1) >= lenF {
			break
		}
		
		n := (string(f[pos]) + string(f[pos+1]))
		name := Name(n)
		
		//Add this NGram to our database (or increment the count of this NGram if it already exists)
		db[name] = NGram{Name:name, Length:int64(len(name)), Count: db[name].Count+1}
	}
	
	fmt.Printf("%+v\n\n", db)
}
