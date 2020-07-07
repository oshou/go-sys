package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	MaxArgs = 2
)

func main() {
	aFlag := flag.Bool("a", false, "Include directory entries whose names begin with a dot (.).")
	tFlg := flag.Bool("t", false, "Sort by time modified (most recently modified first) before sorting the operands by lexicographical order.")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if *aFlag {
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
		}
		return
	}

	if *tFlag {
		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].ModTime().After(fileInfos[i].ModTime())
		})

		for _, fileInfo := range fileInfos {
			if strings.HasPrefix(fileInfo.Name(), ".") {
				continue
			}
			fmt.Println(fileInfo.Name())
		}
		return
	}

	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}
		fmt.Println(fileInfo.Name())
	}
}
