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

// Sets the Interaction response type .
func (c *Ctx) ResponseType(responseType discordgo.InteractionResponseType) *Ctx {
	c.interactionResponse.Type = responseType
	return c
}

// Adds a customId to the interaction response
func (c *Ctx) SetCustomID(id string) *Ctx {
	c.interactionResponse.Data.CustomID = id
	return c
}

// Sets the content of the interaction response.
func (c *Ctx) Content(content string) *Ctx {
	c.interactionResponse.Data.Content = content
	return c
}

// Sets the Text-To-Speech as true
func (c *Ctx) TTS() *Ctx {
	c.interactionResponse.Data.TTS = true
	return c
}

// Adds Message Components to the interaction response.
func (c *Ctx) AddMessageComponents(components []discordgo.MessageComponent) *Ctx {
	c.interactionResponse.Data.Components = append(c.interactionResponse.Data.Components, components...)
	return c
}

// Adds an Embed to the interaction response.
func (c *Ctx) AddEmbed(embed *discordgo.MessageEmbed) *Ctx {
	c.interactionResponse.Data.Embeds = append(c.interactionResponse.Data.Embeds, embed)
	return c
}

// Adds Files to the interaction response.
func (c *Ctx) AddFiles(files []*discordgo.File) *Ctx {
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
func (c *Ctx) SetComponents(actionRows interface{}) (*Ctx, error) {
	messageComponent, ok := actionRows.([]discordgo.MessageComponent)

	if !ok {
		return nil, errors.New("Must be ActionRow")
	}
	if c.interactionResponse.Data.Components != nil {
		c.interactionResponse.Data.Components = append(c.interactionResponse.Data.Components, messageComponent...)
		println(c.interactionResponse.Data.Components)
		return c, nil
	}

	c.interactionResponse.Data.Components = []discordgo.MessageComponent{}
	c.interactionResponse.Data.Components = append(c.interactionResponse.Data.Components, messageComponent...)
	println(c.interactionResponse.Data.Components)
	return c, nil
}
