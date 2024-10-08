package discog

import "github.com/bwmarrin/discordgo"

type TextInput struct {
	textInput discordgo.TextInput
}

func NewTextInput(id string) *TextInput {
	return &TextInput{
		textInput: discordgo.TextInput{
			CustomID: id,
		},
	}
}

func (t *TextInput) SetLabel(label string) *TextInput {
	t.textInput.Label = label
	return t
}

func (t *TextInput) SetPlaceholder(placeholder string) *TextInput {
	t.textInput.Placeholder = placeholder
	return t
}

func (t *TextInput) SetStyle(style discordgo.TextInputStyle) *TextInput {
	t.textInput.Style = style
	return t
}

func (t *TextInput) SetValue(value string) *TextInput {
	t.textInput.Value = value
	return t
}

func (t *TextInput) Required() *TextInput {
	t.textInput.Required = true
	return t
}

func (t *TextInput) SetMinMax(min int, max int) *TextInput {
	t.textInput.MinLength = min
	t.textInput.MaxLength = max
	return t
}

func (t *TextInput) GetComponent() interface{} {
	return t.textInput
}
