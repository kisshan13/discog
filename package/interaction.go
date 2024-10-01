package discog

import "github.com/bwmarrin/discordgo"

type InteractionResponse struct {
	interactionType discordgo.InteractionResponseType
	data            *discordgo.InteractionResponseData
}

func NewInteractionResponse() *InteractionResponse {
	return &InteractionResponse{
		interactionType: discordgo.InteractionResponseChannelMessageWithSource,
		data: &discordgo.InteractionResponseData{
			Content: "",
		},
	}
}
