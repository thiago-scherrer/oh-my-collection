package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	usr := os.Getenv("USER")
	token := os.Getenv("TOKEN")
	url := "https://api.discogs.com/users/" + usr + "/collection"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Discogs token="+token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resBody)
}
