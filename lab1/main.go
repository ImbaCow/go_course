package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type pathConfig struct {
	Paths map[string]string `json:"paths"`
}

func getPathConfig(path string) (*pathConfig, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dat pathConfig
	if err := json.Unmarshal(buff, &dat); err != nil {
		return nil, err
	}
	return &dat, nil
}

func main() {
	filePath := flag.String("f", "", "config file path")
	flag.Parse()

	config, err := getPathConfig(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if val, ok := config.Paths[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusSeeOther)
		}
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8000", nil)
}
