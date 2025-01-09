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

var stopCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "stop",
	Description: "Palworld ゲームサーバーを停止します。",
}

var stopErrMessage = &discordgo.MessageEmbed{
	Color: internal.ColorRed(),
	Title: "ゲームサーバーの停止に失敗しました...",
	Footer: &discordgo.MessageEmbedFooter{
		IconURL: env.Icon,
		Text:    "Pal Server",
	},
}

func stop() *discordgo.MessageEmbed {
	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	conn, err := rcon.Dial(address, env.Password, rcon.SetDeadline(1*time.Second+500*time.Millisecond))
	if err != nil {
		log.Println("dial:", address, err)

		return stopErrMessage
	}
	defer conn.Close()

	status, err := conn.Execute("Info")
	if err != nil {
		log.Println("execute:", address, err)

		return stopErrMessage
	}

	save, err := conn.Execute("Save")
	if err != nil {
		log.Println("execute:", address, err)

		return stopErrMessage
	}
	log.Println(save)

	_, err = conn.Execute("Shutdown 60 This-server-will-be-shutdown-after-a-minute.")
	if err != nil {
		log.Println("execute:", address, err)

		return stopErrMessage
	}

	return &discordgo.MessageEmbed{
		Color: internal.ColorOrange(),
		Title: "1分後にゲームサーバーは停止します！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: env.Icon,
			Text:    status,
		},
	}
}
