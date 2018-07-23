package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Stars       int    `json:"stars"`
}

func main() {
	_ = githubStars("go")
}

func githubStars(lang string) error {
	body := githubRepositories(lang)
	repositories := stringToMap(body)
	repos := starsGolang(repositories)
	err := saveJSON(repos)
	if err != nil {
		return err
	}

	return nil
	// return fmt.Errorf("Not implemented")
}

func githubRepositories(lang string) []byte {
	url := fmt.Sprintf("https://api.github.com/search/repositories?"+
		"q=language:%s&sort=stars&order=desc&per_page=10", lang)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Can't open API Github\n%+v\n", err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}

func stringToMap(body []byte) map[string]interface{} {
	var repositories = make(map[string]interface{})
	err := json.Unmarshal(body, &repositories)
	if err != nil {
		return nil
	}
	return repositories
}

func starsGolang(repositories map[string]interface{}) []Repository {
	var repos []Repository
	for _, repository := range repositories["items"].([]interface{}) {
		r := Repository{
			Name:        repository.(map[string]interface{})["full_name"].(string),
			Description: repository.(map[string]interface{})["description"].(string),
			URL:         repository.(map[string]interface{})["html_url"].(string),
			Stars:       int(repository.(map[string]interface{})["stargazers_count"].(float64)),
		}
		repos = append(repos, r)
	}
	return repos
}

func saveJSON(data []Repository) error {
	filename := "stars.json"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Not created json. %+v", err)
	}

	var jsonFormat bytes.Buffer
	json.Indent(&jsonFormat, jsonData, "", "    ")

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Not opened %+v. %+v", filename, err)
	}
	defer f.Close()
	jsonFormat.WriteTo(f)

	return nil
}
