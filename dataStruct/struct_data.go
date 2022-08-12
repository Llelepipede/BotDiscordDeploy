package dataStruct

type StudentsData struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Nom        string `json:"nom"`
	Prenom     string `json:"prenom"`
	ID_discord string `json:"ID"`
}

type Guild struct {
	Nom    string
	Point  int
	Membre []Complete_Stud
}

type Complete_Stud struct {
	Id      int    `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Point   int    `json:"point"`
	Credit  int    `json:"credit"`
	Guild   string `json:"guild"`
	Discord string `json:"ID_disc"`
}

type Logs struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Log    []Log  `json:"log"`
}

type Log struct {
	Date    string `json:"date"`
	Mentor  string `json:"mentor"`
	Comment string `json:"comment"`
	Clause  string `json:"clause"`
}

type Api struct {
	Id     int    `json:"id"`
	Guild  string `json:"guild"`
	Point  int    `json:"point"`
	Credit int    `json:"credit"`
}

type Studient struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
}
