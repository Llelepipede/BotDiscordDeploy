package main

import (
	"fmt"
	"golang-discord-bot/bot"
	"golang-discord-bot/config"

	//"golang-discord-bot/data"
	//"golang-discord-bot/gitmanage"
	"log"

	"github.com/go-git/go-git/storage"
)

// lien d'invitation du bot dans le serveur
// https://discord.com/oauth2/authorize?client_id=993483152762876014&scope=bot

func main() {
	var i storage.Storer
	fmt.Printf("storage.Storer: %v\n", i)
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
