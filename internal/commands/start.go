package commands

import "github.com/bwmarrin/discordgo"

var startCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "start",
	Description: "Palworld ゲームサーバーを起動します。",
}

func start() string {
	return "start"
}
