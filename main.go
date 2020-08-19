package main

import (
	"fmt"
	"github.com/masato25/blockchain-snademo/cmd"
	"github.com/masato25/blockchain-snademo/maltegolocal"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"

	"github.com/masato25/blockchain-snademo/db"
	ms "github.com/masato25/blockchain-snademo/models"
	"os"

	log "github.com/sirupsen/logrus"
)

func companyName2EthereumAddr(qdb *gorm.DB, input string) {
	companyName := input
	acct := ms.Account{Name: companyName}
	qdb.Where(&acct).First(&acct)
	log.Debugf("acct: %v", acct)
	TRX := maltegolocal.MaltegoTransform{}
	if acct.EthereumAddr != "" {
		NewEnt := TRX.AddEntity("III.ETHADDRESS", acct.EthereumAddr)
		NewEnt.SetType("III.ETHADDRESS")
		NewEnt.SetValue(acct.EthereumAddr)
		NewEnt.SetLinkColor("#FF0000")
		NewEnt.SetLinkColor(maltegolocal.LINK_COLOR_DEFAULT)
		NewEnt.SetLinkStyle(maltegolocal.LINK_STYLE_DOTTED)
		NewEnt.SetWeight(200)
		NewEnt.SetLinkLabel("mapping")
	} else {
		TRX.AddUIMessage("Not Found", maltegolocal.UIM_FATAL)
	}

	fmt.Println(TRX.ReturnOutput())
}

func ethereumAddr2companyName(qdb *gorm.DB, input string) {
	acct := ms.Account{EthereumAddr: input}
	qdb.Where(&acct).First(&acct)
	log.Debugf("acct: %v", acct)
	TRX := maltegolocal.MaltegoTransform{}
	if acct.EthereumAddr != "" {
		NewEnt := TRX.AddEntity("maltego.Company", acct.Name)
		NewEnt.SetType("maltego.Company")
		NewEnt.SetValue(acct.Name)
		NewEnt.SetLinkColor(maltegolocal.LINK_COLOR_DEFAULT)
		NewEnt.SetLinkStyle(maltegolocal.LINK_STYLE_DOTTED)
		NewEnt.SetWeight(200)
		NewEnt.SetLinkLabel("mapping")
	} else {
		TRX.AddUIMessage("Not Found", maltegolocal.UIM_FATAL)
	}
	fmt.Println(TRX.ReturnOutput())
}

func queryReceiveFrom(qdb *gorm.DB, input string) {
	acct := ms.Account{EthereumAddr: input}
	qdb.Where(&acct).Find(&acct)
	log.Debugf("acct: %v", acct)
	TRX := maltegolocal.MaltegoTransform{}
	transactions := []ms.Transactions{}
	flag := true
	if acct.ID != 0 {
		qdb.Set("gorm:auto_preload", true).Where(&ms.Transactions{ToAccountID: acct.ID}).Find(&transactions)
	} else {
		TRX.AddUIMessage("addrsss not found", maltegolocal.UIM_FATAL)
		flag = false
	}
	if len(transactions) > 0 {
		for _, t := range transactions {
			NewEnt := TRX.AddEntity("III.ETHADDRESS", t.GetFrom())
			NewEnt.SetType("III.ETHADDRESS")
			NewEnt.SetValue(t.GetFrom())
			NewEnt.SetLinkColor(maltegolocal.LINK_COLOR_2)
			NewEnt.SetWeight(200)
			NewEnt.SetLinkLabel(fmt.Sprintf("received %v %s", t.Value, t.Token.Unit))
		}
	} else if flag {
		TRX.AddUIMessage("transactions not found", maltegolocal.UIM_FATAL)
		flag = false
	}
	fmt.Println(TRX.ReturnOutput())
}

func querySendTo(qdb *gorm.DB, input string) {
	acct := ms.Account{EthereumAddr: input}
	qdb.Where(&acct).Find(&acct)
	log.Debugf("acct: %v", acct)
	TRX := maltegolocal.MaltegoTransform{}
	transactions := []ms.Transactions{}
	flag := true
	if acct.ID != 0 {
		qdb.Set("gorm:auto_preload", true).Where(&ms.Transactions{FromAccountID: acct.ID}).Find(&transactions)
	} else {
		TRX.AddUIMessage("addrsss not found", maltegolocal.UIM_FATAL)
		flag = false
	}
	if len(transactions) > 0 {
		for _, t := range transactions {
			NewEnt := TRX.AddEntity("III.ETHADDRESS", t.GetTo())
			NewEnt.SetType("III.ETHADDRESS")
			NewEnt.SetValue(t.GetTo())
			NewEnt.SetLinkColor(maltegolocal.LINK_COLOR_2)
			NewEnt.SetWeight(200)
			NewEnt.SetLinkLabel(fmt.Sprintf("send %v %s", t.Value, t.Token.Unit))
		}
	} else if flag {
		TRX.AddUIMessage("transactions not found", maltegolocal.UIM_FATAL)
		flag = false
	}
	fmt.Println(TRX.ReturnOutput())
}

func main() {
	cmd.Execute()
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	qdb := db.Conn()
	defer qdb.Close()
	if len(os.Args) == 1 {
		log.Fatal("ARGV is missing")
	}
	args := []string{""}
	args = append(args, os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	lt := maltegolocal.ParseLocalArguments(args)
	input := lt.Value
	funcName := viper.Get("transformName")
	switch funcName {
	case "companyName2EthereumAddr":
		companyName2EthereumAddr(qdb, input)
	case "ethereumAddr2CompanyName":
		ethereumAddr2companyName(qdb, input)
	case "queryReceiveFrom":
		queryReceiveFrom(qdb, input)
	case "querySendTo":
		querySendTo(qdb, input)
	}
}
