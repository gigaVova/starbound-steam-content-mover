package main

import (
	"flag"
	"fmt"
	"github.com/adhocore/chin"
	"os"
	"sync"
)

type directories struct {
	srcDirectoryPath    *string
	targetDirectoryPath *string
}

func main() {
	var wg sync.WaitGroup
	var d directories
	// spinner lol
	s := chin.New().WithWait(&wg)

	// directory we want to get content from
	d.srcDirectoryPath = flag.String("src", "no default, set it by yourself", "directory to move content from")
	// directory we want to move content to
	d.targetDirectoryPath = flag.String("target", "./content", "directory to move content to")
	scrapTitles := flag.Bool("titles", false, "do scrap content titles or not")
	flag.Parse()

	if *d.srcDirectoryPath == "" {
		fmt.Println("src directory must be set")
		os.Exit(0)
	}

	checkTargetDirectory(*d.targetDirectoryPath)

	go s.Start()
	copyFiles(d, *scrapTitles, &wg)
	s.Stop()
	wg.Wait()
}
