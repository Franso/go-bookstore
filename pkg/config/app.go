package config

// this files main work is to return a variable called db that can be used with the rest
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Enables connection to the database
func Connect() {
	// database, name of user:DbPassword/databasename
	d, err := gorm.Open("mysql", "francis:Testpw@14@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
