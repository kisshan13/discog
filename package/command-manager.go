package discog

// Represents a root command manager for the bot. It holds all the information needed to create & execute commands.
type CommandManager struct {
	groups                  []*CommandGroup
	modalHandler            func(c *Ctx)
	messageComponentHandler func(c *Ctx)
}

// Creates a new instance of CommandManager.
func NewCommandManager() *CommandManager {
	return &CommandManager{groups: []*CommandGroup{}}
}

// Adds a command group to the command manager.
func (cm *CommandManager) AddGroup(group *CommandGroup) {
	cm.groups = append(cm.groups, group)
}

func (cm *CommandManager) AddModalInteractionHandler(handler func(c *Ctx)) {
	cm.modalHandler = handler
}

func (cm *CommandManager) AddMessageComponentInteractionHandler(handler func(c *Ctx)) {
	cm.messageComponentHandler = handler
}

// Finds a command by its name. Returns the handler function if found, nil otherwise.
func (cm *CommandManager) FindCommand(name string) func(ctx *Ctx) {
	if len(cm.groups) != 0 {
		for _, group := range cm.groups {
			for _, command := range group.SlashCommands {
				println(command.Command.Name)
				if command.Command.Name == name {
					return command.Handler
				}
			}
		}
		return nil
	}

	return nil
}

// Registers all command groups with the bot's Discord session.
func (cm *CommandManager) register(bot *Bot) {
	for _, group := range cm.groups {
		group.register(bot.Session)
	}
}

/*
Unregisters all command groups with the bot's Discord session.
This should be done when the bot is shutting down.
This is necessary to prevent any remaining commands from being executed.
It is recommended to use a shutdown hook in your main function to call this function.
For example: `defer cm.Unregister(bot)` in your main function.
This function should not be called manually.
It is automatically called when the bot is shut down.
*/
func (cm *CommandManager) unregister(bot *Bot) {
	for _, group := range cm.groups {
		group.unregister(bot.Session)
	}
}
