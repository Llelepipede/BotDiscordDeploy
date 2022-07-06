package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Token         string
	Prefix        string
	config        *configStruct
	Color_error   int
	Color_reponse int
)

type configStruct struct {
	Token         string `json:"token"`
	Prefix        string `json:"prefix"`
	Color_error   int    `json:"color_error"`
	Color_reponse int    `json:"color_reponse"`
}

func ReadConfig() error {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	Token = config.Token
	Prefix = config.Prefix
	Color_error = config.Color_error
	Color_reponse = config.Color_reponse

	return nil
}
