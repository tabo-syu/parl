package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
)

var statusCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "status",
	Description: "Palworld ゲームサーバーの状態を確認します。",
}

func status() string {
	host := os.Getenv("RCON_HOST")
	port := os.Getenv("RCON_PORT")
	password := os.Getenv("RCON_PASSWORD")

	endpoint := fmt.Sprintf("%s:%s", host, port)
	conn, err := rcon.Dial(endpoint, password)
	if err != nil {
		log.Println("dial: ", endpoint, err)

		return "ゲームサーバーが停止しています。"
	}
	defer conn.Close()

	response, err := conn.Execute("Info")
	if err != nil {
		log.Println("execute: ", endpoint, err)

		return "エラー"
	}

	return response
}
