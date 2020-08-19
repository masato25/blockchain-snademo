package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	ms "github.com/masato25/blockchain-snademo/models"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var qdb *gorm.DB

func seedData(db *gorm.DB) {
	ms.InitAccounts(db)
	ms.InitToken(db)
	ms.InitTransactions(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&ms.Account{})
	if db.Error != nil {
		log.Error(db.Error.Error())
	}
	db.AutoMigrate(&ms.Token{})
	if db.Error != nil {
		log.Error(db.Error.Error())
	}
	db.AutoMigrate(&ms.Transactions{})
	if db.Error != nil {
		log.Error(db.Error.Error())
	}
}

func Conn() *gorm.DB {
	dbfile := viper.GetString("DB_FILE")
	log.Debugf("db path: %s", dbfile)
	var err error
	qdb, err = gorm.Open("sqlite3", dbfile)
	if err != nil {
		log.Error(err.Error())
	}
	seeding := viper.GetBool("seed")
	if seeding {
		migration(qdb)
		seedData(qdb)
	}
	return qdb
}

func GetDB() *gorm.DB {
	if qdb == nil {
		viper.Set("seed", false)
		qdb = Conn()
	}
	return qdb
}
