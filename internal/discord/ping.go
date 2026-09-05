package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var pingCommand = SlashCommand{
	Metadata: discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Pings the bot server",
	},
	Path:       "/ping",
	HandleFunc: handlePing,
}

func handlePing(_ discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	msg := discord.NewMessageCreate()
	msg.Content = "Pong!"
	return e.Respond(discord.InteractionResponseTypeCreateMessage, msg)
}
