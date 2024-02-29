package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/adhocore/chin"
)

// Check if the target folder exist and create one if not
func checkTargetDirectory(targetDir string) {
	if _, err := os.Stat(targetDir); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(targetDir, fs.ModeDir)
		fmt.Printf("Created %s\n", targetDir)
	}
}

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

func copyFiles(srcDirectoryPath, targetDirectoryPath string, wg *sync.WaitGroup) {
	wg.Add(1)
	// all entries (directories) of source folder
	entries, err := os.ReadDir(srcDirectoryPath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Copying has started! ")
	startTime := time.Now()
	// going through all of the directories of source folder
	for _, entry := range entries {
		if entry.IsDir() {
			// save the name (steam id)
			currentDirectoryName := entry.Name()
			// path to the current directory
			folderPath := filepath.Join(srcDirectoryPath, currentDirectoryName)

			// all entries (files) of dirictory with content
			files, err := os.ReadDir(folderPath)
			if err != nil {
				panic(err)
			}

			// going through all of the entries
			for _, file := range files {
				if !file.IsDir() {
					if strings.Split(file.Name(), ".")[1] == "pak" {
						currentFileName := file.Name()
						currentFilePath := filepath.Join(folderPath, currentFileName)

						newFileName := currentDirectoryName + ".pak"
						newFilePath := filepath.Join(targetDirectoryPath, newFileName)

						err := copyFile(currentFilePath, newFilePath)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
	fmt.Printf("\nCopying took %s", time.Since(startTime))
	wg.Done()
}

func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
