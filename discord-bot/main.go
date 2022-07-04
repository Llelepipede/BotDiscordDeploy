package main

import (
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
    "log"
)
func main() {
    err := config.ReadConfig()
    if err != nil {
        log.Fatal(err)
        return
    }
    bot.Run()
    <-make(chan struct{})
    return
}