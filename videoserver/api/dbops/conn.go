package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	// DSN: username:password@protocol(address)/dbname?param=value

	// 使用TCP连接
	// dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/video_server?charset=utf8mb4")

	// 使用unix_socket连接
	// user: pc0
	// password:
	// protocol: unix
	// address: /var/run/mysqld/mysqld.sock
	// dbname: video_server
	// param: charset=utf8mb4
	dbConn, err = sql.Open("mysql", "pc0:@unix(/var/run/mysqld/mysqld.sock)/video_server?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
}
