package commands

import "github.com/bwmarrin/discordgo"

var stopCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "stop",
	Description: "Palworld ゲームサーバーを停止します。",
}

func stop() *discordgo.MessageEmbed {
	return nil
}
