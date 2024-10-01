package discog

import "github.com/bwmarrin/discordgo"

type Ctx struct {
	session     *discordgo.Session
	interaction *discordgo.Interaction
}

func NewCtx(s *discordgo.Session, i *discordgo.Interaction) *Ctx {
	return &Ctx{session: s, interaction: i}
}

func (c *Ctx) Send(content string) {
	c.session.InteractionRespond(c.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content, Flags: discordgo.MessageFlagsEphemeral},
	})
}
