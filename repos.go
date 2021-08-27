package main

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
)

type Repo struct {
	Id   int64  `csv:"id"`
	Name string `csv:"name"`
}

var repos = make(map[int64]string)

func reposFromCSV(file io.Reader) {
	csvReader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var repo Repo
	for {
		if err := dec.Decode(&repo); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		repos[repo.Id] = repo.Name
	}
}
