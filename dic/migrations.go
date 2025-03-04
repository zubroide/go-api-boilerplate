package dic

import (
	_ "github.com/lib/pq"            // for migrations
	_ "github.com/steinbacher/goose" // for migrations
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	ReadConfig()
	InitContainer()
	DB = Container.Get(Db).(*gorm.DB)
}
