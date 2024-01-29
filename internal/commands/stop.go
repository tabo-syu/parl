package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
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
		IconURL: icon,
		Text:    "Pal Server",
	},
}

func stop() *discordgo.MessageEmbed {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := rcon.Dial(address, password, rcon.SetDeadline(1*time.Second+500*time.Millisecond))
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
		Title: "3分後にゲームサーバーは停止します！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: icon,
			Text:    status,
		},
	}
}
