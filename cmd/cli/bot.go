package cli

import (
	"errors"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/hazzardr/go-template/internal/discord"
)

var botToken string

var errNoToken = errors.New("discord token required: pass --token or set $DISCORD_TOKEN")

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Discord bot actions",
}

var botServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Discord bot",
	RunE: func(_ *cobra.Command, _ []string) error {
		if botToken == "" {
			return errNoToken
		}
		b, err := discord.NewBot(botToken)
		if err != nil {
			return err
		}
		return b.Serve()
	},
}

var botSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync slash commands with Discord. Run once after adding a new command.",
	RunE: func(_ *cobra.Command, _ []string) error {
		if botToken == "" {
			return errNoToken
		}
		b, err := discord.NewBot(botToken)
		if err != nil {
			return err
		}
		if err := b.SyncCommands(); err != nil {
			return err
		}
		slog.Info("discord command sync successful")
		return nil
	},
}

func init() {
	botCmd.PersistentFlags().StringVarP(&botToken, "token", "t", os.Getenv("DISCORD_TOKEN"),
		"Discord bot token (defaults to $DISCORD_TOKEN)")
	botCmd.AddCommand(botServeCmd, botSyncCmd)
	rootCmd.AddCommand(botCmd)
}
