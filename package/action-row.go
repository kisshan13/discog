package discog

import (
	"github.com/bwmarrin/discordgo"
)

// Represent a ActionRow
type ActionRow struct {
	components []discordgo.MessageComponent
}

// Creates a new ActionRow
func NewActionRow() *ActionRow {
	return &ActionRow{
		components: []discordgo.MessageComponent{},
	}
}

// Adds a component to the ActionRow
func (r *ActionRow) AddComponent(component discordgo.MessageComponent) {
	r.components = append(r.components, component)
}

// Method used to get the ActionRow components which represent a discordgo.ActionsRow structure
func (r *ActionRow) GetActionRow() discordgo.ActionsRow {
	return discordgo.ActionsRow{
		Components: r.components,
	}
}
