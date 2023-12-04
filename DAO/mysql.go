package DAO

// Data Access Object
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:wangqihang@/bubble?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return Db.DB().Ping()
}
func Close() {
	Db.Close()
}
