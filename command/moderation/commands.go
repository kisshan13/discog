package moderations

import (
	"log"

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

		emailInput := discog.NewTextInput("email-id").SetLabel("Email").SetPlaceholder("kisshan@outlook.com").SetStyle(discordgo.TextInputShort)

		errorr := actionRow.AddComponent(emailInput)

		if errorr != nil {
			log.Println(errorr.Error())
		}

		ctx.SetContent("I can't kick you have to add functionality first")
		_, err := ctx.SetComponents([]discog.Component{actionRow})

		if err != nil {
			println(err.Error())
		}

		ctx.SetTitle("Submit the data ")
		ctx.SetCustomID("this-is-a-custom-id")
		ctx.SetResponseType(discordgo.InteractionResponseModal)

		err = ctx.Send()

		if err != nil {
			println(err.Error())
		}
	})

	return command
}
