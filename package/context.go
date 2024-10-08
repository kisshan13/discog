package discog

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Represents callback ctx for Slash Command Handlers
type Ctx struct {
	Session     *discordgo.Session           // Client Session
	Interaction *discordgo.InteractionCreate // Interaction Create struct

	interactionResponse *discordgo.InteractionResponse
}

// Creates a new Ctx for the given Discord session and interaction.
func newCtx(s *discordgo.Session, i *discordgo.InteractionCreate) *Ctx {

	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{},
	}

	return &Ctx{Session: s, Interaction: i, interactionResponse: interactionResponse}
}

// Sends the response to Discord.
func (c *Ctx) Send() error {
	return c.Session.InteractionRespond(c.Interaction.Interaction, c.interactionResponse)
}

// Sets the title for the response.
func (c *Ctx) SetTitle(title string) *Ctx {
	c.interactionResponse.Data.Title = title
	return c
}

// Sets the Interaction response type .
func (c *Ctx) SetResponseType(responseType discordgo.InteractionResponseType) *Ctx {
	c.interactionResponse.Type = responseType
	return c
}

// Adds a customId to the interaction response
func (c *Ctx) SetCustomID(id string) *Ctx {
	c.interactionResponse.Data.CustomID = id
	return c
}

// Sets the content of the interaction response.
func (c *Ctx) SetContent(content string) *Ctx {
	c.interactionResponse.Data.Content = content
	return c
}

// Sets the Text-To-Speech as true
func (c *Ctx) SetTTS() *Ctx {
	c.interactionResponse.Data.TTS = true
	return c
}

// Adds Files to the interaction response.
func (c *Ctx) SetFiles(files []*discordgo.File) *Ctx {
	c.interactionResponse.Data.Files = append(c.interactionResponse.Data.Files, files...)
	return c
}

// Flags the interaction response as ephemeral.
func (c *Ctx) FlagEphemeral() *Ctx {
	c.interactionResponse.Data.Flags = discordgo.MessageFlagsEphemeral
	return c
}

// Adds message flags to the response.
func (c *Ctx) FlagReply(flag discordgo.MessageFlags) *Ctx {
	c.interactionResponse.Data.Flags = flag
	return c
}

// Adds the components to the interaction response.
func (c *Ctx) SetComponents(actionRows []Component) (*Ctx, error) {
	c.interactionResponse.Data.Components = []discordgo.MessageComponent{}

	for _, row := range actionRows {
		messageComponent, ok := row.GetComponent().(discordgo.ActionsRow)

		if !ok {
			return nil, errors.New("Must be ActionRow")
		}

		c.interactionResponse.Data.Components = append(c.interactionResponse.Data.Components, messageComponent)
	}
	return c, nil
}

// Sets message embeds to the interaction response.
func (c *Ctx) SetMessageEmbeds(embeds []*discordgo.MessageEmbed) {
	c.interactionResponse.Data.Embeds = embeds
}
