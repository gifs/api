package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Post holds the data from reddit
type Post struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Over18 bool   `json:"over_18"`
	Media  struct {
		Oembed struct {
			Description string `json:"description"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"oembed"`
	} `json:"media"`
}

// Get subreddit(s) media (Example: "/r/gifs+woahdude")
func Get(subreddits string, offset string) ([]Post, error) {
	rurl := fmt.Sprintf("http://reddit.com/%s.json?count=25", subreddits)

	resp, err := http.Get(rurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	type redditResponse struct {
		Data struct {
			Children []struct {
				Data Post
			}
		}
	}

	r := new(redditResponse)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	posts := make([]Post, len(r.Data.Children))
	for i, child := range r.Data.Children {
		posts[i] = child.Data
	}
	return posts, nil
}

func main() {
	if len(os.Args) < 2 {
		print("Usage:\nirg /r/reactiongifs+politicalreactiongifs\n")
		os.Exit(1)
	}
	posts, err := Get(os.Args[1], "")
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range posts {
		if p.Media.Oembed.Type == "video" {

			type attribution struct {
				Site string `json:"site"`
				User string `json:"user"`
			}

			type importRequest struct {
				Source      string      `json:"source"`
				Title       string      `json:"title"`
				Tags        []string    `json:"tags"`
				NSFW        bool        `json:"NSFW"`
				Attribution attribution `json:"attribution"`
			}

			payload, _ := json.Marshal(importRequest{
				Source: p.URL,
				Title:  p.Title,
				NSFW:   p.Over18,
				Attribution: attribution{
					Site: "reddit",
					User: p.Author,
				},
			})

			baseURL := "https://api.gifs.com/media/import"
			req, _ := http.NewRequest("POST", baseURL, bytes.NewReader(payload))
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			type importResponse struct {
				Success struct {
					Page string `json:"page"`
				} `json:"success"`
			}

			ires := new(importResponse)
			defer res.Body.Close()
			err = json.NewDecoder(res.Body).Decode(ires)
			if err != nil {
				log.Fatal(err)
			}

			titleCopy := p.Title

			if p.Over18 {
				titleCopy = "[NSFW] " + titleCopy
			}

			titlePrintLen := 40
			titleLen := len(p.Title)
			truncateLen := titlePrintLen
			if titleLen < truncateLen {
				truncateLen = titleLen
			}

			for i := 0; i < titlePrintLen-titleLen; i++ {
				titleCopy = titleCopy + "."
			}

			status := ""
			if res.StatusCode == 200 {
				status = fmt.Sprintf("\033[32m%d\033[0m", res.StatusCode)
			} else {
				status = fmt.Sprintf("\033[33m%s\033[0m", res.Status)
			}

			fmt.Printf("%s... %s %s\n",
				titleCopy[0:titlePrintLen], status, ires.Success.Page)

		}
	}
}
