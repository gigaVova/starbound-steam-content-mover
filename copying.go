package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Check if the target folder exist and create one if not
func checkTargetDirectory(targetDir string) {
	if _, err := os.Stat(targetDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(targetDir, fs.ModeDir)
		if err != nil {
			slog.Error("")
		}
		fmt.Printf("Created %s\n", targetDir)
	}
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

func copyFiles(srcDirectoryPath, targetDirectoryPath string, wg *sync.WaitGroup) {
	wg.Add(1)
	// all entries (directories) of source folder
	entries, err := os.ReadDir(srcDirectoryPath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Copying.. ")
	startTime := time.Now()
	// going through all the directories of source folder
	for _, entry := range entries {
		if entry.IsDir() {
			// save the name (steam id)
			currentDirectoryName := entry.Name()
			// path to the current directory
			folderPath := filepath.Join(srcDirectoryPath, currentDirectoryName)

			// all entries (files) of directory with content
			files, err := os.ReadDir(folderPath)
			if err != nil {
				panic(err)
			}

			// going through all the entries
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
