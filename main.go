package main

import (
	"flag"
	"github.com/adhocore/chin"
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
	d.srcDirectoryPath = flag.String("src", "", "directory to move content from")
	// directory we want to move content to
	d.targetDirectoryPath = flag.String("target", "./content", "directory to move content to")
	flag.Parse()

	checkTargetDirectory(*d.targetDirectoryPath)

	go s.Start()
	copyFiles(d, &wg)
	s.Stop()
	wg.Wait()
}
