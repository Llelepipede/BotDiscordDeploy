package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Api_stud []Api
)

type Guild struct {
	Nom    string
	Point  int
	Membre []Api
}

type Complete_Stud struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Point  int    `json:"point"`
}

type Api struct {
	Id    int    `json:"id"`
	Guild string `json:"guild"`
	Point int    `json:"point"`
}

type Studient struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
}

const api_link = "https://mentor-paris.github.io/jsonApiGuildbot/api.json"

func GetApi() error {

	//Call the KuteGo API and retrieve our cute Dr Who Gopher
	response, err := http.Get(api_link)
	var variable []byte

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	file := response.Body

	file.Read(variable)

	err = json.Unmarshal(variable, &Api_stud)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

func CreateApi() ([]byte, error) {

	var studs [100]Api

	for i := 0; i < 100; i++ {
		studs[i].Id = i
		studs[i].Guild = "none"
		studs[i].Point = 0
	}

	file, err := json.Marshal(studs)

	ioutil.WriteFile("api.json", file, 0777)

	if err != nil {
		log.Fatal(err)
		return file, err
	}

	return file, err
}

func CreateStud() ([]byte, error) {

	var studs [100]Studient

	for i := 0; i < 100; i++ {
		studs[i].Id = i
		studs[i].Nom = "None"
		studs[i].Prenom = "None"
	}

	file, err := json.Marshal(studs)

	ioutil.WriteFile("stud.json", file, 0777)

	if err != nil {
		log.Fatal(err)
		return file, err
	}

	return file, err
}
