package main

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
)

type Event struct {
	Type    string `csv:"type"`
	Id      int64  `csv:"id"`
	ActorId int64  `csv:"actor_id"`
	RepoID  int64  `csv:"repo_id"`
}

var events = make(map[int64]Event)

func eventsFromCSV(file io.Reader) {
	csvReader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var event Event
	for {
		if err := dec.Decode(&event); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch event.Type {
		case "PushEvent", "PullRequestEvent", "WatchEvent":
			events[event.Id] = event
		}
	}
}
