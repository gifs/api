package bulk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type gifResponse struct {
	Page  string `json:"page"`
	Files struct {
		Gif  string `json:"gif"`
		Jpg  string `json:"jpg"`
		Mp4  string `json:"mp4"`
		Webm string `json:"webm"`
	} `json:"files"`
	Embed string `json:"embed"`
	Meta  struct {
		Duration string `json:"duration"`
		Frames   string `json:"frames"`
		Height   string `json:"height"`
		Type     string `json:"type"`
		Width    string `json:"width"`
	} `json:"meta"`
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
