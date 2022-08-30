package main

import (
	"golang-discord-bot/bot"
	"golang-discord-bot/config"

	//"golang-discord-bot/data"
	//"golang-discord-bot/gitmanage"
	"log"
)

// lien d'invitation du bot dans le serveur
// https://discord.com/oauth2/authorize?client_id=993483152762876014&scope=bot

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		bot.Run()
		<-make(chan struct{})

		return
	}

}

// func main() {

// 	fmt.Printf("bot.StartWith(\"/prout        \", \"/prout \"): %v\n", bot.StartWith("/prout        ", "/prout "))
// }
