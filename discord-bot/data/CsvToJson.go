package data

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

type StudentsData struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Nom string `json:"nom"`
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
				}
			}
			StudentsList = append(StudentsList, rec)
		}
	}
	return StudentsList
}

func StudDataGet() []Studient {

	// open file
	f, err := os.Open("Students_cleaning.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// 2. Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Assign successive lines of raw CSV data to fields of the created structs
	StudentsList := createStudentsList(data)

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(StudentsList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	var stud *[]Studient

	json.Unmarshal(jsonData, stud)

	return *stud
	// fmt.Println(string(jsonData))
}
