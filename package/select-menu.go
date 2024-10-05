package discog

import "github.com/bwmarrin/discordgo"

// Represents a SelectMenu
type SelectMenu struct {
	selectMenu discordgo.SelectMenu
}

// Represents a SelectMenuOption
type SelectMenuOption struct {
	option discordgo.SelectMenuOption
}

// Create a new SelectMenu
func NewSelectMenu(menuType discordgo.SelectMenuType) *SelectMenu {
	return &SelectMenu{
		selectMenu: discordgo.SelectMenu{
			MenuType: menuType,
		},
	}
}

// Adds placeholder to select menu
func (sm *SelectMenu) PlaceHolder(placeholder string) {
	sm.selectMenu.Placeholder = placeholder
}

// Sets Min & Max values for SelectMenu
func (sm *SelectMenu) SetMinMax(min int, max int) *SelectMenu {
	sm.selectMenu.MinValues = &min
	sm.selectMenu.MaxValues = max

	return sm
}

// Sets Default value for the SelectMenu
func (sm *SelectMenu) AddDefaultValue(id string, valueType discordgo.SelectMenuDefaultValueType) *SelectMenu {
	if len(sm.selectMenu.DefaultValues) == 0 {
		sm.selectMenu.DefaultValues = []discordgo.SelectMenuDefaultValue{
			{
				ID:   id,
				Type: valueType,
			},
		}

		return sm
	}

	sm.selectMenu.DefaultValues = append(sm.selectMenu.DefaultValues, discordgo.SelectMenuDefaultValue{
		ID:   id,
		Type: valueType,
	})

	return sm
}

// Creates a new SelectMenuOption
func NewSelectMenuOption(label string, value string, description string) *SelectMenuOption {
	return &SelectMenuOption{
		option: discordgo.SelectMenuOption{
			Label:       label,
			Value:       value,
			Description: description,
		},
	}
}

// Adds emoji to options
func (sm *SelectMenuOption) AnyEmoji(emoji *Emoji) *SelectMenuOption {
	sm.option.Emoji = &emoji.emoji

	return sm
}

// Mark the option as default selected for the SelectMenu
func (sm *SelectMenuOption) Default() *SelectMenuOption {
	sm.option.Default = true

	return sm
}
