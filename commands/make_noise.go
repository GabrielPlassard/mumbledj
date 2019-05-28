package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
)

type MakeNoiseCommand struct{}

func (c *MakeNoiseCommand) Aliases() []string {
	return viper.GetStringSlice("commands.makenoise.aliases")
}

func (c *MakeNoiseCommand) Description() string {
	return viper.GetString("commands.makenoise.description")
}

func (c *MakeNoiseCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.makenoise.is_admin")
}

func (c *MakeNoiseCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	return new(AddCommand).Execute(user, "https://www.youtube.com/watch?v=ipN_JADAFIM")
}
