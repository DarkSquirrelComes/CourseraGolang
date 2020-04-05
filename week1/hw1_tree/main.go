package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func prefix(isLast bool) (res string) {
	if !isLast {
		return "├───"
	}
	return "└───"
}

func globalPrefix(isLast bool) (res string) {
	if !isLast {
		return "│\t"
	}
	return "\t"
}

func fileSizeStr(file os.FileInfo) string {
	if file.Size() == 0 {
		return " (empty)"
	}
	return " (" + strconv.FormatInt(file.Size(), 10) + "b)"
}

func dirTreeRecursive(out io.Writer, path string, printFiles bool, oldPrefix string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})
	var isLast bool
	var lastDir int = -1
	for i, currentFile := range list {
		if currentFile.IsDir() {
			lastDir = i
		}
	}
	for i, currentFile := range list {
		if printFiles {
			isLast = (i == len(list)-1)
		} else {
			isLast = (i == lastDir)
		}
		if printFiles && !currentFile.IsDir() {
			fmt.Fprintln(out, oldPrefix+prefix(isLast)+currentFile.Name()+fileSizeStr(currentFile))
		}
		if currentFile.IsDir() {
			fmt.Fprintln(out, oldPrefix+prefix(isLast)+currentFile.Name())
			err = dirTreeRecursive(out, path+"/"+currentFile.Name(), printFiles, oldPrefix+globalPrefix(isLast))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTreeRecursive(out, path, printFiles, "")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
