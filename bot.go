package discog

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot is the main struct for interacting with Discord. It holds the bot's token and command manager.
type Bot struct {
	token          string
	CommandManager *CommandManager
	Session        *discordgo.Session
}

// NewBot creates a new Bot with the provided token and command manager.
func NewBot(token string, manager *CommandManager) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	return &Bot{token: "Bot " + token, Session: session, CommandManager: manager}, nil
}

// Run starts the Discord bot with the provided callback function.
func (b *Bot) Run(callback func(session *discordgo.Session, err error)) {
	err := b.Session.Open()

	if err != nil {
		callback(nil, err)
		return
	}

	callback(b.Session, nil)

	b.CommandManager.register(b)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	b.CommandManager.unregister(b)
	log.Println("Graceful shutdown")
}

// OnReady is a callback function that is called when the bot is ready. It logs the bot's status and registers the interaction handlers.
func (b *Bot) OnReady(callback func(r *discordgo.Ready)) {
	b.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Ready! %+v\n", r)
		callback(r)
	})

	if len(b.CommandManager.groups) != 0 {
		b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			if i.Interaction.Type == discordgo.InteractionModalSubmit && b.CommandManager.modalHandler != nil {
				ctx := newCtx(s, i)
				b.CommandManager.modalHandler(ctx)
				return
			}

			if i.Interaction.Type == discordgo.InteractionMessageComponent && b.CommandManager.messageComponentHandler != nil {
				ctx := newCtx(s, i)
				b.CommandManager.messageComponentHandler(ctx)
				return
			}

			handler := b.CommandManager.FindCommand(i.ApplicationCommandData().Name)

			if handler != nil {
				ctx := newCtx(s, i)

				handler(ctx)
			}
		})
	}
}
