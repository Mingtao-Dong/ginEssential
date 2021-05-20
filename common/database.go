package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB()*gorm.DB{
	driverName:="mysql"
	host:="localhost"
	port:="3306"
	database:="ginessential"
	username:="root"
	password:="dmt123"
	charset:="utf8"
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db,err:=gorm.Open(driverName ,args)
	if err!=nil{
		panic("failed to connect database,err: "+err.Error())
	}
	db.AutoMigrate(&model.User{})//gorm创建数据表
	DB=db
	return db

}
func GetDB()*gorm.DB{
	return DB
}
