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

		actionRow := discog.NewActionRow()

		button := discog.NewButton().AddContent("", "Hello", "https://kisshan.xyz", false).AddStyle(discordgo.LinkButton)

		actionRow.AddComponent(button.GetComponent())

		ctx.Content("I can't kick you have to add functionality first")
		ctx.SetComponents([]discordgo.MessageComponent{actionRow.GetActionRow()})
		err := ctx.Send()

		if err != nil {
			println(err.Error())
		}
	})

	return command
}
