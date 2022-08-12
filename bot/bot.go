package bot

import (
	"encoding/json"
	"fmt"
	"golang-discord-bot/config"
	"golang-discord-bot/data"
	"golang-discord-bot/dataStruct"
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
	config.Aide()

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
	// add -t {type} -e {etudiant (nom prenom | prenom nom)} -v {valeur}
	if StartWith(m.Content, "add") {
		if !admin {
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			splited := other.Splitdot(m.Content, " -")
			if len(splited) != 4 {
				message := other.C_embed("ERROR", "```argument incorrecte pour la commande \"add\"\n\n      ->utilisation : add -t {type} -e {etudiant (nom prenom | prenom nom)} -v {valeur}```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				_, nope := gitmanage.Pull()
				if !nope {
					message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				} else {
					all_stud := config.ReadApi()

					all_stud, indexEtud, _ := other.Add(splited[1:], all_stud)
					if indexEtud == -1 {
						message := other.C_embed("ERROR", "```erreur de la fonction add```", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					} else {
						file, err := json.Marshal(all_stud)
						ioutil.WriteFile("./ApiData/api.json", file, 0777)

						if err != nil {
							message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							message := other.C_embed("GAIN DE POINT", "```"+all_stud[indexEtud].Prenom+" "+all_stud[indexEtud].Nom+" \n   ->ID: "+strconv.Itoa(indexEtud)+"\n   ->Point total: "+strconv.Itoa(all_stud[indexEtud].Point)+"\n   ->Guilde: "+all_stud[indexEtud].Guild+"\n   ->Credit: "+strconv.Itoa(all_stud[indexEtud].Credit)+"```", config.Color_reponse)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							gitmanage.Commit(m.Content)
						}
					}

				}

			}
		}
	}

	if StartWith(m.Content, "remove") {
		if !admin {
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			splited := other.Splitdot(m.Content, " -")
			if len(splited) != 4 {
				message := other.C_embed("ERROR", "```argument incorrecte pour la commande \"remove\"\n\n      ->utilisation : remove -t {type} -e {etudiant (nom prenom , prenom nom)} -v {valeur}```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				_, nope := gitmanage.Pull()
				if !nope {
					message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				} else {
					all_stud := config.ReadApi()

					all_stud, indexEtud, _ := other.Remove(splited[1:], all_stud)
					if indexEtud == -1 {
						message := other.C_embed("ERROR", "```erreur de la fonction remove```", config.Color_error)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					} else {
						file, err := json.Marshal(all_stud)
						ioutil.WriteFile("./ApiData/api.json", file, 0777)

						if err != nil {
							message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							message := other.C_embed("RETRAIT DE POINT", "```"+all_stud[indexEtud].Prenom+" "+all_stud[indexEtud].Nom+" \n   ->ID: "+strconv.Itoa(indexEtud)+"\n   ->Point total: "+strconv.Itoa(all_stud[indexEtud].Point)+"\n   ->Guilde: "+all_stud[indexEtud].Guild+"\n   ->Credit: "+strconv.Itoa(all_stud[indexEtud].Credit)+"```", config.Color_reponse)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							gitmanage.Commit(m.Content)
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
				all_stud := config.ReadApi()
				to_print := ""
				for _, v := range all_stud {
					if v.Discord == m.Author.ID {
						to_print += "```" + v.Nom + " " + v.Prenom + " \n   -> guilde : " + v.Guild + " \n   -> point : " + strconv.Itoa(v.Point) + " \n   -> Crédit : " + strconv.Itoa(v.Credit) + "```\n\n"
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
			all_stud := config.ReadApi()

			allGuild, err := other.Find_Guild(all_stud)
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

			// presente le leader board des étudiants.
		} else if StartWith(response[1], "liste") {
			if !admin {
				// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
				// adresse_m := &message
				// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				all_stud := config.ReadApi()
				all_stud, _ = other.Tri_Stud(all_stud)
				to_print := ""
				for i, v := range all_stud {
					to_print += "``` n°: " + strconv.Itoa(i+1) + "\n" + v.Nom + " " + v.Prenom + " \n   -> ID : " + strconv.Itoa(v.Id) + " \n   -> guilde : " + v.Guild + " \n   -> point : " + strconv.Itoa(v.Point) + "```\n\n"
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
				if StartWith(response[1], "guild") {
					// cible tout le monde
					if response[2] == "all" {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud := config.ReadApi()
							for i := range all_stud {
								all_stud[i].Guild = response[3]
							}

							file, err := json.Marshal(all_stud)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)

							if err != nil {
								message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								message := other.C_embed("ATTRIBUTION DE GUILDE", "```Tout les Etudiants ont été update```", config.Color_reponse)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								gitmanage.Commit(m.Content)
								// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
							}
						}
						// sinon, cherche le nom dans la base de donnée.
					} else {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						}
						all_stud := config.ReadApi()
						id_stud, nope := other.Find_in_stud(response[2], all_stud, "nom")
						if !nope {
							message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud[id_stud].Guild = response[3]

							file, err := json.Marshal(all_stud)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)

							if err != nil {
								message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								message := other.C_embed("ATTRIBUTION DE GUILDE", "```"+all_stud[id_stud].Prenom+" "+all_stud[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point: "+strconv.Itoa(all_stud[id_stud].Point)+"\n   ->Nouvelle Guilde: "+all_stud[id_stud].Guild+"```", config.Color_reponse)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								gitmanage.Commit(m.Content)
								// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))

							}
						}
					}
				}
				// permet de changer les points de la cible.
				if StartWith(response[1], "point") {
					if response[2] == "all" {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud := config.ReadApi()
							for i := range all_stud {
								all_stud[i].Point, _ = strconv.Atoi(response[3])
							}

							file, err := json.Marshal(all_stud)
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
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud := config.ReadApi()
							id_stud, nope := other.Find_in_stud(response[2], all_stud, "nom")
							if !nope {
								message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								all_stud[id_stud].Point, _ = strconv.Atoi(response[3])

								file, err := json.Marshal(all_stud)
								ioutil.WriteFile("./ApiData/api.json", file, 0777)

								if err != nil {
									message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								} else {
									message := other.C_embed("ATTRIBUTION DES POINTS", "```"+all_stud[id_stud].Prenom+" "+(all_stud)[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point total: "+strconv.Itoa(all_stud[id_stud].Point)+"\n   ->Guilde: "+all_stud[id_stud].Guild+"```", config.Color_reponse)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
									// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
								}

							}
						}

					}
				}
				if StartWith(response[1], "credit") {
					if response[2] == "all" {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud := config.ReadApi()
							for i := range all_stud {
								all_stud[i].Credit, _ = strconv.Atoi(response[3])
							}

							file, err := json.Marshal(all_stud)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)

							if err != nil {
								message := other.C_embed("ERROR", "l'ecriture dans le fichier api.json a échoué", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								message := other.C_embed("ATTRIBUTION DE CREDIT", "```Tout les Etudiants ont été update```", config.Color_reponse)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								// _, _ = s.ChannelMessageSend(m.ChannelID, "l'id de "+(*stud_list)[id_stud].Prenom+" "+(*stud_list)[id_stud].Nom+" est : "+strconv.Itoa(id_stud))
							}
						}
					} else {
						_, nope := gitmanage.Pull()
						if !nope {
							message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							all_stud := config.ReadApi()
							id_stud, nope := other.Find_in_stud(response[2], all_stud, "nom")
							if !nope {
								message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								all_stud[id_stud].Credit, _ = strconv.Atoi(response[3])

								file, err := json.Marshal(all_stud)
								ioutil.WriteFile("./ApiData/api.json", file, 0777)

								if err != nil {
									message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								} else {
									message := other.C_embed("ATTRIBUTION DES POINTS", "```"+all_stud[id_stud].Prenom+" "+(all_stud)[id_stud].Nom+" \n   ->ID: "+strconv.Itoa(id_stud)+"\n   ->Point total: "+strconv.Itoa(all_stud[id_stud].Point)+"\n   ->Guilde: "+all_stud[id_stud].Guild+"```", config.Color_reponse)
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

	if StartWith(m.Content, "initLogs") {
		// if false {
		if !admin {
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			response := other.Split(m.Content)
			if len(response) <= 0 {
				// message := other.C_embed("ERROR", "```Argument incorrecte pour la commande \"initLogs\"```", config.Color_error)
				// adresse_m := &message
				// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				if response[1] == "&!&" {
					data.AddLogsALL(config.ReadApi())
				}
			}
		}
	}

	// leader board
	if StartWith(m.Content, "top") {
		response := other.Split(m.Content)
		if len(response) == 1 {
			all_stud := config.ReadApi()
			all_stud, _ = other.Tri_Stud(all_stud)
			to_print := ""
			for i, v := range all_stud[0:5] {
				if i == len(all_stud)-1 {
					break
				}
				to_print += "``` n°: " + strconv.Itoa(i+1) + "\n" + v.Nom + " " + v.Prenom + " \n   -> ID : " + strconv.Itoa(v.Id) + " \n   -> guilde : " + v.Guild + " \n   -> point : " + strconv.Itoa(v.Point) + " \n   -> Crédit : " + strconv.Itoa(v.Credit) + "```\n\n"
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
				all_stud := config.ReadApi()
				allGuild, err := other.Find_Guild(all_stud)

				if err != nil {
					message := other.C_embed("ERROR", "```Probleme dans la recherche des guilde au sein de l'API```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
					return
				} else {
					triGuild, _ := other.Tri_Guild(allGuild)
					to_string := ""
					for i, v := range triGuild {
						//fmt.Printf("other.List_Membre(v): %v\n", other.List_Membre(v, *stud))
						tri, _ := other.Tri_Stud(triGuild[i].Membre)
						to_string += "**" + v.Nom + "**" + "```" +
							"\nclassement: " + strconv.Itoa(i+1) + " | point: " + strconv.Itoa(triGuild[i].Point) +
							"\n   ->nombre de membre: " + strconv.Itoa(len(triGuild[i].Membre)) +
							"\n   ->top 3: " +
							"\n      1." + (all_stud)[(tri[0].Id)].Nom + " " + (all_stud)[(tri[0].Id)].Prenom +
							"\n         point: " + strconv.Itoa(tri[0].Point) +
							"\n      2." + (all_stud)[(tri[1].Id)].Nom + " " + (all_stud)[(tri[1].Id)].Prenom +
							"\n         point: " + strconv.Itoa(tri[1].Point) +
							"\n      3." + (all_stud)[(tri[2].Id)].Nom + " " + (all_stud)[(tri[2].Id)].Prenom +
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
			message := other.C_embed("GUILD-BOT\n   ->Les commandes :eyes: ",
				"``` aide\n   ->visualise les commandes du bot ```\n"+
					"``` point\n   ->visualise vos points ```\n"+
					"``` point guilde\n   ->visualise les points des guildes ```\n"+
					"``` top\n   ->visualise le top 5 des étudiants de chaque guilde ```\n",
				config.Color_reponse)
			adresse_m := &message
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			message = other.C_embed("GUILD-BOT\n\n----ADMIN----",
				"``` point liste\n\t\t-> Visualise les points de tout les étudiants ```\n"+
					config.Aide_add+config.Aide_remove+config.Aide_report+
					"\n``` set\n   ->change la guilde, les points, ou les crédits de l'étudiant pour la valeur precisé.\n"+
					"\t{nom | all}\t-> Precise l'etudiant qui doit avoir des valeur modifié. si all est ecrit , cela siblera toute la base de donnée. (1 seul element)\n"+
					"\t{point | guild | credit}\t-> Precise la valeur a changer. (1 seul element)\n"+
					"\t{valeur}\t-> Valeur a affecté a/aux étudiant/s cible. (1 seul element)```\n"+
					"``` push\n\t> Sauvegarde sur git les données ```\n",
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
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			response := other.Split(m.Content)
			if len(response) <= 1 {
				message := other.C_embed("PUSHING", "```pushing ...```", config.Color_reponse)
				adresse_m := &message
				embed_message, _ := s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				if gitmanage.Push(m.Content) {
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

	if StartWith(m.Content, "createAPI") {
		dataStruct.CreateApi()
	}

	if StartWith(m.Content, "get") {
		if !admin {
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			all_stud := config.ReadApi()
			get := other.Get(m.Content, all_stud)
			if get == nil {
				message := other.C_embed("RESULTAT DE LA REQUETE", "```Requète erronée```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				to_print := ""
				for i, v := range get {
					to_print += "```" + v + "```"
					if i%10 == 9 && i != len(get)-1 {
						message := other.C_embed("RESULTAT DE LA REQUETE", to_print, config.Color_reponse)
						adresse_m := &message
						_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						to_print = ""
					}
				}
				message := other.C_embed("RESULTAT DE LA REQUETE", to_print, config.Color_reponse)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			}
		}
	}

	if StartWith(m.Content, "report") {
		if !admin {
			// message := other.C_embed("ERROR", "```Vous n'etes pas dans le bon salon pour faire cette commande```", config.Color_error)
			// adresse_m := &message
			// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
		} else {
			splited := other.Splitdot(m.Content, " -")
			if len(splited) <= 4 {
				message := other.C_embed("ERROR", "```argument incorrecte pour la commande \"remove\"\n\n      ->utilisation : remove -t {type} -e {etudiant (nom prenom , prenom nom)} -v {valeur}```", config.Color_error)
				adresse_m := &message
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
			} else {
				_, nope := gitmanage.Pull()
				if !nope {
					message := other.C_embed("ERROR", "```Le Pull a echoué```", config.Color_error)
					adresse_m := &message
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
				} else {
					all_stud := config.ReadApi()
					all_logs := config.ReadLogs()
					all_logs, all_stud, indexEtud := data.AddLog(m.Content, all_stud, all_logs)

					if all_logs != nil && all_stud != nil {
						fmt.Printf("all_logs[indexEtud].Log: %v\n", all_logs[indexEtud].Log)
						if indexEtud == -1 {
							message := other.C_embed("ERROR", "```erreur de la fonction```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							file, err := json.Marshal(all_stud)
							ioutil.WriteFile("./ApiData/api.json", file, 0777)
							if err != nil {
								message := other.C_embed("ERROR", "```l'ecriture dans le fichier api.json a échoué```", config.Color_error)
								adresse_m := &message
								_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
							} else {
								file, err := json.Marshal(all_logs)
								ioutil.WriteFile("./ApiData/logsGeneral.json", file, 0777)

								if err != nil {
									message := other.C_embed("ERROR", "```l'ecriture dans le fichier logsGeneral.json a échoué```", config.Color_error)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
								} else {
									message := other.C_embed("ECRITURE DANS LES LOGS", "```\nétudiant affecté par les logs :\n\n"+all_stud[indexEtud].Prenom+" "+all_stud[indexEtud].Nom+" \n   ->ID: "+strconv.Itoa(indexEtud)+"\n   ->Point total: "+strconv.Itoa(all_stud[indexEtud].Point)+"\n   ->Guilde: "+all_stud[indexEtud].Guild+"\n   ->Credit: "+strconv.Itoa(all_stud[indexEtud].Credit)+"```", config.Color_reponse)
									adresse_m := &message
									_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
									gitmanage.Commit(m.Content)
								}
							}
						}
					} else {
						if indexEtud == -1 {
							message := other.C_embed("ERROR", "```étudiant introuvable```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						} else {
							message := other.C_embed("ERROR", "```erreur dans la commande report\nexemple d'utilisation:\n\t-> $:report -e Gomis kwency -t credit -c remove -d c'est absenté du cours pendant 3h -m Paul -v 2```", config.Color_error)
							adresse_m := &message
							_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
						}

					}

				}
			}
		}

	}
	// if StartWith(m.Content, "test") {
	// 	fmt.Printf("other.FormDisplay(config.ReadApi()): %v\n", other.FormDisplay(config.ReadApi()))
	// }
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