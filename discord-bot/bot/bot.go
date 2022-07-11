package bot

import (
	"encoding/json"
	"fmt"
	"golang-discord-bot/config"
	"golang-discord-bot/data"
	"golang-discord-bot/gitmanage"
	"golang-discord-bot/other"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {

	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}

	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}

	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// Commande qui ajoute des point a l"etudiant
	if StartWith(m.Content, "gain") {
		response := other.Split(m.Content)

		// verifie qu'il y a bien un certain nombre de paramètre
		if len(response) <= 2 {
			message := other.C_embed("ERROR", "Pas assez d'argument pour la commande \"gain\"", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			// lie la liste des etudiant pour y trouver l'etudiant correspondant a la requète
			stud_list, err := config.ReadStud()

			// file, err := ioutil.ReadFile("./stud.json")
			// var stud_list *[]data.Studient

			// // si on arrive pas a ouvrir le fichier json, renvoie une erreur
			// if err != nil {
			// 	message := other.C_embed("ERROR", "Probleme de lecture du Json étudiant", config.Color_error)
			// 	adresse_m := &message
			// 	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			// 	return
			// }

			// // si on arrive pas a Unmarshal le json, renvoie une erreur
			// err = json.Unmarshal(file, &stud_list)
			if err != nil {
				message := other.C_embed("ERROR", "Probleme dand la lecture du fichier étudiant", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				return
			}
			// cherche dans le json l'eudiant en argument
			id_stud, nope := other.Find_in_stud(response[1], *stud_list, "nom")
			if !nope {
				message := other.C_embed("ERROR", "étudiant introuvable", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				_, nope = gitmanage.Pull()
				if !nope {
					message := other.C_embed("ERROR", "Le Pull a echoué", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				} else {
					api := config.ReadApi()
					to_add, _ := strconv.Atoi(response[2])
					api[id_stud].Point += to_add

					file, err := json.Marshal(api)
					ioutil.WriteFile("./ApiData/api.json", file, 0777)

					if err != nil {
						message := other.C_embed("ERROR", "l'ecriture dans le fichier api.json a échoué", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					} else {
						message := other.C_embed("GAIN DE POINT", (*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" \nID: "+strconv.Itoa(id_stud), config.Color_reponse)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
					}
				}
			}
		}
	}

	if StartWith(m.Content, "point") {
		response := other.Split(m.Content)

		// verifie qu'il y a bien un certain nombre de paramètre
		if len(response) <= 1 {
			_, nope := gitmanage.Pull()
			if !nope {
				message := other.C_embed("ERROR", "Le Pull a echoué", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				api := config.ReadApi()
				stud, _ := config.ReadStud()
				to_print := ""
				for _, v := range api {

					to_print += (*stud)[v.Id].Nom + " " + (*stud)[v.Id].Prenom + " : " + strconv.Itoa(v.Point) + " | guilde : " + v.Guild + "\n"
				}
				message := other.C_embed("LISTE DES POINT", to_print, config.Color_reponse)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)

			}
		} else if StartWith(response[1], "guilde") {
			api := config.ReadApi()
			stud, err := config.ReadStud()
			if err != nil {
				message := other.C_embed("ERROR", "Probleme dand la lecture du fichier étudiant", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				return
			} else {
				allGuild, err := other.Find_Guild(api, *stud)
				if err != nil {
					message := other.C_embed("ERROR", "Probleme dans la recherche des guilde au sein de l'API", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					return
				} else {
					message := other.C_embed("LISTE DES GUILDE", other.List_Guild(allGuild), config.Color_reponse)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					channel, _ := s.UserChannelCreate(m.Author.ID)
					_, _ = s.ChannelMessageSendEmbed(channel.ID, adresse_m)
				}
			}

		} else {
			message := other.C_embed("ERROR", "Argument incorrecte pour la commande \"gain\"", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		}
	}

	if StartWith(m.Content, "changeguilde") {
		response := other.Split(m.Content)
		// verifie qu'il y a bien un certain nombre de paramètre
		if len(response) <= 2 {
			message := other.C_embed("ERROR", "Pas assez d'argument pour la commande \"changeguilde\"", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			message := other.C_embed("ATTRIBUT UNE GUILDE ", "change la guilde", config.Color_reponse)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		}
	}

	// if StartWith(m.Content, "embed") {
	// 	description :=
	// 		"Lorem ipsum dolor sit amet consectetur,\n adipisicing elit. Omnis corporis fuga ducimus ea incidunt? Atque quo sint aliquam. Debitis,\n quidem. Rem dolor, modi ad labore natus porro quas quae ducimus?"
	// 	message := other.C_embed("SELECTION D'ETUDIANT", description, 1)

	// 	adresse := &message
	// 	fmt.Printf("message: %v\n", message)
	// 	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse)
	// }

	if StartWith(m.Content, "push") {

		response := other.Split(m.Content)
		if len(response) <= 1 {
			message := other.C_embed("PUSHING", "pushing ...", config.Color_reponse)
			adresse_m := &message
			embed_message, _ := s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			if gitmanage.Commit() {
				message := other.C_embed("PUSHING", "pushing ... succeed", config.Color_reponse)
				adresse_m := &message
				s.ChannelMessageEditEmbed(m.ChannelID, embed_message.ID, adresse_m)
				// s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... succeed")
			} else {
				message := other.C_embed("PUSHING", "pushing ... failed", config.Color_reponse)
				adresse_m := &message
				s.ChannelMessageEditEmbed(m.ChannelID, embed_message.ID, adresse_m)
				// s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... failed")
			}

		}
	}

	if StartWith(m.Content, "/createStud") {
		response := other.Split(m.Content)
		if len(response) <= 1 {

			_, _ = s.ChannelMessageSend(m.ChannelID, "creating the json file...")
			_, err := data.CreateStud()
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "création echoué")
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, "création reussi")
			}

		}
	}
}

func StartWith(content string, patern string) bool {
	if len(patern) > len(content) {
		return false
	}
	for i, v := range patern {
		if rune(content[i]) != v {
			return false
		}
	}
	return true
}
