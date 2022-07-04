package bot

import (
    "golang-discord-bot/config"
    "log"
	"fmt"
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
    goBot.AddHandler(OnReady)
    err = goBot.Open()

    if err != nil {
        return
    }
	fmt.Println("Bot is running !")
}

// OnReady is called whenaconnection to Discord is first established
func OnReady(s *discordgo.Session,  m *discordgo.Ready) {

	// Set the playing status.
	s.UpdateStatus(0, "@ me!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

// Ignore all messages created by the bot itself
    if m.Author.ID == BotID {
        return
    }
    
// If the message is "Hi" reply with "Hi Back!!"
    if m.Content == "Hi" || m.Content == "hi" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")
    }

	if m.Content == "Guild" || m.Content == "guild" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Your Guild is Horde")
    }

	if m.Content == "TopTop" || m.Content == "toptop" || m.Content == "Toptop" || m.Content == "topTop" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Vasy monte sur ma TopTop et je t'emmerai au bout du monde")
    }

	if m.Content == "Code" || m.Content == "code" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Le langage TopTop est un nouveau langage qui va revolutionner le monde de l'informatique. Avec le TopTop, vous serez au top du top !")
    }
}