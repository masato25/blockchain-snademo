package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Account struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);unique_index:account_name"`
	EthereumAddr string `gorm:"type:varchar(100);unique_index:account_ethaddr"`
	Nonce        int    `gorm:""`
}

func InitAccounts(qdb *gorm.DB) {
	errors := []error{}

	d := Account{Name: "管理局", EthereumAddr: "0x0000000000000000000000000000000000000001"}
	qdb.Create(&d)
	errors = append(errors, qdb.Error)
	d2 := Account{Name: "公司A", EthereumAddr: "0x0000000000000000000000000000000000000002"}
	qdb.Create(&d2)
	errors = append(errors, qdb.Error)
	d3 := Account{Name: "公司B", EthereumAddr: "0x0000000000000000000000000000000000000003"}
	qdb.Create(&d3)
	errors = append(errors, qdb.Error)
	d4 := Account{Name: "公司C", EthereumAddr: "0x0000000000000000000000000000000000000004"}
	qdb.Create(&d4)
	errors = append(errors, qdb.Error)

	for _, err := range errors {
		if err != nil {
			log.Error(err.Error())
		}
	}
}
