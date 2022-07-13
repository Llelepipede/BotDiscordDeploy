package config

import (
	"encoding/json"
	"golang-discord-bot/data"
	"io/ioutil"
	"log"
)

var (
	Token         string
	Prefix        string
	config        *configStruct
	Color_error   int
	Color_reponse int
	Admin_chanel  string
)

type configStruct struct {
	Token         string `json:"token"`
	Prefix        string `json:"prefix"`
	Color_error   int    `json:"color_error"`
	Color_reponse int    `json:"color_reponse"`
	Admin_chanel  string `json:"admin_chanel"`
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
	Admin_chanel = config.Admin_chanel
	return nil
}

func ReadApi() []data.Api {
	var ret []data.Api

	file, err := ioutil.ReadFile("./ApiData/api.json")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = json.Unmarshal(file, &ret)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return ret
}

func ReadStud() (*[]data.Studient, error) {
	// lie la liste des etudiant pour y trouver l'etudiant correspondant a la requète
	file, err := ioutil.ReadFile("./stud.json")
	var stud_list *[]data.Studient

	// si on arrive pas a ouvrir le fichier json, renvoie une erreur
	if err != nil {
		return nil, nil
	}

	// si on arrive pas a Unmarshal le json, renvoie une erreur
	err = json.Unmarshal(file, &stud_list)
	if err != nil {

		return nil, nil
	}
	return stud_list, nil
}
