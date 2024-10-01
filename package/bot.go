package discog

import "github.com/bwmarrin/discordgo"

type Bot struct {
	token   string
	Session *discordgo.Session
}

func NewBot(token string) *Bot {
	return &Bot{token: "Bot " + token}
}

func (b *Bot) Run(callback func(session *discordgo.Session, err error)) {
	session, err := discordgo.New(b.token)

	if err != nil {
		callback(nil, err)
		return
	}

	b.Session = session

	err = b.Session.Open()

	if err != nil {
		callback(nil, err)
		return
	}

	callback(session, nil)
}

func (b *Bot) OnReady(callback func(r *discordgo.Ready)) {
	b.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		callback(r)
	})
}
