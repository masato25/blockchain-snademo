package test

import (
	"github.com/masato25/blockchain-snademo/db"
	ms "github.com/masato25/blockchain-snademo/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestAccountsSpec(t *testing.T) {
	qdb := db.GetDB()
	Convey("test init accounts", t, func() {
		accounts := []ms.Account{}
		qdb.Find(&accounts)
		So(len(accounts), ShouldEqual, 4)
	})

	Convey("test init tokens", t, func() {
		tokens := []ms.Token{}
		qdb.Find(&tokens)
		So(len(tokens), ShouldEqual, 1)
		token := tokens[0]
		So(token.Name, ShouldEqual, "medison1")
	})

	Convey("test init transactions", t, func() {
		transactions := []ms.Transactions{}
		qdb.Find(&transactions)
		log.Debug(transactions)
		So(len(transactions), ShouldEqual, 6)
	})

	Convey("test query address", t, func() {
		acct := ms.Account{EthereumAddr: "0x0000000000000000000000000000000000000001"}
		qdb.Where(&acct).Find(&acct)
		if qdb.Error != nil {
			log.Error(qdb.Error.Error())
		}
		transactions := []ms.Transactions{}
		//qdb.Where(&ms.Transactions{To: &acct}).Find(&transactions)
		//qdb.Find(&transactions)
		qdb.Set("gorm:auto_preload", true).Where(&ms.Transactions{ToAccountID: acct.ID}).Find(&transactions)

		//So(transactions[0].ToAccountID, ShouldEqual, "")
		//So(transactions[0].To, ShouldEqual, "")
		So(len(transactions), ShouldEqual, 1)
	})
}
