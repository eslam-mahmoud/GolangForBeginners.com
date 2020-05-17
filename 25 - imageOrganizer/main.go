package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sync"
	"time"
)

/*
	take input source folder
	take input destination folder
	get all files
	match each file
	extract info from file name
	create destination folder
	move file
*/
type fileToBeMoved struct {
	sourcePath      string
	fileName        string
	destinationPath string
}

func main() {
	var sourcePath string
	var sourceFiles []os.FileInfo
	var destinationPath string
	var err error

	for {
		fmt.Println("Please enter source path")
		fmt.Scanln(&sourcePath)
		sourceFiles, err = ioutil.ReadDir(sourcePath)
		if err != nil {
			log.Fatal(err)
		} else if len(sourceFiles) == 0 {
			log.Println("Empty source folder")
		} else {
			break
		}
	}

	for {
		fmt.Println("Please enter destination path")
		fmt.Scanln(&destinationPath)
		_, err := ioutil.ReadDir(destinationPath)
		if err == nil {
			break
		}
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	filesChan := make(chan fileToBeMoved)
	for i := 0; i < 1; i++ {
		fmt.Println("add worker", i)
		wg.Add(1)
		go move(&wg, filesChan)
	}
	start := time.Now()
	for i := 0; i < len(sourceFiles); i++ {
		re := regexp.MustCompile(`([0-9]{4})([0-9]{2})([0-9]{2})`)
		matches := re.FindStringSubmatch(sourceFiles[i].Name())
		if len(matches) == 0 {
			continue
		}
		filesChan <- fileToBeMoved{
			sourcePath:      sourcePath,
			fileName:        sourceFiles[i].Name(),
			destinationPath: path.Join(destinationPath, matches[1], matches[2], matches[3]),
		}
	}
	close(filesChan)
	fmt.Printf("took %v\n", time.Since(start))
	wg.Wait()
	fmt.Println("main function will exit now")
}
func move(wg *sync.WaitGroup, filesChan chan fileToBeMoved) {
	defer wg.Done()
	for f := range filesChan {
		// make sure folder exist or create it
		_, err := os.Stat(f.destinationPath)
		if os.IsNotExist(err) {
			os.MkdirAll(f.destinationPath, os.ModePerm)
		}
		// move the file
		err = os.Rename(path.Join(f.sourcePath, f.fileName), path.Join(f.destinationPath, f.fileName))
		if err != nil {
			log.Println(err)
		}
		// log result
		// fmt.Println("moved", path.Join(f.sourcePath, f.fileName), "to", path.Join(f.destinationPath, f.fileName))
	}
}
