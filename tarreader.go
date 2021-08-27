package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func readTar(srcFile string) {
	f, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			switch name {
			case "data/commits.csv":
				commitsFromCSV(tarReader)
			case "data/repos.csv":
				reposFromCSV(tarReader)
			case "data/actors.csv":
				actorsFromCSV(tarReader)
			case "data/events.csv":
				eventsFromCSV(tarReader)
			}
		default:
			log.Fatalf("Unexpected type: %c in file %s",
				header.Typeflag,
				name,
			)
		}
	}
}
