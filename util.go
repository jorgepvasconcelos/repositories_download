package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RepositoriesStruct struct {
	Html_url string `json:"html_url"`
}

func GetPublicRepos(user string) []RepositoriesStruct {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var repositoriesList []RepositoriesStruct
	if err := json.Unmarshal(body, &repositoriesList); err != nil {
		log.Fatal(err)
	}

	return repositoriesList
}
