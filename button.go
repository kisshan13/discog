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

// Adds label to Button.
func (b *Button) SetLabel(label string) *Button {
	b.button.Label = label
	return b
}

// Disabled the button.
func (b *Button) SetDisabled() *Button {
	b.button.Disabled = true
	return b
}

// Sets ID for the button.
func (b *Button) SetID(id string) *Button {
	b.button.CustomID = id
	return b
}

// Sets URL for the button.
func (b *Button) SetURL(url string) *Button {
	b.button.URL = url
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
