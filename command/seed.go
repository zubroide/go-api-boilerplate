package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zubroide/go-api-boilerplate/db/seeds"
	"github.com/zubroide/go-api-boilerplate/dic"
	"gorm.io/gorm"
)

var seederName string

func init() {
	seedCmd.PersistentFlags().StringVar(&seederName, "seeder", "", "Seeder name")
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run seeder",
	Run: func(cmd *cobra.Command, args []string) {

		// get gormDB from dic
		gormDB, ok := dic.Container.Get(dic.Db).(*gorm.DB)
		if !ok || gormDB == nil {
			fmt.Println("DB instance does not resolved")
		}

		// prepare auto seeders collection
		auto := seeds.NewSeeds(gormDB)
		auto.AppendSeeder(&seeds.SeedUsers{SeederBase: seeds.SeederBase{"users"}})

		// prepare manual seeder collection
		manual := seeds.NewSeeds(gormDB)
		manual.AppendSeeder(&seeds.SeedUsers{SeederBase: seeds.SeederBase{"users"}})

		// run target seeders
		if len(seederName) > 0 {
			if err := manual.RunSeederByName(seederName); nil != err {
				panic(err)
			}
		} else {
			if err := auto.RunSeeds(); nil != err {
				panic(err)
			}
		}
	},
}
