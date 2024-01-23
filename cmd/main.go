package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/tabo-syu/parl/internal"
	"github.com/tabo-syu/parl/internal/commands"
)

const DISCORD_TOKEN = ""

func main() {
	discord, err := discordgo.New("Bot " + DISCORD_TOKEN)
	if err != nil {
		log.Fatalf("invalid discord token: %s", err)
	}

	bot := internal.NewBot(discord, commands.NewParl())
	if err := bot.Start(); err != nil {
		log.Fatalf("cannot open the session: %s", err)
	}
	defer bot.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("press Ctrl+C to exit")
	<-stop

	log.Println("gracefully shutting down.")
}
