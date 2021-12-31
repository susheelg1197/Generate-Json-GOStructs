package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//go:generate go run gen.go
//go:generate go fmt codegen/user.gen.go

//OS Read Directory
func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, strings.ReplaceAll(file.Name(), ".json", ""))
	}
	return files, nil
}

func main() {
	root := "../apidata/userList"
	files, err := OSReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		var resp map[string]interface{}
		in, _ := os.Open(fmt.Sprintf("../apidata/userList/%v.json", file))
		b, _ := ioutil.ReadAll(in)
		json.Unmarshal(b, &resp)

		data := struct {
			Name   string
			Fields map[string]interface{}
		}{
			strings.ReplaceAll(file, " ", ""),
			resp,
		}

		tpl, _ := template.New("template.tpl").Funcs(template.FuncMap{
			"Title": strings.Title,
			"TypeOf": func(v interface{}) string {
				if v == nil {
					return "string"
				}
				return strings.ToLower(reflect.TypeOf(v).String())
			},
		}).ParseFiles("template.tpl")

		out, _ := os.Create(fmt.Sprintf("codegen/generated_%v.go", file))
		defer out.Close()

		tpl.Execute(out, data)
	}
}
