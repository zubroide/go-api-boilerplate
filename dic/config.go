package dic

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig() {
	var err error

	viper.SetConfigFile("base.env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Override config parameters from environment variables if specified
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}
