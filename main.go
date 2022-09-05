package main

import (
	"log"
	"os/exec"
)

func main() {
	user := "jorgepvasconcelos" // your target user must be replaced here
	publicReposUrl := GetPublicRepos(user)
	for _, value := range publicReposUrl {
		args := []string{"clone", value.Html_url}

		if err := exec.Command("git", args...).Run(); err != nil {
			log.Fatal(err)
		}
	}
}
