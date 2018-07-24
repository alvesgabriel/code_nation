package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type site struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CreatedBy     string   `json:"created_by"`
	ReleasedAt    string   `json:"released_at"`
	Repositories  int      `json:"repositories"`
	RelatedTopics []string `json:"related_topics"`
}

func main() {
	_ = parseHTML("golang.html")
}

func parseHTML(page string) error {
	html, err := getHTML(page)
	if err != nil {
		return err
	}
	s, err := createSite(html)
	if err != nil {
		return err
	}
	if err := saveJSON(*s); err != nil {
		return err
	}
	return nil
	// return fmt.Errorf("Not implemented")
}

func getHTML(page string) (io.Reader, error) {
	file, err := ioutil.ReadFile("golang.html")
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(file)), err
}

func createSite(html io.Reader) (*site, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	name := doc.Find("h1").Text()
	description := doc.Find("body > div.application-main > div.container-lg.topic.p-responsive > div.py-6 > div > div > p").Text()
	released := doc.Find("body > div.application-main > div.container-lg.topic.p-responsive > div.py-6 > div > ul > li:nth-child(3)").Text()
	created := doc.Find("body > div.application-main > div.container-lg.topic.p-responsive > div.py-6 > div > ul > li:nth-child(2)").Text()
	created = strings.Split(created, "Created by ")[1]

	repositories := doc.Find("h2 > span").Text()
	repositories = strings.Join(strings.Split(repositories, ","), "")
	total, _ := strconv.Atoi(repositories)

	var topics []string
	doc.Find("body > div.application-main > div.container-lg.topic.p-responsive > div.d-md-flex.gutter-md > div.col-md-4.mt-6.mt-md-0 > a").
		Each(func(i int, s *goquery.Selection) {
			re := regexp.MustCompile(`\w+`)
			body := []byte(s.Text())
			repository := re.Find(body)
			topics = append(topics, string(repository))
		})

	s := site{
		Name:        name,
		Description: description,
		CreatedBy: strings.Trim(created, ` 
		`),
		ReleasedAt: strings.Trim(strings.Split(released, "Released ")[1], `
        `),
		Repositories:  total,
		RelatedTopics: topics,
	}

	return &s, nil
}

func saveJSON(data site) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Not created json. %+v", err)
	}

	var jsonFormat bytes.Buffer
	json.Indent(&jsonFormat, jsonData, "", "    ")

	f, err := createFileJSON()
	if err != nil {
		return err
	}
	defer f.Close()
	jsonFormat.WriteTo(f)

	return nil
}

func createFileJSON() (*os.File, error) {
	filename := "golang.json"
	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("Not opened %+v. %+v", filename, err)
	}

	return f, nil
}
