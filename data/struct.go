package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-discord-bot/dataStruct"
	"golang-discord-bot/other"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	Api_stud []dataStruct.Api
)

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

func AddLogsALL(from []dataStruct.Complete_Stud) {
	// lie la liste des etudiant pour y trouver l'etudiant correspondant a la requète

	var logs []dataStruct.Logs
	var logs_ *dataStruct.Logs
	var log_ *dataStruct.Log

	// si on arrive pas a ouvrir le fichier json, renvoie une erreur

	for i, v := range from {
		log_ = new(dataStruct.Log)
		logs_ = new(dataStruct.Logs)
		logs_.Id = v.Id
		logs_.Nom = v.Nom
		logs_.Prenom = v.Prenom
		log_.Clause = "Création des logs"
		log_.Comment = "- - -"
		log_.Date = time.Now().String()
		log_.Mentor = "groupe mentor"
		logs = append(logs, *logs_)
		logs[i].Log = append(logs[i].Log, *log_)
	}
	file, _ := json.Marshal(logs)

	err := os.WriteFile("ApiData/logsGeneral.json", file, 0777)

	if err != nil {
		log.Print(err)
		return
	}
}

func AddLog(comment string, from []dataStruct.Complete_Stud, logs []dataStruct.Logs) ([]dataStruct.Logs, []dataStruct.Complete_Stud, []bool) {
	ret := logs
	ret1 := from
	var indexOfStud = -1
	var ind []bool
	var err error
	newlog := new(dataStruct.Log)
	clause := ""
	typeOf := ""
	value := ""
	etud := ""
	indexTypeWhere := 0

	splited := other.Splitdot(comment, " -")
	if len(splited) != 7 {
		return nil, nil, ind
	}
	newlog.Mentor = "non renseigné"
	wrong := false
	//report -e {prenom nom} -m {prenom_mentor} -d {comment ...} -c remove credit XXX
	for _, v := range splited[1:] {
		splice := other.Split(v)
		if strings.EqualFold(splice[0], "w") {
			if len(splice) == 4 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {
							indexTypeWhere = i
						}
					}
				}
				ind = other.Opperator(indexTypeWhere, splice[2], splice[3], from)
			} else {
				return ret, ret1, ind
			}
			newlog.Date = time.Now().String()
		} else if strings.EqualFold(splice[0], "m") {
			newlog.Mentor = v[1:]
		} else if strings.EqualFold(splice[0], "d") {
			newlog.Comment = v[1:]
		} else if strings.EqualFold(splice[0], "c") {
			clause = v[1:]
		} else if strings.EqualFold(splice[0], "t") {
			typeOf = v
		} else if strings.EqualFold(splice[0], "v") {
			value = v
		} else {
			wrong = true
		}
		if clause != "" && typeOf != "" && value != "" && etud != "" {
			clause = other.Split(clause)[0]
			newlog.Clause = clause + " -" + etud + " -" + typeOf + " -" + value
			switch other.Splitdot(newlog.Clause, " -")[0] {
			case "remove":
				fmt.Printf("newlog.Clause: %v\n", newlog.Clause)
				ret1, ind, err = other.Remove(other.Splitdot(newlog.Clause, " -")[1:], from)
			case "add":
				ret1, ind, err = other.Add(other.Splitdot(newlog.Clause, " -")[1:], from)
			default:
				err = errors.New("wrong commande")
			}
			clause = ""
		}
	}
	fmt.Printf("wrong: %v\n", wrong)
	if wrong || err != nil || len(ind) != 0 {
		return nil, nil, ind
	} else {
		if indexOfStud != -1 {
			if err != nil {
				return nil, nil, ind
			} else {
				for i, v := range ind {
					if v {
						ret[i].Log = append(ret[i].Log, *newlog)
						fmt.Printf(" OK: %v\n", ret[i].Log)
					}
				}
				return ret, ret1, ind
			}

		} else {
			return nil, nil, ind
		}

	}
}
