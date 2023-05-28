package database

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// NewEntClient use if not wrap sql db
//
//goland:noinspection GoUnusedExportedFunction
func NewEntClient() *ent.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.configs.username"),
		viper.GetString("db.configs.password"),
		viper.GetString("db.configs.host"),
		viper.GetString("db.configs.port"),
		viper.GetString("db.configs.database"))
	log.Debugf("DSN ", dsn)

	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			log.Debugf(time.Now().Format("2006-01-02 15:04:05"), v)
		}
	}))

	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	//from docs define close on this function, but will impact cant create DB session on repository
	//defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Infof("initialized database x orm : success")
	return client
}
