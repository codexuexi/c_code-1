package c_code

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func MysqlConnect(MyUser, Password, Host, Port, Db string) (db *gorm.DB, e error) {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, e = gorm.Open("mysql", connArgs)
	if e != nil {
		return
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(time.Second * 15)
	return
}
