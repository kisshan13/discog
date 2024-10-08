package discog

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Component interface {
	GetComponent() interface{}
}

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
func (r *ActionRow) AddComponent(component Component) error {

	isValid, ok := component.GetComponent().(discordgo.MessageComponent)

	if !ok {
		return fmt.Errorf("component is not a valid message component")
	}

	r.components = append(r.components, isValid)
	return nil
}

// Method used to get the ActionRow components which represent a discordgo.ActionsRow structure
func (r *ActionRow) GetComponent() interface{} {
	return discordgo.ActionsRow{
		Components: r.components,
	}
}
