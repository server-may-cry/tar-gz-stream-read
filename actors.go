package main

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
)

type Actor struct {
	Id       int64  `csv:"id"`
	Username string `csv:"username"`
}

var actors = make(map[int64]string)

func actorsFromCSV(file io.Reader) {
	csvReader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var user Actor
	for {
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		actors[user.Id] = user.Username
	}
}
