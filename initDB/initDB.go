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
	DB, _ = sql.Open("mysql", "root:root@tcp(localhost:3306)/edu?charset=utf8")
	fmt.Println("mysql connected")
}
