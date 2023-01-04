package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	path := "actual"
	function := "addPrefix"
	files, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	} else {
		sortByTime(files)
		applyFunction(files, path, function)

	}
}

// Helper function
func changeModTime(files []fs.DirEntry) {
	tick := time.Tick(time.Millisecond)
	for _, file := range files {
		<-tick
		err := os.Chtimes("actual/"+file.Name(), time.Now(), time.Now())
		if err != nil {
			panic(err)
		}
	}
}

func applyFunction(files []fs.DirEntry, path string, function string) {
	for i, file := range files {
		oldPath := strings.Join([]string{path, file.Name()}, "/")
		newPath := getNewPath(file, path, i, function)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Successfully renamed files")
}

func getNewPath(file fs.DirEntry, path string, index int, function string) string {
	if function == "addPrefix" {
		newName := strconv.Itoa(index+1) + ". " + file.Name()
		return strings.Join([]string{path, newName}, "/")
	} else if function == "removePrefix" {
		newName := strings.SplitN(file.Name(), ". ", 2)[1]
		return strings.Join([]string{path, newName}, "/")
	} else {
		panic("Invalid function")
	}
}

func sortByTime(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		infoI, errI := files[i].Info()
		if errI != nil {
			panic(errI)
		}
		infoJ, errJ := files[j].Info()
		if errJ != nil {
			panic(errJ)
		}
		return infoI.ModTime().Before(infoJ.ModTime())
	})
}

// Helper function
func sortByName(files []fs.DirEntry) []fs.DirEntry {
	sortedFiles := make([]fs.DirEntry, len(files))
	copy(sortedFiles, files)
	sort.Slice(sortedFiles, func(i, j int) bool {
		iNum, _ := strconv.Atoi(strings.SplitN(sortedFiles[i].Name(), ".", 2)[0])
		jNum, _ := strconv.Atoi(strings.SplitN(sortedFiles[j].Name(), ".", 2)[0])
		return iNum < jNum
	})
	return sortedFiles
}
