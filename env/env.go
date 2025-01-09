package env

import "os"

var (
	// Palworld
	Host       = os.Getenv("API_HOST")
	Port       = os.Getenv("API_PORT")
	Password   = os.Getenv("API_PASSWORD")
	ServerPath = os.Getenv("SERVER_PATH")
	Icon       = os.Getenv("DISCORD_ICON_URL")

	// Discord
	DiscordToken = os.Getenv("DISCORD_TOKEN")
)
