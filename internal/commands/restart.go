package commands

import "github.com/bwmarrin/discordgo"

var restartCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "restart",
	Description: "Palworld ゲームサーバーを再起動します。",
}

func restart() *discordgo.MessageEmbed {
	return nil
}
