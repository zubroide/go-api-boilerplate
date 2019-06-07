package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("APP_NAME") + " 0.0.1 -- HEAD")
	},
}