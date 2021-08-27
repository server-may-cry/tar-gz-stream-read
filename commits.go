package main

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
)

type Commit struct {
	EventId int64 `csv:"event_id"`
}

var commits []Commit

func commitsFromCSV(file io.Reader) {
	csvReader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var commit Commit
	for {
		if err := dec.Decode(&commit); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		commits = append(commits, commit)
	}
}
