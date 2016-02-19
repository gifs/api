package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gifs/bulk-uploader"
)

const baseURL = "http://api.gifs.com"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gifs-bulk-uploader ./path/to/directory")
		os.Exit(1)
	}

	searchDir := os.Args[1]
	fileList := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		ext := filepath.Ext(f.Name())
		if ext == ".gif" || ext == ".mp4" || ext == ".webm" {
			fileList = append(fileList, path)
		}
		return nil
	})

	for _, file := range fileList {
		fmt.Print("Uploading " + file + " to gifs API... ")
		res, err := bulk.Upload(baseURL+"/media/upload", file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("successfully added media to " + res.Page)
	}

}
