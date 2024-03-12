package main

import (
	"flag"
	"github.com/adhocore/chin"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// spinner lol
	s := chin.New().WithWait(&wg)

	// directory we want to get content from
	var srcDirectoryPath = flag.String("src", "", "directory to move content from")
	// directory we want to move content to
	var targetDirectoryPath = flag.String("target", "./content", "directory to move content to")
	flag.Parse()

	checkTargetDirectory(*targetDirectoryPath)

	go s.Start()
	copyFiles(*srcDirectoryPath, *targetDirectoryPath, &wg)
	s.Stop()
	wg.Wait()
}
