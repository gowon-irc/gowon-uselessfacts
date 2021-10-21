package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://uselessfacts.jsph.pl/random.json?language=en"
)

type factJson struct {
	Text string `json:"text"`
}

func (j factJson) String() string {
	return j.Text
}

func fact() (msg string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	j := &factJson{}

	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
