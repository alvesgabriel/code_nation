package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	var files []fileStruct
	dirname := path.Dir(dir)
	filesDir, err := ioutil.ReadDir(dirname)
	if err != nil {
		return fmt.Errorf("Not a dir. %+v", err)
	}
	for _, f := range filesDir {
		name := f.Name()
		pathname := fmt.Sprintf("%v/%v", dirname, name)
		// fmt.Println(name, pathname)

		// fi, err := os.Stat(name)
		// if err != nil {
		// 	return fmt.Errorf("Not a file. %+v", err)
		// }
		// switch mode := fi.Mode(); {
		// case mode.IsDir():
		// 	jsonify(pathname)
		// case mode.IsRegular():
		// }
		file := fileStruct{
			Name: name,
			Path: pathname,
		}
		files = append(files, file)

		// fmt.Printf("%+v\n", files)
	}
	err = saveJson(files)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func saveJson(data []fileStruct) error {
	filename := "files.json"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Not created json. %+v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Not opened %+v. %+v", filename, err)
	}
	defer f.Close()
	f.Write([]byte(jsonData))

	return nil
}
