package initDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// db init
// todo:  divide the string field
func init() {
	DB, _ = sql.Open("mysql", "root:root@tcp(49.232.162.126:3306)/miniprogram?charset=utf8")
	fmt.Println("mysql connected")
}
