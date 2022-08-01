package bot

import (
	"encoding/json"
	"fmt"
	"golang-discord-bot/config"
	"golang-discord-bot/data"
	"golang-discord-bot/gitmanage"
	"golang-discord-bot/other"
	"io/ioutil"
	"strconv"

	"github.com/apex/log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {

	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err.Error())
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

	admin := false
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	if m.ChannelID == config.Admin_chanel {

		admin = true
	}
	// Commande qui ajoute des point a l"etudiant
	if StartWith(m.Content, "gain") {
		if !admin {
			message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			response := other.Split(m.Content)

			// verifie qu'il y a bien un certain nombre de paramètre
			if len(response) != 3 {
				message := other.C_embed("ERROR", "```argument incorrecte pour la commande \"gain\"\n\n      ->utilisation : gain {nom_etudiant} {to_add```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				// lie la liste des etudiant pour y trouver l'etudiant correspondant a la requète
				stud_list, err := data.StudDataGet()

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
					message := other.C_embed("ERROR", "```Probleme dand la lecture du fichier étudiant```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					return
				}
				// cherche dans le json l'eudiant en argument
				id_stud, nope := other.Find_in_stud(response[1], stud_list, "nom")
				if !nope {
					message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				} else {
					_, nope = gitmanage.Pull()
					if !nope {
						message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					} else {
						api := config.ReadApi()
						to_add, _ := strconv.Atoi(response[2])
						api[id_stud].Point += to_add

						file, err := json.Marshal(api)
						ioutil.WriteFile("./ApiData/api.json", file, 0777)

						if err != nil {
							message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							message := other.C_embed("GAIN DE POINT", "```"+(stud_list)[id_stud].Prenom+" "+(stud_list)[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point total: "+strconv.Itoa(api[id_stud].Point)+"\n   ->Guilde: "+api[id_stud].Guild+"```", config.Color_reponse)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
						}
					}
				}
			}
		}
	}

	// presente les points celon des paramètres.
	if StartWith(m.Content, "point") {
		response := other.Split(m.Content)

		// verifie qu'il y a bien un certain nombre de paramètre
		// si aucun paramètre, envoie en privé le score de l'etudiant.
		if len(response) <= 1 {
			_, nope := gitmanage.Pull()
			if !nope {
				message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				api := config.ReadApi()
				stud, _ := data.StudDataGet()
				to_print := ""
				for i, v := range stud {
					if v.ID_discord == m.Author.ID {
						to_print += "```" + v.Nom + " " + v.Prenom + " \n   -> guilde : " + api[i].Guild + " \n   -> point : " + strconv.Itoa(api[i].Point) + "```\n\n"
						break
					}
				}
				if to_print == "" {
					message := other.C_embed("ERROR", "```on sait pas qui t'es```", config.Color_error)
					adresse_m := &message
					channel, _ := s.UserChannelCreate(m.Author.ID)
					_, _ = s.ChannelMessageSendEmbed(channel.ID, adresse_m)
				} else {
					message := other.C_embed("LISTE DES POINTS", to_print, config.Color_reponse)
					adresse_m := &message
					channel, _ := s.UserChannelCreate(m.Author.ID)
					_, _ = s.ChannelMessageSendEmbed(channel.ID, adresse_m)
				}

			}
			// presente le leader board des guildes.
		} else if StartWith(response[1], "guilde") {
			api := config.ReadApi()
			stud, err := data.StudDataGet()
			if err != nil {
				message := other.C_embed("ERROR", "```Probleme dand la lecture du fichier étudiant```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				return
			} else {
				allGuild, err := other.Find_Guild(api, stud)
				if err != nil {
					message := other.C_embed("ERROR", "```Probleme dans la recherche des guilde au sein de l'API```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					return
				} else {
					to_print := other.List_Guild(allGuild)
					message := other.C_embed("LISTE DES GUILDE", to_print, config.Color_reponse)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				}
			}
			// presente le leader board des étudiants.
		} else if StartWith(response[1], "liste") {
			if !admin {
				message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				api := config.ReadApi()
				stud, _ := data.StudDataGet()
				api, _ = other.Tri_Stud(api, stud)
				to_print := ""
				for i, v := range api {
					if i == len(stud)-1 {
						break
					}
					to_print += "``` n°: " + strconv.Itoa(i+1) + "\n" + (stud)[v.Id].Nom + " " + (stud)[v.Id].Prenom + " \n   -> ID : " + strconv.Itoa(v.Id) + " \n   -> guilde : " + api[v.Id].Guild + " \n   -> point : " + strconv.Itoa(v.Point) + "```\n\n"
					if i%10 == 9 {
						message := other.C_embed("LISTE DES POINTS", to_print, config.Color_reponse)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						to_print = ""
					}
				}
				message := other.C_embed("LISTE DES POINTS", to_print, config.Color_reponse)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			}
			// sinon, presente un message d'erreur.
		} else {
			message := other.C_embed("ERROR", "```Argument incorrecte pour la commande \"point\"```", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		}
	}

	// attribut des valeur precise a des cibles.
	if StartWith(m.Content, "set") {
		if !admin {
			message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			response := other.Split(m.Content)
			if len(response) != 4 {
				message := other.C_embed("ERROR", "```Argument incorrecte pour la commande \"set\"```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				// permet de changer la guilde de la cible.
				if StartWith(response[1], "guilde") {
					stud_list, err := data.StudDataGet()
					if err != nil {
						message := other.C_embed("ERROR", "```Probleme dand la lecture du fichier étudiant```", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						return
					}
					// cible tout le monde
					if response[2] == "all" {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							api := config.ReadApi()
							for i := range api {
								api[i].Guild = response[3]
							}

							file, err := json.Marshal(api)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)

							if err != nil {
								message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								message := other.C_embed("ATTRIBUTION DE GUILDE", "```Tout les Etudiants ont été update```", config.Color_reponse)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
							}
						}
						// sinon, cherche le nom dans la base de donnée.
					} else {
						id_stud, nope := other.Find_in_stud(response[2], stud_list, "nom")
						if !nope {
							message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							_, nope = gitmanage.Pull()
							if !nope {
								message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								api := config.ReadApi()
								api[id_stud].Guild = response[3]

								file, err := json.Marshal(api)
								ioutil.WriteFile("./ApiData/api.json", file, 0777)

								if err != nil {
									message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								} else {
									message := other.C_embed("ATTRIBUTION DE GUILDE", "```"+(stud_list)[id_stud].Prenom+" "+(stud_list)[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point: "+strconv.Itoa(api[id_stud].Point)+"\n   ->Nouvelle Guilde: "+api[id_stud].Guild+"```", config.Color_reponse)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
									// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
								}
							}
						}
					}
				}
				// permet de changer les points de la cible.
				if StartWith(response[1], "point") {
					stud_list, err := data.StudDataGet()

					if err != nil {
						message := other.C_embed("ERROR", "```Probleme dand la lecture du fichier étudiant```", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						return
					}
					if response[2] == "all" {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							api := config.ReadApi()
							for i := range api {
								api[i].Point, _ = strconv.Atoi(response[3])
							}

							file, err := json.Marshal(api)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)

							if err != nil {
								message := other.C_embed("ERROR", "l'ecriture dans le fichier api.json a échoué", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								message := other.C_embed("ATTRIBUTION DE POINT", "```Tout les Etudiants ont été update```", config.Color_reponse)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
							}
						}
					} else {
						id_stud, nope := other.Find_in_stud(response[2], stud_list, "nom")
						if !nope {
							message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							_, nope = gitmanage.Pull()
							if !nope {
								message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								api := config.ReadApi()
								api[id_stud].Point, _ = strconv.Atoi(response[3])

								file, err := json.Marshal(api)
								ioutil.WriteFile("./ApiData/api.json", file, 0777)

								if err != nil {
									message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								} else {
									message := other.C_embed("ATTRIBUTION DES POINTS", "```"+(stud_list)[id_stud].Prenom+" "+(stud_list)[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point total: "+strconv.Itoa(api[id_stud].Point)+"\n   ->Guilde: "+api[id_stud].Guild+"```", config.Color_reponse)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
									// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
								}
							}
						}
					}
				}
			}
		}
	}

	// leader board
	if StartWith(m.Content, "top") {
		response := other.Split(m.Content)
		if len(response) == 1 {
			api := config.ReadApi()
			stud, _ := data.StudDataGet()
			api, _ = other.Tri_Stud(api, stud)
			to_print := ""
			for i, v := range api[0:5] {
				if i == len(stud)-1 {
					break
				}
				to_print += "``` n°: " + strconv.Itoa(i+1) + "\n" + (stud)[v.Id].Nom + " " + (stud)[v.Id].Prenom + " \n   -> ID : " + strconv.Itoa(v.Id) + " \n   -> guilde : " + api[v.Id].Guild + " \n   -> point : " + strconv.Itoa(v.Point) + "```\n\n"
				if i%10 == 9 {
					message := other.C_embed("LISTE DES POINTS", to_print, config.Color_reponse)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					to_print = ""
				}
			}
			message := other.C_embed("LISTE DES POINTS", to_print, config.Color_reponse)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else if len(response) != 2 {
			message := other.C_embed("ERROR", "```Argument incorrecte pour la commande \"top\"```", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else if len(response) == 2 {
			if StartWith(response[1], "guilde") {
				api := config.ReadApi()
				stud, _ := data.StudDataGet()
				//api, _ = other.Tri_Stud(api, stud)
				allGuild, err := other.Find_Guild(api, stud)

				if err != nil {
					message := other.C_embed("ERROR", "```Probleme dans la recherche des guilde au sein de l'API```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					return
				} else {
					triGuild, _ := other.Tri_Guild(allGuild)
					to_string := ""
					for i, v := range triGuild {
						stud, _ = data.StudDataGet()
						//fmt.Printf("other.List_Membre(v): %v\n", other.List_Membre(v, *stud))
						tri, _ := other.Tri_Stud(triGuild[i].Membre, stud)
						to_string += "**" + v.Nom + "**" + "```" +
							"\nclassement: " + strconv.Itoa(i+1) + " | point: " + strconv.Itoa(triGuild[i].Point) +
							"\n   ->nombre de membre: " + strconv.Itoa(len(triGuild[i].Membre)) +
							"\n   ->top 3: " +
							"\n      1." + (stud)[(tri[0].Id)].Nom + " " + (stud)[(tri[0].Id)].Prenom +
							"\n         point: " + strconv.Itoa(tri[0].Point) +
							"\n      2." + (stud)[(tri[1].Id)].Nom + " " + (stud)[(tri[1].Id)].Prenom +
							"\n         point: " + strconv.Itoa(tri[1].Point) +
							"\n      3." + (stud)[(tri[2].Id)].Nom + " " + (stud)[(tri[2].Id)].Prenom +
							"\n         point: " + strconv.Itoa(tri[2].Point) +
							"```"

					}
					message := other.C_embed("LISTE DES GUILDE", to_string, config.Color_reponse)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				}
			}
		}
	}

	// commande aide pour visualisé les utilisation des autres commandes
	if StartWith(m.Content, "aide") {
		if admin {
			message := other.C_embed("GUILD-BOT\n   ->Les commandes",
				"``` aide\n   ->visualise les commandes du bot ```\n"+
					"``` point\n   ->visualise vos points ```\n"+
					"``` point guilde\n   ->visualise les points des guildes ```\n"+
					"``` top\n   ->visualise le top 3 des étudiants de chaque guilde \n       <<EN DEVELOPEMENT>> ```\n",
				config.Color_reponse)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			message = other.C_embed("GUILD-BOT\n\n----ADMIN----",
				"``` point liste\n   ->visualise les points de tout les étudiants ```\n"+
					"``` gain {nom_etudiant} {to_add}\n   ->ajoute un certain nombre de point a l'etudiant (nom de l'étudiant en majuscule) ```\n"+
					"``` set {guilde|point} {all|nom_etudiant {to_set}\n   ->change la guilde ou les points de l'étudiant pour la valeur preciser.\n   ->Si vous ecrivez \"all\" a la place de l'etudiant, cible tout les étudiants ```\n"+
					"``` push\n   ->sauvegarde sur git les données ```\n",
				config.Color_reponse)
			adresse_m = &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			message := other.C_embed("GUILD-BOT\n   ->Les commandes",
				"``` aide\n   ->visualise les commandes du bot ```\n"+
					"``` point\n   ->visualise vos points ```\n"+
					"``` point guilde\n   ->visualise les points des guildes ```\n"+
					"``` top\n   ->visualise le top 3 des étudiants de chaque guilde \n       <<EN DEVELOPEMENT>> ```\n",
				config.Color_reponse)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		}

	}

	// push les données api en ligne.
	if StartWith(m.Content, "push") {
		if !admin {
			message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			response := other.Split(m.Content)
			if len(response) <= 1 {
				message := other.C_embed("PUSHING", "```pushing ...```", config.Color_reponse)
				adresse_m := &message
				embed_message, _ := s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				if gitmanage.Commit() {
					message := other.C_embed("PUSHING", "```pushing ... succeed```", config.Color_reponse)
					adresse_m := &message
					s.ChannelMessageEditEmbed(m.ChannelID, embed_message.ID, adresse_m)
					// s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... succeed")
				} else {
					message := other.C_embed("PUSHING", "```pushing ... failed```", config.Color_error)
					adresse_m := &message
					s.ChannelMessageEditEmbed(m.ChannelID, embed_message.ID, adresse_m)
					// s.ChannelMessageEdit(m.ChannelID, commit.ID, "pushing ... failed")
				}
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
