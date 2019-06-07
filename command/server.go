package command

import (
	"go-api-boilerplate/dic"
	"go-api-boilerplate/route"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var serverPort string
	defaultServerPort := viper.GetString("SERVER_PORT")
	serverCmd.PersistentFlags().StringVar(&serverPort, "port", defaultServerPort, "Server port")
	viper.BindPFlag("SERVER_PORT", serverCmd.PersistentFlags().Lookup("port"))

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		router := route.Setup(dic.Builder)
		router.Run(":" + viper.GetString("SERVER_PORT"))
	},
}
