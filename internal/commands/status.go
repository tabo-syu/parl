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

func status(api *internal.API) *discordgo.MessageEmbed {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	info, err := api.ServerInfo(ctx)
	if err != nil {
		log.Printf("failed to status command: %s\n", err)

		return statusErrMessage
	}

	return &discordgo.MessageEmbed{
		Color: internal.ColorGreen(),
		Title: "ゲームサーバーは稼働中です！",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: env.Icon,
			Text:    fmt.Sprintf("%s [%s] ", info.ServerName, info.Version),
		},
	}
}
