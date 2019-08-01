package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
//因为我们需要在其他地方使用SqlDB这个变量，所以需要大写代表public
var SqlDB *sql.DB
//初始化方法
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
