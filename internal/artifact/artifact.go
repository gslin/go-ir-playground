package artifact

import (
	"encoding/json"
	"io/ioutil"
)

type Article struct {
	Id string
	Body string
	Title string
}

func Read(filename string) []Article {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var payload []Article
	err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}

	return payload
}
