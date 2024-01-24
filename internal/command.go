package internal

import "github.com/bwmarrin/discordgo"

type RootCommand interface {
	GetCommand() *discordgo.ApplicationCommand
	Handle(*discordgo.ApplicationCommandInteractionData) *discordgo.MessageEmbed
}
