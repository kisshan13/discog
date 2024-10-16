package discog

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Command Group to manage related commands like grouping a set of commands for a guild,moderation, utils, user etc.
type CommandGroup struct {
	Name          string
	Description   string
	GuildId       string
	SlashCommands []*SlashCommand
	commands      map[string]*discordgo.ApplicationCommand
}

// Creates a new Command Group with the given name, description, and Guild ID.
func NewCommandGroup(name string, description string, guildId string) *CommandGroup {
	return &CommandGroup{
		Name:          name,
		Description:   description,
		GuildId:       guildId,
		SlashCommands: make([]*SlashCommand, 0),
		commands:      make(map[string]*discordgo.ApplicationCommand),
	}
}

// Adds a Slash Command to the command group.
func (cg *CommandGroup) AddCommand(command *SlashCommand) {
	cg.SlashCommands = append(cg.SlashCommands, command)

	cg.commands[command.Command.Name] = command.Command
}

// Registers all the Slash Commands in the command group to Discord.
func (cg *CommandGroup) register(session *discordgo.Session) {
	for _, command := range cg.SlashCommands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, cg.GuildId, command.Command)

		if err != nil {
			log.Printf("❎ [⚡ Discog ] [ CG : %s] Error registering command %s: %v\n", cg.Name, command.Command.Name, err)
			continue
		}

		cg.commands[cmd.Name] = cmd
		log.Printf("✅ [⚡ Discog ] [ CG : %s] Registered command /%s \n", cg.Name, command.Command.Name)
	}
}

// Removes all the Slash Commands in the command group from Discord.
func (cg *CommandGroup) unregister(session *discordgo.Session) {
	for _, command := range cg.commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, cg.GuildId, command.ID)

		if err != nil {
			log.Printf("[⚡ Discog ] [ CG : %s] Error deleting command /%s: %v\n", cg.Name, command.Name, err)
			continue
		}

		delete(cg.commands, command.Name)
		log.Printf("✅ [⚡ Discog ] Deleted command /%s \n", command.Name)
	}
}
