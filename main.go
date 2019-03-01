package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type school struct {
	URL  string `json:"url"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

var (
	jPath = "./edu.json"
	tPath = "./edu.txt"
	// TODO 增加一个学校
	// append  = kingpin.Flag("append", "append one school").Bool()
	j2t = kingpin.Flag("j2t", "json to txt").Bool()
	t2j = kingpin.Flag("t2j", "txt to json").Bool()
)

func main() {
	kingpin.Parse()

	if *j2t {
		json2text()
	}
	if *t2j {
		text2json()
	}
	f, err := os.Open(jPath)
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
	for _, v := range schools {
		fmt.Printf("Tag:%s	Name:%s	URL:%s \n", v.Tag, v.Name, v.URL)
	}
	fmt.Printf("本库共收录%d所高校\n", len(schools))
}

func json2text() {
	f, err := os.Open(jPath)
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
	var str string
	for _, v := range schools {
		str = str + fmt.Sprintf("Tag:%s	Name:%s	URL:%s \n", v.Tag, v.Name, v.URL)
	}
	ioutil.WriteFile(tPath, []byte(str), 0644)
	fmt.Println("写入完毕")
	os.Exit(0)
}

func text2json() {
	f, err := os.Open(tPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	// fmt.Println(b)
	var schools []school
	var str = strings.Split(string(b), "\n")
	for _, s := range str {
		l := strings.Split(s, "	")
		if len(l) == 3 {
			l[2] = strings.Trim(l[2], " ")
			if l[2][len(l[2])-1:len(l[2])] == "/" {
				l[2] = strings.Trim(l[2], "/")
			}
			schools = append(schools, school{Tag: fmt.Sprintf(l[0][4:]), Name: fmt.Sprintf(l[1][5:]), URL: fmt.Sprintf("%s/", l[2][4:(len(l[2]))])})
		}
	}
	sj, _ := json.Marshal(&schools)
	ioutil.WriteFile(jPath, []byte(sj), 0644)
	os.Exit(0)
}
