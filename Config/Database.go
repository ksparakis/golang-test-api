package Config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func BuildDB() (error){
	//ip_address := os.Getenv("MYSQL_IP")
	//dsn := fmt.Sprint("docker:docker@tcp(localhost:3306)/test_db?charset=utf8&parseTime=True&loc=Local", ip_address)
	initDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "docker:docker@tcp(localhost:3306)/test_db?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	DB = initDb

	return err
}
