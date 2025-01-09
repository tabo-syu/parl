package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
	"github.com/tabo-syu/parl/env"
	"github.com/tabo-syu/parl/internal"
)

var statusCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "status",
	Description: "Palworld ゲームサーバーの状態を確認します。",
}

var statusErrMessage = &discordgo.MessageEmbed{
	Color: internal.ColorRed(),
	Title: "ゲームサーバーは停止しています...",
	Footer: &discordgo.MessageEmbedFooter{
		IconURL: env.Icon,
		Text:    "Pal Server",
	},
}

func status() *discordgo.MessageEmbed {
	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	conn, err := rcon.Dial(address, env.Password, rcon.SetDeadline(500*time.Millisecond))
	if err != nil {
		log.Println("dial:", address, err)

		return statusErrMessage
	}
	defer conn.Close()

	response, err := conn.Execute("Info")
	if err != nil {
		log.Println("execute:", address, err)

		return statusErrMessage
	}

	return &discordgo.MessageEmbed{
		Color: internal.ColorGreen(),
		Title: "ゲームサーバーは稼働中です！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: env.Icon,
			Text:    response,
		},
	}
}
