package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

/*

JAVASCRIPT TO RUN WHEN ONLINE

// {
//   "french": "that.childNodes[3].childNodes[5].childNodes[5].dataset.title",
//   "english": "that.childNodes[3].childNodes[5].childNodes[5].dataset.translation",
//   "image": "that.childNodes[1].childNodes[1].src",
//   "audio": "that.dataset.audioPath"
// }

var arr = []
var entries = $('.js-group-entry')

for(i=0; i<entries.length; i++) {
	var that = entries[i];
	var e = {
		"french": that.childNodes[3].childNodes[5].childNodes[5].dataset.title,
		"english": that.childNodes[3].childNodes[5].childNodes[5].dataset.translation,
		"image": that.childNodes[1].childNodes[1].src,
		"audio": that.dataset.audioPath,
		"base": that.childNodes[3].childNodes[5].childNodes[5].dataset.title.match(/[a-zA-ZÀ-ÿ]+/g).join('-').toLowerCase()
	}
	arr.push(e)
}

copy JSON array data to a file called data.json

*/

type FrenchCard struct {
	French  string `json:"french"`
	English string `json:"english"`
	Image   string `json:"image"`
	Audio   string `json:"audio"`
	Base    string `json:"base"`
}

func main() {
	// get JSON data from file the lazy way
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal into struct array
	cards := make([]FrenchCard, 0)
	if err := json.Unmarshal(data, &cards); err != nil {
		log.Fatal(err)
	}

	// create channels for workers
	jobs := make(chan FrenchCard, 5)
	done := make(chan FrenchCard, len(cards))

	// spin up workers
	for i := 0; i < 5; i++ {
		go worker(jobs, done)
	}

	// assign jobs to workers
	for _, card := range cards {
		jobs <- card
	}
	close(jobs)

	// wait for everything to be finished
	for i := 0; i < len(cards); i++ {
		card := <-done
		fmt.Printf("French:\t\t%s\nEnglish:\t%s\n\n", card.French, card.English)
	}
}

// worker takes in a FrenchCard containing the assets to be
// downloaded for each vocabulary word. will pass the data off
// to the completed worker after it finishes, error or not
func worker(c chan FrenchCard, done chan FrenchCard) {
	for card := range c {
		g, _ := errgroup.WithContext(context.Background())
		g.Go(downloader(card.Base, card.Image))
		g.Go(downloader(card.Base, card.Audio))

		if err := g.Wait(); err != nil {
			log.Printf("unable to download: %v\n", card)
		}
		done <- card
	}
}

// downloader takes in the base filename from the JSON feed
// and downloads the URL of the asset. Either an MP3 or JPG
// file and returns a function that can be used with errgroup
func downloader(base, u string) func() error {
	return func() error {
		ext := filepath.Ext(u)

		f, err := os.Create(fmt.Sprintf("%s%s", base, ext))
		if err != nil {
			return err
		}
		defer f.Close()

		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(f, resp.Body)

		return nil
	}
}
