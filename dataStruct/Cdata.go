package dataStruct

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func CreateApi() ([]byte, error) {

	var studs [100]Api

	for i := 0; i < 100; i++ {
		studs[i].Id = i
		studs[i].Guild = "none"
		studs[i].Point = 0
		studs[i].Credit = 10
	}

	file, err := json.Marshal(studs)

	ioutil.WriteFile("ApiData/api.json", file, 0777)

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

func CreateCStud(api []Api, stud []StudentsData) []Complete_Stud {
	var ret []Complete_Stud

	for i := range stud {
		var new_CStud Complete_Stud
		new_CStud.Credit = api[i].Credit
		new_CStud.Point = api[i].Point
		new_CStud.Guild = api[i].Guild
		new_CStud.Prenom = stud[i].Prenom
		new_CStud.Nom = stud[i].Nom
		new_CStud.Id = api[i].Id
		new_CStud.Discord = stud[i].ID_discord
		ret = append(ret, new_CStud)
	}
	return ret
}
