package other

import (
	"golang-discord-bot/data"
)

func Find_in_stud(to_find string, in []data.Studient, by_what string) (int, bool) {
	var ret int
	for i, v := range in {
		if v.Nom == to_find && (by_what == "Nom" || by_what == "nom") {
			return i, true
		}
		if v.Prenom == to_find && (by_what == "Prenom" || by_what == "prenom") {
			return i, true
		}
	}
	return ret, false
}

func Find_in_api(to_find string, in []data.Api) (int, bool) {
	var ret int
	for i, v := range in {
		if v.Guild == to_find {
			return i, true
		}
	}
	return ret, false
}
