package discog

import "github.com/bwmarrin/discordgo"

// Represents an Emoji
type Emoji struct {
	emoji discordgo.ComponentEmoji
}

// Create a new Emoji
func NewEmoji(name string, animated bool, id string) *Emoji {
	return &Emoji{
		emoji: discordgo.ComponentEmoji{
			Name:     name,
			Animated: animated,
			ID:       id,
		},
	}
}
