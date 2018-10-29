package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type xkcd struct {
	Month      string `json:"month"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	var arrJSON []xkcd
	arg, _ := strconv.ParseInt(os.Args[1], 10, 64)
	a := int(arg)
	for i := 1; i <= a; i++ {
		rJSON := new(xkcd)
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		fmt.Printf("Getting URL : %s\n", url)
		err := getJSON(url, rJSON)
		if err != nil {
			log.Fatal(err)
		}
		arrJSON = append(arrJSON, *rJSON)
	}
	for i := 0; i < a; i++ {
		getImg(arrJSON[i].Img, arrJSON[i].Title)
	}
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getImg(url string, fileName string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	file, err := os.Create("./" + fileName + ".png")
	_, err = io.Copy(file, r.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
