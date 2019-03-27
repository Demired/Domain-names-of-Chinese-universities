package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type item struct {
	Domain string `json:"domain"`
	Tag    string `json:"tag"`
	Name   string `json:"name"`
}

var (
	tPath   = kingpin.Flag("tPath", "txt file path").Default("edu.txt").ExistingFile()
	jPath   = kingpin.Flag("jPath", "json file path").Default("edu.json").ExistingFile()
	compile = kingpin.Flag("compile", "json to txt").Bool()
)

func main() {
	kingpin.Parse()

	if *compile {
		txt2json()
	}

	f, err := os.Open(*jPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var items []item
	json.Unmarshal(b, &items)
	for _, v := range items {
		fmt.Printf("Tag:%s	Name:%s	Domain:%s \n", v.Tag, v.Name, v.Domain)
	}
	fmt.Printf("本库共收录%d所高校\n", len(items))
}

func txt2json() {
	f, err := os.Open(*tPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var items []item
	var str = strings.Split(string(b), "\n")
	for _, s := range str {
		l := strings.Split(s, "\t")
		if len(l) > 2 {
			items = append(items, item{
				Tag:    strings.Trim(strings.TrimLeft(l[0], "Tag:"), " "),
				Name:   strings.Trim(strings.TrimLeft(l[1], "Name:"), " "),
				Domain: strings.Trim(strings.TrimLeft(l[2], "Domain:"), " "),
			})
		} else {
			fmt.Println(len(l))
			fmt.Printf("Invalid format %s\n", s)
		}
	}
	sj, _ := json.Marshal(&items)
	ioutil.WriteFile(*jPath, []byte(sj), 0644)
	fmt.Printf("本库共收录%d所高校\n", len(items))
	os.Exit(0)
}
