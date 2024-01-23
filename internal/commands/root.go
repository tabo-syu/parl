package commands

import "github.com/bwmarrin/discordgo"

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

func (p *Parl) Handle(request *discordgo.ApplicationCommandInteractionData) string {
	if len(request.Options) == 0 {
		return "入力値が不正です！"
	}

	subCmd := request.Options[0]
	response := ""
	switch subCmd.Name {
	case statusCmd.Name:
		response = status()
	case startCmd.Name:
		response = start()
	case stopCmd.Name:
		response = stop()
	case restartCmd.Name:
		response = restart()
	default:
		response = "入力値が不正です！"
	}

	return response
}