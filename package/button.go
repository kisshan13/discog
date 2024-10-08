package discog

import "github.com/bwmarrin/discordgo"

// Represents a Button Component
type Button struct {
	button discordgo.Button
}

// Creates a new Button Component
func NewButton() *Button {
	return &Button{
		button: discordgo.Button{},
	}
}

// Adds content to Button Component like id, label, url and disabled.
func (b *Button) SetContent(id string, label string, url string, disabled bool) *Button {
	b.button.Label = label
	b.button.URL = url
	b.button.CustomID = id
	b.button.Disabled = disabled

	return b
}

// Adds style to the Button Component
func (b *Button) SetStyle(style discordgo.ButtonStyle) *Button {
	b.button.Style = style

	return b
}

// Adds emoji to the Button Component
func (b *Button) AnyEmoji(emoji Emoji) *Button {
	b.button.Emoji = &emoji.emoji

	return b
}

// Returns the Button Component
func (b *Button) GetComponent() interface{} {
	return b.button
}
