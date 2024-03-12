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

var (
	newFileName string
	// If the folder contains more than one .pak file (e.g. Stardust Core has two .pak files
	// in its folder), a prefix will be added to the result files
	prefix int
)

// Check if the target folder exist and create one if not
func checkTargetDirectory(targetDir string) {
	if _, err := os.Stat(targetDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(targetDir, fs.ModeDir)
		if err != nil {
			slog.Error("couldn't create a directory", "error", err, "targetDir", targetDir)
		}

		slog.Info("successfully created", "targetDir", targetDir)
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

func copyFiles(d directories, scrapTitles bool, wg *sync.WaitGroup) {
	wg.Add(1)
	// all entries (directories) of source folder
	entries, err := os.ReadDir(*d.srcDirectoryPath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Copying.. ")
	startTime := time.Now()

	if scrapTitles {
		// going through all the directories of source folder
		for _, entry := range entries {
			if entry.IsDir() {
				// reset the prefix
				prefix = 1
				// save the name (Steam id)
				currentDirectoryName := entry.Name()
				// path to the current directory
				folderPath := filepath.Join(*d.srcDirectoryPath, currentDirectoryName)

				// all entries (files) of directory with content (.pak files expectedly)
				files, err := os.ReadDir(folderPath)
				if err != nil {
					panic(err)
				}

				// going through all the entries
				for _, file := range files {
					if !file.IsDir() {
						// check if the current file is ".pak"
						if filepath.Ext(file.Name()) == ".pak" {
							// get the path of .pak file from Starbound's workshop folder
							currentFileName := file.Name()
							currentFilePath := filepath.Join(folderPath, currentFileName)

							// set new file name's to a title in the Steam workshop
							newFileName = getWorkshopItemTitle(currentDirectoryName) + ".pak"
							if prefix > 1 {
								newFileName = getWorkshopItemTitle(currentDirectoryName) + "_" + fmt.Sprint(prefix) + ".pak"
							}
							prefix++

							// path to a new file with its name and target folder
							newFilePath := filepath.Join(*d.targetDirectoryPath, newFileName)

							// skip if the file already exists
							if _, err := os.Stat(newFilePath); errors.Is(err, os.ErrNotExist) {
								err := copyFile(currentFilePath, newFilePath)
								if err != nil {
									panic(err)
								}
							}
						}
					}
				}
			}
		}
	} else {
		// going through all the directories of source folder
		for _, entry := range entries {
			if entry.IsDir() {
				// reset the prefix
				prefix = 1
				// save the name (Steam id)
				currentDirectoryName := entry.Name()
				// path to the current directory
				folderPath := filepath.Join(*d.srcDirectoryPath, currentDirectoryName)

				// all entries (files) of directory with content (.pak files expectedly)
				files, err := os.ReadDir(folderPath)
				if err != nil {
					panic(err)
				}

				// going through all the entries
				for _, file := range files {
					if !file.IsDir() {
						// check if the current file is ".pak"
						if strings.Split(file.Name(), ".")[1] == "pak" {
							// get the path of .pak file from Starbound's workshop folder
							currentFileName := file.Name()
							currentFilePath := filepath.Join(folderPath, currentFileName)

							// set new file name's to an ID in the Steam workshop
							newFileName = currentDirectoryName + ".pak"
							if prefix > 1 {
								newFileName = currentDirectoryName + "_" + fmt.Sprint(prefix) + ".pak"
							}
							prefix++

							// path to a new file with its name and target folder
							newFilePath := filepath.Join(*d.targetDirectoryPath, newFileName)

							// skip if the file already exists
							if _, err := os.Stat(newFilePath); errors.Is(err, os.ErrNotExist) {
								err := copyFile(currentFilePath, newFilePath)
								if err != nil {
									panic(err)
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("\nCopying took %s", time.Since(startTime))
	wg.Done()
}
