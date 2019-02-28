package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type school struct {
	URL  string `json:"url"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

func main() {
	f, err := os.Open("./new-edu.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var schools []school
	json.Unmarshal(b, &schools)
	fmt.Printf("本库共收录%d所高校\n", len(schools))
}
