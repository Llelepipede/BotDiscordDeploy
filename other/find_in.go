package other

import (
	"golang-discord-bot/dataStruct"
	"reflect"
	"strconv"
	"strings"

	"github.com/apex/log"
)

func Find_in_stud(to_find string, in []dataStruct.Complete_Stud, by_what string) (int, bool) {
	var ret int
	for i, v := range in {
		if strings.EqualFold(v.Nom, to_find) && (by_what == "Nom" || by_what == "nom") {
			return i, true
		}
		if strings.EqualFold(v.Prenom, to_find) && (by_what == "Prenom" || by_what == "prenom") {
			return i, true
		}
	}
	return ret, false
}

func Find_in_api(to_find string, in []dataStruct.Api) (int, bool) {
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

func Splitdot(str string, split string) []string {
	runes := str
	count := 0
	prevIsLeter := false
	for i := 0; i < len(runes)-(len(split)); i++ {
		if runes[i:i+len(split)] == split {
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
	y := 0
	start := 0
	prevIsLeter = false
	for i := 0; i < len(runes)-(len(split)); i++ {
		if runes[i:i+len(split)] == split {

			if prevIsLeter {
				resArr[y] = string(runes[start:i])
				y++
				prevIsLeter = false
			}
			start = i + len(split)
		} else {
			prevIsLeter = true
		}
	}
	if prevIsLeter {
		resArr[y] = string(runes[start:])
	}
	return resArr
}

func Find_Guild(in []dataStruct.Complete_Stud) ([]dataStruct.Guild, error) {
	var ret []dataStruct.Guild
	var new_guild *dataStruct.Guild
	var new_member *dataStruct.Complete_Stud
	var is_in bool
	var guild_id int
	for i := range in {
		new_guild = new(dataStruct.Guild)
		new_member = new(dataStruct.Complete_Stud)
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
			new_guild.Nom = in[i].Guild

			new_member.Id = in[i].Id
			new_member.Point = in[i].Point
			new_member.Credit = in[i].Credit

			new_guild.Membre = append(new_guild.Membre, *new_member)

			new_guild.Point = in[i].Point

			ret = append(ret, *new_guild)
		} else {
			new_member.Id = in[i].Id
			new_member.Point = in[i].Point
			new_member.Credit = in[i].Credit

			ret[guild_id].Membre = append(ret[guild_id].Membre, *new_member)

			ret[guild_id].Point += in[i].Point
		}
	}
	return ret, nil
}
func List_Guild(allGuild []dataStruct.Guild) string {
	to_string := ""
	for _, v := range allGuild {

		to_string += "``` " + v.Nom + " \n   ->" + strconv.Itoa(v.Point) + " points\n```"
	}
	return to_string
}

// in = mon api (id,point,guilde)    with = mes etudiants (nom,prenom)
func Tri_Stud(in []dataStruct.Complete_Stud) ([]dataStruct.Complete_Stud, error) {
	cpy_in := make([]dataStruct.Complete_Stud, len(in))
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

func Tri_Guild(in []dataStruct.Guild) ([]dataStruct.Guild, error) {
	cpy_in := make([]dataStruct.Guild, len(in))
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

func List_Membre(Guild dataStruct.Guild, with []dataStruct.StudentsData) string {
	to_string := ""
	for _, v := range Guild.Membre {
		to_string += "nom:" + with[v.Id].Nom + " " + with[v.Id].Prenom + "\n"
	}
	return to_string
}

//get
func Get(comment string, from []dataStruct.Complete_Stud) []string {
	var get []int
	var finded bool
	splited := Split(comment)
	for i, v := range splited {
		if i != 0 {
			finded = false
			if strings.ToUpper(v) != "WHERE" {
				if strings.ToUpper(v) == "ALL" {
					var get []int
					for j := 0; j < reflect.ValueOf(&from[0]).Elem().NumField(); j++ {
						get = append(get, j)
					}

					return Where(splited[i+1:], from, get)
				} else {
					Sample := reflect.ValueOf(&from[0]).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(v, Sample.Type().Field(i).Name) {
							finded = true
							get = append(get, i)
						}
					}
					if !finded {
						return nil
					}
				}
			} else {
				return Where(splited[i:], from, get)
			}
		}
	}
	return Where(splited[len(splited):], from, get)
}

func Where(splited []string, from []dataStruct.Complete_Stud, get []int) []string {
	var whereInd int
	var opperand string
	var filter []bool

	if len(splited) == 0 {
		for range from {
			filter = append(filter, true)
		}
		return Show(from, get, filter)
	}
	if strings.ToUpper(splited[0]) != "WHERE" {
		for range from {
			filter = append(filter, true)
		}
		return Show(from, get, filter)
	}
	if len(splited)%4 != 0 {
		return nil
	}
	for j, v := range splited {
		if j != 0 {
			if j%4 == 0 {
				if strings.EqualFold(v, "add") {
					return nil
				}
			} else if j%4 == 1 {
				Sample := reflect.ValueOf(&from[0]).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if strings.EqualFold(v, Sample.Type().Field(i).Name) {
						whereInd = i
					}
				}
			} else if j%4 == 2 {
				opperand = v
			} else {
				if j/4 == 0 {
					filter = Opperator(whereInd, opperand, v, from)
				} else {
					filter = WhereAdd(filter, Opperator(whereInd, opperand, v, from))
				}

			}

		}
	}
	return Show(from, get, filter)
}

func Show(from []dataStruct.Complete_Stud, get []int, filter []bool) []string {
	var ret []string
	if filter == nil {
		return nil
	}
	for j, v := range from {
		temp := ""
		if filter[j] {
			Sample := reflect.ValueOf(&v).Elem()
			for _, w := range get {
				for i := 0; i < Sample.NumField(); i++ {
					if i == w {
						if strings.EqualFold(Sample.Type().Field(i).Type.Name(), "string") {
							temp += Sample.Type().Field(i).Name + " : "
							temp += Sample.Field(i).String() + "\n"
						} else if strings.EqualFold(Sample.Type().Field(i).Type.Name(), "int") {
							temp += Sample.Type().Field(i).Name + " : "
							temp += strconv.Itoa((int(Sample.Field(i).Int()))) + "\n"
						}
					}
				}
			}
			ret = append(ret, temp)
		}
	}
	return ret
}

func Opperator(whereInd int, opperand string, whereCond string, from []dataStruct.Complete_Stud) []bool {
	var retBool []bool
	var toWorkWithInt int
	var toWorkWithStr string
	inted := false
	switch reflect.ValueOf(&from[0]).Elem().Type().Field(whereInd).Type.Name() {
	case "int":
		toWorkWithInt, _ = strconv.Atoi(whereCond)
		inted = true
	case "string":
		toWorkWithStr = whereCond
	}

	if inted {
		switch opperand {
		case "<":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() < int64(toWorkWithInt))
					}
				}
			}
		case "<=":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() <= int64(toWorkWithInt))
					}
				}
			}
		case "=":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() == int64(toWorkWithInt))
					}
				}
			}
		case "==":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() == int64(toWorkWithInt))
					}
				}
			}
		case ">":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() > int64(toWorkWithInt))
					}
				}
			}
		case ">=":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, Sample.Field(i).Int() >= int64(toWorkWithInt))
					}
				}
			}
		default:
			return nil

		}

	} else {
		switch opperand {
		case "=":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, strings.EqualFold(Sample.Field(i).String(), toWorkWithStr))
					}
				}
			}
		case "==":
			for _, c := range from {
				Sample := reflect.ValueOf(&c).Elem()
				for i := 0; i < Sample.NumField(); i++ {
					if i == whereInd {
						retBool = append(retBool, strings.EqualFold(Sample.Field(i).String(), toWorkWithStr))
					}
				}
			}
		default:
			return nil
		}
	}

	return retBool
}

func WhereAdd(first []bool, second []bool) []bool {
	var ret []bool
	if len(first) != len(second) {
		return nil
	}
	for i := range first {
		ret = append(ret, first[i] && second[i])
	}
	return ret
}

func Add(splited []string, from []dataStruct.Complete_Stud) ([]dataStruct.Complete_Stud, []bool, error) {
	ret := from

	var err error
	var new_value dataStruct.Complete_Stud

	indexType := -1
	indexTypeWhere := -1
	var indexEtud []bool
	value := ""

	for _, v := range splited[0:] {
		splice := Split(v)
		switch splice[0] {
		case "t":
			if len(splice) == 2 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {

							indexType = i
						}
					}
				}
			}
		case "w":
			if len(splice) == 4 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {
							indexTypeWhere = i
						}
					}
				}
				indexEtud = Opperator(indexTypeWhere, splice[2], splice[3], from)
			} else {
				return ret, indexEtud, err
			}
		case "v":
			if len(splice) == 2 {
				value = splice[1]

			}
		default:
			return ret, indexEtud, err
		}
	}
	log.Info("taille de la liste d'étudiant:" + strconv.Itoa(len(indexEtud)))
	log.Info("index de la variable cible:" + strconv.Itoa(indexType))
	log.Info("index de la valeur a attribuer :" + value)
	if len(indexEtud) == 0 || (indexType == -1) || (value == "") {
		return ret, indexEtud, err
	} else {
		for j, c := range from {
			new_value = c
			Sample := reflect.ValueOf(&new_value).Elem()
			if indexEtud[j] {
				for i := 0; i < Sample.NumField(); i++ {
					if i == indexType {
						inted, _ := strconv.Atoi(value)
						log.Info(strconv.Itoa(int(Sample.Field(i).Int() + int64(inted))))
						Sample.Field(i).SetInt(Sample.Field(i).Int() + int64(inted))
					}
				}
				ret[j] = new_value
			}

		}
		return ret, indexEtud, err
	}
}

func Remove(splited []string, from []dataStruct.Complete_Stud) ([]dataStruct.Complete_Stud, []bool, error) {
	ret := from

	var err error
	var new_value dataStruct.Complete_Stud

	indexType := -1
	indexTypeWhere := -1
	var indexEtud []bool
	value := ""

	for _, v := range splited[0:] {
		splice := Split(v)
		switch splice[0] {
		case "t":
			if len(splice) == 2 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {

							indexType = i
						}
					}
				}
			}
		case "w":
			if len(splice) == 4 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {
							indexTypeWhere = i
						}
					}
				}
				indexEtud = Opperator(indexTypeWhere, splice[2], splice[3], from)
			} else {
				return ret, indexEtud, err
			}
		case "v":
			if len(splice) == 2 {
				value = splice[1]

			}
		default:
			return ret, indexEtud, err
		}
	}
	if len(indexEtud) == 0 || (indexType == -1) || (value == "") {
		return ret, indexEtud, err
	} else {
		for j, c := range from {
			new_value = c
			Sample := reflect.ValueOf(&new_value).Elem()
			if indexEtud[j] {
				for i := 0; i < Sample.NumField(); i++ {
					if i == indexType {
						inted, _ := strconv.Atoi(value)
						log.Info(strconv.Itoa(int(Sample.Field(i).Int() - int64(inted))))
						Sample.Field(i).SetInt(Sample.Field(i).Int() - int64(inted))
					}
				}
				ret[j] = new_value
			}

		}
		return ret, indexEtud, err
	}
}

func Set(splited []string, from []dataStruct.Complete_Stud) ([]dataStruct.Complete_Stud, []bool, error) {
	ret := from

	var err error
	var new_value dataStruct.Complete_Stud

	indexType := -1
	indexTypeWhere := -1
	var indexEtud []bool
	value := ""

	for _, v := range splited[0:] {
		splice := Split(v)
		switch splice[0] {
		case "s":
			if len(splice) == 4 {
				if splice[2] == "<=" {
					for _, c := range from {
						Sample := reflect.ValueOf(&c).Elem()
						for i := 0; i < Sample.NumField(); i++ {
							if strings.EqualFold(splice[3], Sample.Type().Field(i).Name) {
								indexType = i
							}
						}
					}
					if indexType != -1 {
						value = splice[1]
					} else {
						return ret, indexEtud, err
					}
				} else if splice[2] == "=>" {
					for _, c := range from {
						Sample := reflect.ValueOf(&c).Elem()
						for i := 0; i < Sample.NumField(); i++ {
							if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {
								indexType = i
							}
						}
					}
					if indexType != -1 {
						value = splice[3]
					} else {
						return ret, indexEtud, err
					}
				} else {
					return ret, indexEtud, err
				}
			} else {
				return ret, indexEtud, err
			}
		case "w":
			if len(splice) == 4 {
				for _, c := range from {
					Sample := reflect.ValueOf(&c).Elem()
					for i := 0; i < Sample.NumField(); i++ {
						if strings.EqualFold(splice[1], Sample.Type().Field(i).Name) {
							indexTypeWhere = i
						}
					}
				}
				indexEtud = Opperator(indexTypeWhere, splice[2], splice[3], from)
			} else {
				return ret, indexEtud, err
			}
		default:
			return ret, indexEtud, err
		}
	}
	if len(indexEtud) == 0 || (indexType == -1) || (value == "") {
		return ret, indexEtud, err
	} else {
		for j, c := range from {
			new_value = c
			Sample := reflect.ValueOf(&new_value).Elem()
			if indexEtud[j] {
				for i := 0; i < Sample.NumField(); i++ {
					if i == indexType {
						if Sample.Type().Field(i).Type.Name() == "string" {
							log.Info("attribution d'un string")
							Sample.Field(i).SetString(value)
						} else if Sample.Type().Field(i).Type.Name() == "int" {
							inted, err := strconv.Atoi(value)
							if err != nil {
								return ret, indexEtud, err
							}
							log.Info("attribution d'un int")
							Sample.Field(i).SetInt(int64(inted))
						}
					}
				}
				ret[j] = new_value
			}

		}
		return ret, indexEtud, err
	}
}
