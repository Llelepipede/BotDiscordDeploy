package config

import (
	"bytes"
	"encoding/json"
	"golang-discord-bot/dataStruct"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	Token         string
	Prefix        string
	config        *configStruct
	Color_error   int
	Color_reponse int
	Admin_chanel  string
	ApiLink       string
	ApiLog        string
	Aide_report   string
	Aide_remove   string
	Aide_add      string
	Aide_top      string
	Aide_get      string
	Aide_point    string
	Aide_set      string
	Aide_push     string
	Aide_initLogs string
)

type configStruct struct {
	Token         string `json:"token"`
	Prefix        string `json:"prefix"`
	Color_error   int    `json:"color_error"`
	Color_reponse int    `json:"color_reponse"`
	Admin_chanel  string `json:"admin_chanel"`
	ApiLink       string `json:"apiLink"`
	ApiLog        string `json:"apiLog"`
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
	ApiLink = config.ApiLink
	ApiLog = config.ApiLog
	return nil
}

func ReadApi() ([]dataStruct.Complete_Stud, []dataStruct.Complete_Stud) {
	ret := new([]dataStruct.Complete_Stud)
	ret2 := new([]dataStruct.Complete_Stud)
	link := config.ApiLink
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, ret)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, ret2)
	if err != nil {
		log.Fatal(err)
	}
	return *ret, *ret2
}

func ReadLogs() []dataStruct.Logs {
	var ret []dataStruct.Logs

	file, err := ioutil.ReadFile("./ApiData/logsGeneral.json")

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

func ReadStud() (*[]dataStruct.Studient, error) {
	// lie la liste des etudiant pour y trouver l'etudiant correspondant a la requète
	file, err := ioutil.ReadFile("./stud.json")
	var stud_list *[]dataStruct.Studient

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

func UpdateStud(api []dataStruct.Complete_Stud, filter []bool) error {

	client := &http.Client{}
	for y, v := range api {
		if !filter[y] {
			continue
		}
		payload, err := json.Marshal(v)
		to_ret := ApiLink
		Sample := reflect.ValueOf(&v).Elem()
		to_ret += "?"
		for i := 0; i < Sample.NumField(); i++ {
			if i != 0 {
				to_ret += "&"
			}
			to_ret += Sample.Type().Field(i).Name + "="
			if strings.EqualFold(Sample.Type().Field(i).Type.Name(), "string") {
				to_ret += Sample.Field(i).String()
			}
			if strings.EqualFold(Sample.Type().Field(i).Type.Name(), "int") {
				to_ret += strconv.Itoa(int(Sample.Field(i).Int()))

			}
		}
		if err != nil {
			log.Fatal(err)
			return err
		}

		// 3.
		req, err := http.NewRequest(http.MethodPatch, to_ret, bytes.NewBuffer(payload))
		req.Header.Set("Accept", " application/json")
		req.Header.Set("Authorization", "Bearer rnd_29L4LSmlKTj6GjcEuO9YJU66su8F")
		if err != nil {
			log.Fatal(err)
			return err
		}

		// 4.
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// 5.
		defer resp.Body.Close()

		// 6.
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		// log.Println("-------------")
		// log.Println(string(body))
		// fmt.Printf("v: %v\n", v)
		// fmt.Printf("req: %v\n", req)

	}
	return nil
}

func Aide() {
	Aide_report = "``` report\n\t\t-> Ajoute un log a l'étudiant cible. Necessite le paramètre -c, -m, -e, -t, et -v\n```" +
		"```\t-e\t-> La chaine de caractère suivante sera considéré comme le nom et prenom de l'étudiant , quelque que soit l'ordre (2 elements en général)\n```" +
		"```\t-t\t-> La chaine de caractère suivante sera considéré comme le nom de la variable a modifié. (1 seul element)\n```" +
		"```\t-v\t-> La chaine de caractère suivante sera considéré comme la valeur a ajouté a la variable t. (1 seul element)\n```" +
		"```\t-c\t-> La chaine de caractère suivante sera considéré comme l'effet du fonctionement de la commande (remove ou add). (1 seul element)\n```" +
		"```\t-m\t-> La chaine de caractère suivante sera considéré comme le mentor ayant ecrit ce message. (X elements)\n```" +
		"```\t-d\t-> La chaine de caractère suivante sera considéré comme la description du report. (X elements)\n```" +
		"```\tEXEMPLE: Augmenter les crédits de l'étudiant Paul de 100 et le certifié dans les logs\n\t\t-> $:report -d à craché sur la table -m Rayane Mokri -c add -e Paul CHESNEAU -t credit -v 100```\n"
	Aide_remove = "``` remove\n\t\t-> Retire une valeur v a une variable t l'etudiant cible. Necessite le paramètre -e (ou -w), -t, et -v\n```" +
		"```\t-e\t-> La chaine de caractère suivante sera considéré comme le nom et prenom de l'étudiant , quelque que soit l'ordre (2 elements en général)\n```" +
		"```\t-w\t-> La chaine de caractère suivante sera considéré comme la condition suffisante pour affecté l'etudiant par la modification. implique que la commande puisse cibler plusieurs personnes. (3 element de la forme {Variable Condition Valeur})\n```" +
		"```\t-t\t-> La chaine de caractère suivante sera considéré comme le nom de la variable a modifié. (1 seul element)\n```" +
		"```\t-v\t-> La chaine de caractère suivante sera considéré comme la valeur a retirer a la variable t. (1 seul element)\n```" +
		"```\tEXEMPLE: Reduire les crédits de l'étudiant Paul de 100\n\t\t-> $:remove -e Paul CHESNEAU -t credit -v 100```\n\n"
	Aide_add = "``` add\n\t\t-> Ajoute une valeur v a une variable t à l'étudiant cible. Necessite le paramètre -e (ou -w), -t, et -v\n```" +
		"```\t-e\t-> La chaine de caractère suivante sera considéré comme le nom et prenom de l'étudiant , quelque que soit l'ordre (2 elements en général)\n```" +
		"```\t-w\t-> La chaine de caractère suivante sera considéré comme la condition suffisante pour affecté l'etudiant par la modification. implique que la commande puisse cibler plusieurs personnes. (3 element de la forme {Variable Condition Valeur})\n```" +
		"```\t-t\t-> La chaine de caractère suivante sera considéré comme le nom de la variable a modifié. (1 seul element)\n```" +
		"```\t-v\t-> La chaine de caractère suivante sera considéré comme la valeur a ajouté a la variable t. (1 seul element)\n```" +
		"```\tEXEMPLE: Augmenter les points de l'étudiant Paul de 50\n\t\t-> $:add -e Paul CHESNEAU -t point -v 50```\n\n"
	Aide_top = "``` top \n\t\t-> visualise le top 5\n```" +
		"```\tEXEMPLE: voir le top 5\n\t\t-> $:top```\n\n"
	Aide_get = "``` get\n\t\t-> Recupére les étudiant en accord a la requète, et affiche les données demandées. Si aucune requète \"where\" est demandé, tout les étudiants seron affiché. \n```" +
		"```\t{données à récupérer | all}\t-> Suite de toute les données a afficher. Si \"all\" est écrit, recupére toute les données (X elements ou 1 seul pour all)\n```" +
		"```\t{where}\t-> Paramètre suplémentaire , permetant de faire un filtre au sein de la selection. Ce filtre n'est pas obligatoire pour le bon fonctionnement de \"get\" (1 + 3 elements obligatoire apres le \"where\")\n```" +
		"```\t{variable {operateur} valeur}\t-> Compare le contenue de la variable avec la valeur selon l'opperateur logique. Si cette comparaison est vrais, la ligne sera afficher dans la visualisation final (3 elements obligatoire apres le \"where\")\n```" +
		"```\t{and}\t-> permet de rehitéré les comparaisons, afin d'affiné le filtre (1 + 3 elements obligatoire apres le \"and\")\n```" +
		"```\tEXEMPLE: visualisé les noms et prenom des étudiants ayant des point compris entre 100 et 1000\n\t\t-> $:get nom prenom where point <= 1000 and point >= 100```\n"
	Aide_point = "``` point\n\t\t-> visualise les points de l'utilisateur de la commande.\n```" +
		"```\t{-g | -l}\t-> paramètre suplémentaire. visualise les points des guildes, ou alors de tout les étudiants. (-l est une commande que seul les mentors peuvent faire) (1 seul element par commande)\n```" +
		"```\tEXEMPLE: visualisé vos propre point\n\t\t-> $:point```\n"
	Aide_set = "``` set\n\t\t-> modifie une variable pour quelle corresponde a la valeur précisé. \n```" +
		"```\t{guild | point | credit}\t-> precise la variable a changer. (1 seule variable par commande)\n```" +
		"```\t{nom_prenom | all}\t-> precise l'étudiant cible. all ciblera tout les étudiants. (1 seul element par commande) \n```" +
		"```\t{valeur}\t-> precise la valeur reference pour la modification. (1 seul element par commande)\n```" +
		"```\tEXEMPLE: changer la guilde de l'étudiant GOMIS pour quelle soit \"Yeti\"\n\t\t-> $:set guild GOMIS Yeti```\n"
}
