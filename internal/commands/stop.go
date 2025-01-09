package commands

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
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

func stop(api *internal.API) *discordgo.MessageEmbed {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	info, err := api.ServerInfo(ctx)
	if err != nil {
		log.Printf("failed to stop command during 1st info request: %s", err)

		return statusErrMessage
	}

	if err = api.ShutdownServer(ctx, 1, "1分後にゲームサーバーは停止します！"); err != nil {
		log.Printf("failed to stop command during 'shutdown server': %s", err)

		return stopErrMessage
	}

	return &discordgo.MessageEmbed{
		Color: internal.ColorOrange(),
		Title: "1分後にゲームサーバーは停止します！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: env.Icon,
			Text:    fmt.Sprintf("%s [%s] ", info.ServerName, info.Version),
		},
	}
}
