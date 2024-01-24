package commands

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/tabo-syu/parl/internal"
)

var (
	host     = os.Getenv("RCON_HOST")
	port     = os.Getenv("RCON_PORT")
	password = os.Getenv("RCON_PASSWORD")
)

type Parl struct {
	Command *discordgo.ApplicationCommand
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
				restartCmd,
			},
		},
	}
}

func (p *Parl) GetCommand() *discordgo.ApplicationCommand {
	return p.Command
}

func (p *Parl) Handle(request *discordgo.ApplicationCommandInteractionData) *discordgo.MessageEmbed {
	if len(request.Options) == 0 {
		return &discordgo.MessageEmbed{
			Color: internal.Color("ff0000"),
			Title: "ゲームサーバーは停止しています...",
		}
	}

	subCmd := request.Options[0]
	switch subCmd.Name {
	case statusCmd.Name:
		return status()
	case startCmd.Name:
		return start()
	case stopCmd.Name:
		return stop()
	case restartCmd.Name:
		return restart()
	default:
		return &discordgo.MessageEmbed{
			Color: internal.Color("ff0000"),
			Title: "ゲームサーバーは停止しています...",
		}
	}
}
