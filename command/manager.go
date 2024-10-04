package commands

import (
	moderations "github.com/kisshan13/discog/command/moderation"
	discog "github.com/kisshan13/discog/package"
)

func GetManager() *discog.CommandManager {
	manager := discog.NewCommandManager()

	manager.AddGroup(moderations.GetGroup())

	return manager
}
