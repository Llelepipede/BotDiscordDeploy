package data

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

type StudentsData struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Nom        string `json:"nom"`
	Prenom     string `json:"prenom"`
	ID_discord string `json:"ID"`
}

func createStudentsList(data [][]string) []StudentsData {
	// convert csv lines to array of structs
	var StudentsList []StudentsData
	for i, line := range data {
		if i > 0 { // omit header line
			var rec StudentsData
			for j, field := range line {
				if j == 0 {
					rec.Nom = field
				} else if j == 1 {
					rec.Prenom = field
				} else if j == 2 {
					rec.ID_discord = field
				}
			}
			StudentsList = append(StudentsList, rec)
		}
	}
	return StudentsList
}

func StudDataGet() (*[]StudentsData, error) {

	// retrieve the path directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, errors.New("error")
	}
	// open file
	f, err := os.Open(path + "/data/Students_cleaning.csv")
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error")
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// 2. Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error")
	}

	// 3. Assign successive lines of raw CSV data to fields of the created structs
	StudentsList := createStudentsList(data)

	return &StudentsList, nil
	// fmt.Println(string(jsonData))
}
