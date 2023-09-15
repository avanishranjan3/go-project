package Controllers

import (
	"fmt"
)

type User struct {
	InvestorId string `json:"investorId" gorm:"column:investorId"`
}

func FetchDb() {
	var dta User
	db, _ := GormDb()
	rowsResErr := db.Raw("select investorId from uat2OnBoardDb.InvestorAccountDetails").Find(&dta).Error //Exec("CREATE TABLE IF NOT EXISTS USERS (id varchar(40) primary key,Name varchar(1000))")
	if rowsResErr != nil {
		fmt.Println("DB rowsResError ", rowsResErr)
		return
	}
	fmt.Println("result: ", dta)

}
