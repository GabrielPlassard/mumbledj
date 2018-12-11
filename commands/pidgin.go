package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
)

type PidginCommand struct{}

func (c *PidginCommand) Aliases() []string {
	return viper.GetStringSlice("commands.pidgin.aliases")
}

func (c *PidginCommand) Description() string {
	return viper.GetString("commands.pidgin.description")
}

func (c *PidginCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.pidgin.is_admin")
}

func (c *PidginCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	return new(AddCommand).Execute(user, "https://youtube.com/watch?v=XF9Me96KMU4")
}
