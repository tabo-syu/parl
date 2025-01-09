package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tabo-syu/parl/env"
	"github.com/tabo-syu/parl/internal"
)

var invalidRequestErrMessage = &discordgo.MessageEmbed{
	Color: internal.ColorRed(),
	Title: "コマンドの入力値が不正です...",
	Footer: &discordgo.MessageEmbedFooter{
		IconURL: env.Icon,
		Text:    "Pal Server",
	},
}

type Parl struct {
	Command *discordgo.ApplicationCommand
	API     *internal.API
}

func NewParl() *Parl {
	return &Parl{
		Command: &discordgo.ApplicationCommand{
			Name:        "parl",
			Description: "Palworld ゲームサーバーの起動や停止、状態の確認ができます",
			Options: []*discordgo.ApplicationCommandOption{
				statusCmd,
				startCmd,
				stopCmd,
			},
		},
		API: internal.NewAPI(env.Password),
	}
}

func (p *Parl) GetCommand() *discordgo.ApplicationCommand {
	return p.Command
}

func (p *Parl) Handle(request *discordgo.ApplicationCommandInteractionData) *discordgo.MessageEmbed {
	if len(request.Options) == 0 {
		return invalidRequestErrMessage
	}

	subCmd := request.Options[0]
	switch subCmd.Name {
	case statusCmd.Name:
		return status(p.API)
	case startCmd.Name:
		return start(p.API)
	case stopCmd.Name:
		return stop(p.API)
	default:
		return invalidRequestErrMessage
	}
}
