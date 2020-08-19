package models

import (
	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	From          Account `gorm:"foreignkey:FromAccountID;PRELOAD:true"`
	FromAccountID uint
	To            Account `gorm:"foreignkey:ToAccountID;PRELOAD:true"`
	ToAccountID   uint
	Token         Token `gorm:"foreignkey:"TokenID;PRELOAD:true"`
	TokenID       uint
	Value         int
}

func (self *Transactions) GetFrom() string {
	if self.From.ID == 0 {
		return "0x0"
	} else {
		return self.From.EthereumAddr
	}
}
func (self *Transactions) GetTo() string {
	if self.To.ID == 0 {
		return "0x0"
	} else {
		return self.To.EthereumAddr
	}
}

func InitTransactions(qdb *gorm.DB) {
	person1 := Account{Name: "管理局"}
	qdb.Where(&person1).First(&person1)
	person2 := Account{Name: "公司A"}
	qdb.Where(&person2).First(&person2)
	person3 := Account{Name: "公司B"}
	qdb.Where(&person3).First(&person3)
	person4 := Account{Name: "公司C"}
	qdb.Where(&person4).First(&person4)

	token := Token{Symbol: "md1"}
	qdb.Where(&token).First(&token)

	qdb.Create(&Transactions{ToAccountID: person1.ID, TokenID: token.ID, Value: 1000})
	qdb.Create(&Transactions{FromAccountID: person1.ID, ToAccountID: person2.ID, TokenID: token.ID, Value: 300})
	qdb.Create(&Transactions{FromAccountID: person1.ID, ToAccountID: person3.ID, TokenID: token.ID, Value: 200})
	qdb.Create(&Transactions{FromAccountID: person1.ID, ToAccountID: person4.ID, TokenID: token.ID, Value: 50})

	qdb.Create(&Transactions{FromAccountID: person2.ID, ToAccountID: person4.ID, TokenID: token.ID, Value: 100})
	qdb.Create(&Transactions{FromAccountID: person4.ID, ToAccountID: person3.ID, TokenID: token.ID, Value: 30})
}
