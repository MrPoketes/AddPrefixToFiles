package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var selectedPrefixOption string
	var path string

	a := app.New()
	window := a.NewWindow("AddPrefixToFiles")
	window.Resize(fyne.NewSize(400, 400))

	// BaseForm
	form := widget.NewForm()
	form.OnSubmit = func() {
		fmt.Println("submit")
		fmt.Println(selectedPrefixOption)
		fmt.Println(path)
		// TODO: Validate and then do bellow actions

		// files, err := os.ReadDir(path)

		// if err != nil {
		// 	panic(err)
		// } else {
		// 	sortByTime(files)
		// 	applyFunction(files, path, function)
		// }
	}

	// SelectField
	selectField := widget.NewSelect([]string{"Add Prefix", "Remove Prefix"}, func(option string) {
		selectedPrefixOption = option
	})
	selectField.SetSelectedIndex(0)

	// Path button and path folder selector
	// As the button and the folder selector both call each other,
	// we defined tapped function seperately, after creating pathSelector
	pathButton := widget.NewButton("Path", func() {})
	pathSelector := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil {
			// TODO: Find a better way to handle this
			panic(err)
		}
		path = lu.Path()
		pathButton.Text = path
		form.Refresh()
	}, window)

	pathButton.OnTapped = func() { pathSelector.Show() }

	form.AppendItem(widget.NewFormItem("Select function", selectField))
	form.AppendItem(widget.NewFormItem("Select folder", pathButton))

	// form.AppendItem(functionSelect)
	window.SetContent(
		container.NewVBox(
			form,
		))

	window.ShowAndRun()

	fmt.Println("hello")
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
