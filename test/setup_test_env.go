package test

import (
	"github.com/masato25/blockchain-snademo/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {
	viper.SetConfigName(".env.test")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbfile := viper.GetString("DB_FILE")
	viper.Set("seed", true)
	os.Remove(dbfile)
	db.Conn()
}
