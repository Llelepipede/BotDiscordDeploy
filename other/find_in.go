package other

import (
	"fmt"
	"golang-discord-bot/data"
	"strconv"
)

func Find_in_stud(to_find string, in []data.StudentsData, by_what string) (int, bool) {
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

func Find_Guild(in []data.Api, stud []data.StudentsData) ([]data.Guild, error) {
	var ret []data.Guild
	new_guild := new(data.Guild)
	new_member := new(data.Api)
	var is_in bool
	var guild_id int
	for i, v := range stud {
		new_guild = new(data.Guild)
		new_member = new(data.Api)
		fmt.Printf("i: %v\n", i)
		guild_id = 0
		is_in = true
		for y, m := range ret {
			if in[i].Guild == m.Nom {
				is_in = false
				guild_id = y
				break
			}
		}
		if is_in {
			fmt.Printf("création guilde %v\n", v.Nom+" "+in[i].Guild)
			new_guild.Nom = in[i].Guild

			new_member.Id = in[i].Id
			new_member.Point = in[i].Point

			new_guild.Membre = append(new_guild.Membre, *new_member)

			new_guild.Point = in[i].Point

			ret = append(ret, *new_guild)
		} else {
			fmt.Printf("update %v\n", v.Nom+" "+in[i].Guild)
			new_member.Id = in[i].Id
			new_member.Point = in[i].Point

			ret[guild_id].Membre = append(ret[guild_id].Membre, *new_member)

			ret[guild_id].Point += in[i].Point
		}
	}
	return ret, nil
}
func List_Guild(allGuild []data.Guild) string {
	to_string := ""
	for _, v := range allGuild {

		to_string += "``` " + v.Nom + " \n   ->" + strconv.Itoa(v.Point) + " points\n```"
	}
	return to_string
}

// in = mon api (id,point,guilde)    with = mes etudiants (nom,prenom)
func Tri_Stud(in []data.Api, with []data.StudentsData) ([]data.Api, error) {
	cpy_in := make([]data.Api, len(in))
	copy(cpy_in, in)

	cpy_with := make([]data.StudentsData, len(with))
	copy(cpy_with, with)

	is_tri := false
	for !is_tri {
		is_tri = true
		for i := 0; i < len(cpy_with)-1 && i < len(cpy_in)-1; i++ {
			if cpy_in[i].Point < cpy_in[i+1].Point {
				dump := cpy_in[i]
				cpy_in[i] = cpy_in[i+1]
				cpy_in[i+1] = dump
				is_tri = false
			}
		}
	}

	return cpy_in, nil
}

func Tri_Guild(in []data.Guild) ([]data.Guild, error) {
	cpy_in := make([]data.Guild, len(in))
	copy(cpy_in, in)

	is_tri := false
	for !is_tri {
		is_tri = true
		for i := 0; i < len(cpy_in)-1; i++ {
			if cpy_in[i].Point < cpy_in[i+1].Point {
				dump := cpy_in[i]
				cpy_in[i] = cpy_in[i+1]
				cpy_in[i+1] = dump
				is_tri = false
			}
		}
	}
	return cpy_in, nil
}

func List_Membre(Guild data.Guild, with []data.StudentsData) string {
	to_string := ""
	for _, v := range Guild.Membre {
		to_string += "nom:" + with[v.Id].Nom + " " + with[v.Id].Prenom + "\n"
	}
	return to_string
}
