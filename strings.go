package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

// A drawable directory under res
type ValuesDirectory struct {
	Name      string
	Path      string
	StringFiles []*StringFile
}

func NewValuesDirectory(path string) (*ValuesDirectory, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Get all string files in directory
	stringFiles := []*StringFile{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, stringInfo := range files {
		stringFile, err := NewStringFile(filepath.Join(path, stringInfo.Name()))
		// Ignore files which produce errors
		if err == nil && stringFile != nil {
			stringFiles = append(stringFiles, stringFile)
		}
	}

	return &ValuesDirectory{
		fileInfo.Name(),
		path,
		stringFiles,
	}, nil
}

// A string file
type StringFile struct {
	Name string
	Path string
}

// Create Drawable from file at path
func NewStringFile(path string) (*StringFile, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(fileInfo.Name(), "strings")	{
		return nil, nil
	}

	fmt.Println("String Files", fileInfo.Name())

	return &StringFile{
		fileInfo.Name(),
		path,
	}, nil
}