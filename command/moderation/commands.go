package moderations

import (
	"github.com/bwmarrin/discordgo"
	discog "github.com/kisshan13/discog/package"
)

func GetKickCommand() *discog.SlashCommand {
	command := discog.NewSlashCommand(discordgo.ApplicationCommand{
		Name:        "kick",
		Description: "Command to kick a user",
	})

	command.SetHandler(func(ctx *discog.Ctx) {
		ctx.Content("I can't kick you have to add functionality first")
		ctx.Send()
	})

	return command
}
