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

		button := discog.NewButton().AddContent("", "Hello", "https://discord.gg", false).AddStyle(discordgo.LinkButton)

		// selectMenu := discog.NewSelectMenu(discordgo.StringSelectMenu)
		// selectMenu.PlaceHolder("Add information here...")
		// selectMenu.AddID("id")
		// selectMenu.AddOptions([]discog.SelectMenuOption{
		// 	*discog.NewSelectMenuOption("Hello", "why", "This is a optional"),
		// 	*discog.NewSelectMenuOption("Hii", "hui", "This is a optional"),
		// })

		// err := actionRow.AddComponent(selectMenu)
		err := actionRow.AddComponent(button)

		if err != nil {
			println(err.Error())
		}

		ctx.Content("I can't kick you have to add functionality first")
		_, err = ctx.SetComponents(actionRow.GetComponent())

		if err != nil {
			println(err.Error())
		}

		err = ctx.Send()

		if err != nil {
			println(err.Error())
		}

	})

	return command
}
