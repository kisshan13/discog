package discog

import "github.com/bwmarrin/discordgo"

// Represents a Slash Command for the bot
type SlashCommand struct {
	Command *discordgo.ApplicationCommand // Slash Command information
	Handler func(ctx *Ctx)                // Handler function to handle any interaction
}

// Creates a new Slash Command
func NewSlashCommand(command discordgo.ApplicationCommand) *SlashCommand {
	return &SlashCommand{Command: &command, Handler: nil}
}

// Sets handler for the command
func (c *SlashCommand) SetHandler(handler func(ctx *Ctx)) {
	c.Handler = handler
}
