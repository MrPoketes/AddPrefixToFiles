package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
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
