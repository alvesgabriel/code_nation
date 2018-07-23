package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

type fileStruct struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {
	args := os.Args[1:]
	//@todo read dir name from stdin
	_ = jsonify(args[0])
}

func jsonify(dir string) error {
	f, err := createFileJSON()
	if err != nil {
		return err
	}
	f.Close()

	files := crateJSON(dir)
	sort.Slice(files, func(i, j int) bool {
		return path.Dir(files[i].Path) <= path.Dir(files[j].Path)
	})

	err = saveJSON(files)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func crateJSON(dir string) []fileStruct {
	var files []fileStruct

	filesDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	for _, f := range filesDir {
		name := f.Name()
		pathname := "./" + path.Join(dir, name)

		if f.IsDir() {
			files = append(files, crateJSON(pathname)...)
		} else {
			file := fileStruct{
				Name: name,
				Path: pathname,
			}
			files = append(files, file)
		}
	}
	return files
}

func saveJSON(data []fileStruct) error {
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
	filename := "files.json"
	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("Not opened %+v. %+v", filename, err)
	}

	return f, nil
}
