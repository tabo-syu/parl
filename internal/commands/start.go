package commands

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
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

func start(api *internal.API) *discordgo.MessageEmbed {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	if _, err := api.ServerInfo(ctx); err == nil {
		log.Printf("failed to start command already started: %s\n", err)

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
