package other

import (
	"github.com/bwmarrin/discordgo"
)

func C_embed(title string, description string, color int) discordgo.MessageEmbed {
	var ret discordgo.MessageEmbed

	ret.Title = title
	ret.Description = description
	ret.Type = "rich"
	ret.Color = color

	return ret
}
