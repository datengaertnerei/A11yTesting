package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type pallytasks []struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name"`
	URL         string        `json:"url"`
	Timeout     int           `json:"timeout"`
	Wait        int           `json:"wait"`
	Standard    string        `json:"standard"`
	Ignore      []interface{} `json:"ignore"`
	Actions     []string      `json:"actions"`
	Annotations []struct {
		Type    string `json:"type"`
		Date    int64  `json:"date"`
		Comment string `json:"comment"`
	} `json:"annotations,omitempty"`
	Headers struct {
	} `json:"headers"`
}

func main() {

	var filename string
	flag.StringVar(&filename, "filename", "pally.json", "The Pa11y export file to process")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("not found")
		return
	}
	defer f.Close()
	tasks := pallytasks{}
	err = json.NewDecoder(f).Decode(&tasks)

	for i, task := range tasks {
		task.ID = ""
		task.Annotations = nil
		text, _ := json.Marshal(task)
		ioutil.WriteFile(filename+strconv.Itoa(i), text, 0666)
	}
}
