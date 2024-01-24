package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
)

var statusCmd = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "status",
	Description: "Palworld ゲームサーバーの状態を確認します。",
}

func status() string {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := rcon.Dial(address, password)
	if err != nil {
		log.Println("dial: ", address, err)

		return "ゲームサーバーが停止しています。"
	}
	defer conn.Close()

	response, err := conn.Execute("Info")
	if err != nil {
		log.Println("execute: ", address, err)

		return "エラー"
	}

	return response
}
