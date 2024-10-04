package moderations

import (
	discog "github.com/kisshan13/discog/package"
)

func GetGroup() *discog.CommandGroup {
	group := discog.NewCommandGroup("Moderations", "Moderation Commands for the server", "1076465065093513248")

	group.AddCommand(GetKickCommand())

	return group
}
