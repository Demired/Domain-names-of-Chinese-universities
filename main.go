package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type school struct {
	URL  string `json:"url"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

var (
	schools []school
	jPath   = "./edu.json"
	tPath   = "./edu.txt"
	j2t     = kingpin.Flag("j2t", "install program").Bool()
)

func main() {
	kingpin.Parse()

	f, err := os.Open(jPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &schools)
	if *j2t {
		json2txt()
	}
	for _, v := range schools {
		fmt.Printf("Tag:%s	Name:%s	URL:%s \n", v.Tag, v.Name, v.URL)
	}
	fmt.Printf("本库共收录%d所高校\n", len(schools))
}

func json2txt() {
	var str string
	for _, v := range schools {
		str = str + fmt.Sprintf("Tag:%s	Name:%s	URL:%s \n", v.Tag, v.Name, v.URL)
	}
	ioutil.WriteFile(tPath, []byte(str), 0644)
	fmt.Println("写入完毕")
	os.Exit(0)
}
