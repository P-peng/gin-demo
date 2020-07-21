package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

/**
@link: http://gorm.book.jasperxu.com/
*/
func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:Lv123456+@tcp(47.100.17.141:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	Eloquent.LogMode(true)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
