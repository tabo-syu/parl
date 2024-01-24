package internal

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	sess *discordgo.Session
	cmd  RootCommand
}

func NewBot(sess *discordgo.Session, cmd RootCommand) *Bot {
	return &Bot{sess, cmd}
}

func (b *Bot) Start() error {
	// Bot がサーバーに参加したとき
	b.sess.AddHandler(func(d *discordgo.Session, g *discordgo.GuildCreate) {
		_, err := b.sess.ApplicationCommandCreate(b.sess.State.User.ID, g.Guild.ID, b.cmd.GetCommand())
		if err != nil {
			log.Printf("could not regiter a command at %s guild: %s\n", g.Guild.Name, err)
		} else {
			log.Printf("register a command at %s(%s) guild", g.Guild.Name, g.Guild.ID)
		}
	})

	// コマンドを受け付けたとき
	b.sess.AddHandler(func(d *discordgo.Session, i *discordgo.InteractionCreate) {
		request := i.ApplicationCommandData()
		if b.cmd.GetCommand().Name != request.Name {
			return
		}

		response := b.cmd.Handle(&request)

		d.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{response},
			},
		})
	})

	if err := b.sess.Open(); err != nil {
		return fmt.Errorf("could not open the session: %w", err)
	}

	return nil
}

func (h *Bot) Stop() error {
	if err := h.sess.Close(); err != nil {
		return fmt.Errorf("could not close the session: %w", err)
	}

	return nil
}
