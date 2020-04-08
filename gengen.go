// +build ignore

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// tool to generate the go generate commands

var excludeAlgorithms = map[string]struct{}{
	"german2":         struct{}{},
	"kraaij_pohlmann": struct{}{},
	"lovins":          struct{}{},
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("must specify algorithms directory")
	}

	files, err := ioutil.ReadDir(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		languange := strings.TrimSuffix(file.Name(), ".sbl")
		if _, ok := excludeAlgorithms[languange]; ok {
			continue
		}
		fmt.Printf("//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/%s -go -o %s/%s_stemmer -gop %s -gor github.com/blevesearch/snowballstem\n",
			file.Name(), languange, languange, languange)
		fmt.Printf("//go:generate gofmt -s -w %s/%s_stemmer.go\n",
			languange, languange)
		fmt.Println()
	}
}
