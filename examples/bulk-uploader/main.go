package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const baseURL = "http://api.gifs.com"

type gifResponse struct {
	Success struct {
		Page  string `json:"page"`
		Files struct {
			Gif  string `json:"gif"`
			Jpg  string `json:"jpg"`
			Mp4  string `json:"mp4"`
			Webm string `json:"webm"`
		} `json:"files"`
		Oembed string `json:"oembed"`
		Embed  string `json:"embed"`
		Meta   struct {
			Duration string `json:"duration"`
			Frames   string `json:"frames"`
			Height   string `json:"height"`
			Width    string `json:"width"`
		} `json:"meta"`
	} `json:"success"`
}

func Upload(url, file string) (gres gifResponse, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open(file)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("file", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}

	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &gres)
	if err != nil {
		panic(err.Error())
	}

	return
}

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
		res, err := Upload(baseURL+"/media/upload", file)
		if err != nil {
			fmt.Printf("\033[31m[ERROR]\033[0m %s\n", err.Error())
		} else {
			fmt.Printf("\033[32m[SUCCESS]\033[0m Successfully uploaded to %s\n", res.Success.Page)
		}
	}
}
