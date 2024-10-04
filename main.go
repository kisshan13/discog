package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	commands "github.com/kisshan13/discog/command"
	discog "github.com/kisshan13/discog/package"
)

func main() {

	manager := commands.GetManager()
	bot, err := discog.NewBot("OTU0NzUyNTE5NjE4MzIyNTIz.GprFhd.Xl-zpxYhGBawoCeDcxXwy_7xvVoD7qiE4tu-dM", manager)

	if err != nil {
		panic(err)
	}

	bot.OnReady(func(r *discordgo.Ready) {
		log.Printf("Bot has connected to Discord as %s (%s)\n", r.User.Username, r.User.ID)
	})

	bot.Session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {
		if e.Author.ID != s.State.User.ID {
			guild, err := s.Guild(e.GuildID)

			if err != nil {
				println(err.Error())
			}

			println(guild.ID)

			s.ChannelMessageSend(e.ChannelID, "Kya re lund ke")
		}
	})

	bot.Run(func(session *discordgo.Session, err error) {
		if err != nil {
			panic(err)
		}
		log.Println("Bot is running")
	})
}