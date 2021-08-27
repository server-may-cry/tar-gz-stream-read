package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
)

func main() {
	log.SetFlags(log.Lshortfile)
	file := flag.String("path", "./data.tar.gz", "path to data.tar.gz file")
	flag.Parse()

	readTar(*file)

	actorCommits := make(map[int64]int64)
	repoCommits := make(map[int64]int64)
	for _, commit := range commits {
		event, ok := events[commit.EventId]
		if !ok {
			log.Fatalf("Event with id %d is not found", commit.EventId)
		}
		actorCommits[event.ActorId]++
		repoCommits[event.RepoID]++
	}

	actorPRs := make(map[int64]int64)
	repoWatchers := make(map[int64]int64)
	for _, event := range events {
		switch event.Type {
		case "WatchEvent":
			repoWatchers[event.RepoID]++
		case "PullRequestEvent":
			actorPRs[event.ActorId]++
		}
	}

	fmt.Println("Top 10 active users sorted by amount of PRs created and commits pushed")
	fmt.Println(names(top10actors(actorPRs, actorCommits), actors))
	fmt.Println("Top 10 repositories sorted by amount of commits pushed")
	fmt.Println(names(top10(repoCommits), repos))
	fmt.Println("Top 10 repositories sorted by amount of watch events")
	fmt.Println(names(top10(repoWatchers), repos))
}

func top10actors(prs, commits map[int64]int64) []int64 {
	type s struct {
		index, prs, commits int64
	}

	slice := make([]s, 0, len(prs))
	for i, v := range prs {
		slice = append(slice, s{i, v, commits[i]})
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].prs == slice[j].prs {
			return slice[i].commits > slice[j].commits
		}
		return slice[i].prs > slice[j].prs
	})

	r := make([]int64, 10)
	var i int
	for _, v := range slice {
		if i > 9 {
			break
		}
		r[i] = v.index
		i++
	}
	return r
}

func top10(m map[int64]int64) []int64 {
	type s struct {
		index, value int64
	}

	slice := make([]s, 0, len(m))
	for i, v := range m {
		slice = append(slice, s{i, v})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].value > slice[j].value
	})

	r := make([]int64, 10)
	var i int
	for _, v := range slice {
		if i > 9 {
			break
		}
		r[i] = v.index
		i++
	}
	return r
}

func names(ids []int64, m map[int64]string) []string {
	r := make([]string, 0, len(ids))
	for _, v := range ids {
		name, ok := m[v]
		if !ok {
			log.Fatalf("Can't find the name for ID %d", v)
		}
		r = append(r, name)
	}
	return r
}
