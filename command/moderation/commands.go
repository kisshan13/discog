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

		embed := discog.NewMessageEmbed()

		embed.SetTitle("This is a title").SetDescription("A cool description here")

		ctx.SetMessageEmbeds([]*discordgo.MessageEmbed{
			embed.GetComponent(),
		})

		ctx.SetResponseType(discordgo.InteractionResponseChannelMessageWithSource)

		err := ctx.Send()

		if err != nil {
			println(err.Error())
		}
	})

	return command
}
