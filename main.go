package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db:=common.InitDB()//初始化数据库

	defer db.Close()//延迟关闭

	r:=gin.Default()
	r=CollectRoute(r)

	panic(r.Run())//8080

}





