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
			file, err := ioutil.ReadFile("./stud.json")
			var stud_list *[]data.Studient

			// si on arrive pas a ouvrir le fichier json, renvoie une erreur
			if err != nil {
				message := other.C_embed("ERROR", "Probleme de lecture du Json étudiant", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				return
			}

			// si on arrive pas a Unmarshal le json, renvoie une erreur
			err = json.Unmarshal(file, &stud_list)
			if err != nil {
				message := other.C_embed("ERROR", "Probleme de Unmarshal", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				return
			}
			// cherche dans le json 'eudiant en argument
			id_stud, nope := other.Find_in_stud(response[1], *stud_list, "nom")
			if !nope {
				message := other.C_embed("ERROR", "étudiant introuvable", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				gitmanage.Pull()
				message := other.C_embed("GAIN DE POINT", (*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" \nID: "+strconv.Itoa(id_stud), config.Color_reponse)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
			}

		}
	}

	if StartWith(m.Content, "embed") {
		description :=
			"Lorem ipsum dolor sit amet consectetur,\n adipisicing elit. Omnis corporis fuga ducimus ea incidunt? Atque quo sint aliquam. Debitis,\n quidem. Rem dolor, modi ad labore natus porro quas quae ducimus?"
		message := other.C_embed("SELECTION D'ETUDIANT", description, 1)

		adresse := &message
		fmt.Printf("message: %v\n", message)
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse)
	}

	if StartWith(m.Content, "push") {

		response := other.Split(m.Content)
		if len(response) <= 1 {
			commit, _ := s.ChannelMessageSend(m.ChannelID, "pushing ...")
			if gitmanage.Commit() {
				s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... succeed")
			} else {
				s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... failed")
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
