package discord

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/disgoorg/snowflake/v2"
)

// SlashCommand wires an application command to its handler.
type SlashCommand struct {
	Metadata   discord.ApplicationCommandCreate
	Path       string
	HandleFunc handler.SlashCommandHandler
}

var commands = []SlashCommand{
	pingCommand,
}

const closeTimeout = 5 * time.Second

type Bot struct {
	client *bot.Client
}

func NewBot(token string) (*Bot, error) {
	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
				gateway.IntentGuilds,
				gateway.IntentDirectMessages,
			),
		),
		bot.WithEventListeners(mux()),
	)
	if err != nil {
		return nil, err
	}
	return &Bot{client: client}, nil
}

// Serve opens the gateway connection and blocks until the process is
// interrupted, then closes the connection.
func (b *Bot) Serve() error {
	if err := b.client.OpenGateway(context.Background()); err != nil {
		return err
	}
	slog.Info("discord connection successful")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), closeTimeout)
	defer cancel()
	b.client.Close(ctx)
	return nil
}

// SyncCommands registers the slash commands with Discord.
// Run once after adding a new command.
func (b *Bot) SyncCommands() error {
	creates := make([]discord.ApplicationCommandCreate, len(commands))
	for i, cmd := range commands {
		creates[i] = cmd.Metadata
	}
	return handler.SyncCommands(b.client, creates, make([]snowflake.ID, 0))
}

func mux() *handler.Mux {
	r := handler.New()
	r.Use(middleware.Logger)
	for _, cmd := range commands {
		r.SlashCommand(cmd.Path, cmd.HandleFunc)
	}
	return r
}
