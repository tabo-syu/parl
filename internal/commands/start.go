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

var startCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "start",
	Description: "Palworld ゲームサーバーを起動します。",
}

var startErrMessage = &discordgo.MessageEmbed{
	Color: internal.ColorRed(),
	Title: "ゲームサーバーの起動に失敗しました...",
	Footer: &discordgo.MessageEmbedFooter{
		IconURL: env.Icon,
		Text:    "Pal Server",
	},
}

var startAlreadyMessage = &discordgo.MessageEmbed{
	Color: internal.ColorOrange(),
	Title: "ゲームサーバーはすでに起動しています...",
	Footer: &discordgo.MessageEmbedFooter{
		IconURL: env.Icon,
		Text:    "Pal Server",
	},
}

func start() *discordgo.MessageEmbed {
	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	conn, err := rcon.Dial(address, env.Password, rcon.SetDeadline(200*time.Millisecond))
	if err == nil {
		conn.Close()

		return startAlreadyMessage
	}

	if err := internal.Nohup("/bin/bash", env.ServerPath); err != nil {
		log.Println("nohup:", err)

		return startErrMessage
	}

	return &discordgo.MessageEmbed{
		Color: internal.ColorGreen(),
		Title: "ゲームサーバーを起動しました！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: env.Icon,
			Text:    "Pal Server",
		},
	}
}
