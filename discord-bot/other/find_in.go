package other

import (
	"golang-discord-bot/data"
	"strconv"
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

func Split(str string) []string {
	runes := []rune(str)
	count := 0
	prevIsLeter := false
	for _, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			if prevIsLeter {
				count++
				prevIsLeter = false
			}
		} else {
			prevIsLeter = true
		}
	}
	if prevIsLeter {
		count++
	}
	resArr := make([]string, count)
	i := 0
	start := 0
	prevIsLeter = false
	for ind, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			if prevIsLeter {
				resArr[i] = string(runes[start:ind])
				i++
				prevIsLeter = false
			}
			start = ind + 1
		} else {
			prevIsLeter = true
		}
	}
	if prevIsLeter {
		resArr[i] = string(runes[start:])
	}
	return resArr
}

func Find_Guild(in []data.Api, stud []data.Studient) ([]data.Guild, error) {
	var ret []data.Guild
	var new_guild data.Guild
	var new_member data.Complete_Stud
	var bool bool
	var guild_id int
	for _, v := range in {

		guild_id = 0
		bool = true
		for y, m := range ret {
			if v.Guild == m.Nom {
				bool = false
				guild_id = y
				break
			}
		}
		if bool {
			new_guild.Nom = v.Guild

			new_member.Id = v.Id
			new_member.Point = v.Point

			new_member.Nom = stud[v.Id].Nom
			new_member.Prenom = stud[v.Id].Prenom

			new_guild.Membre = append(new_guild.Membre, new_member)

			new_guild.Point = v.Point

			ret = append(ret, new_guild)
		} else {
			new_member.Id = v.Id
			new_member.Point = v.Point

			new_member.Nom = stud[v.Id].Nom
			new_member.Prenom = stud[v.Id].Prenom

			ret[guild_id].Membre = append(ret[guild_id].Membre, new_member)

			ret[guild_id].Point += v.Point
		}
	}
	return ret, nil
}

func List_Guild(allGuild []data.Guild) string {
	to_string := ""
	for _, v := range allGuild {
		to_string += "la guilde " + v.Nom + " à : " + strconv.Itoa(v.Point) + " points\n"
	}
	return to_string
}
