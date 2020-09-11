package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type comic struct {
	Image      url.URL  `json:"img,string"`
	Transcript []string `json:"transcript"`
}

func getComic(index int, c *comic) error {
	resp, err := http.Get("https://xkcd.com/" + string(rune(index)) + "/info.0.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return err
	}
	log.Println(resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		resp.Body.Close()
		return err
	}
	return nil
}

func main() {
	var c comic
	err := getComic(571, &c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c.Image)
	// for _, l := range c.Transcript {
	// 	fmt.Println(l)
	// }
}
