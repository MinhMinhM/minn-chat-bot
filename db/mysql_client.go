package db

import (
	"fmt"
	"gorm.io/gorm/logger"
	"net/url"
	"os"

	//"net/url"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	//"gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
)

func NewMySQLGorm() *gorm.DB{
	mySQLUser:=os.Getenv("mySqlUser")
	mySQLPassword:=os.Getenv("mySqlPassword")
	mySQLHost:=os.Getenv("mySqlHost")
	mySQLPort:=os.Getenv("mySqlPort")
	mySQLDBName:=os.Getenv("mySqlDBName")
	connectionString:=fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",mySQLUser,url.QueryEscape(mySQLPassword),mySQLHost,mySQLPort,mySQLDBName)
	//connectionString:="root:42306387@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Println(connectionString)
	//db,err:= gorm.Open(sqlserver.Open(connectionString),&gorm.Config{Logger:logger.Default.LogMode(logger.Info)})

	//dsn := "root:42306387@tcp(localhost:3306)/MinhDatabase"
	//dsn:= "admin:password@tcp(0.tcp.ap.ngrok.io:19123)/information_schema"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger:logger.Default.LogMode(logger.Info)})
	if err!=nil{
		fmt.Println(err)
		panic("failed to connect to DB")
	}
	return db
}