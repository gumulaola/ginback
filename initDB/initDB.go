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
	db := " "
	DB, _ = sql.Open("mysql", db)
	fmt.Println("mysql connected")
}
