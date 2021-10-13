package db

import (
	"fmt"
	"gorm.io/gorm/logger"

	//"net/url"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	//"gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
)

func NewMySQLGorm() *gorm.DB{
	//mySQLUser:="root"
	//mySQLPassword:="42306387"
	//mySQLHost:="localhost"
	//mySQLPort:="3306"
	//mySQLDBName:="MinhDatabase"
	//connectionString:=fmt.Sprintf(
	//	"sqlserver://%s:%s@%s?database%s",mySQLUser,url.QueryEscape(mySQLPassword),mySQLHost,mySQLPort,mySQLDBName)
	//connectionString:="root:42306387@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Println(connectionString)
	//db,err:= gorm.Open(sqlserver.Open(connectionString),&gorm.Config{Logger:logger.Default.LogMode(logger.Info)})
	dsn := "root:42306387@tcp(localhost:3306)/MinhDatabase"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger:logger.Default.LogMode(logger.Info)})
	if err!=nil{
		fmt.Println(err)
		panic("failed to connect to DB")
	}
	return db
}