package main

// Import connect mysql module
import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db = init_db()

	res, err := db.Query("SHOW TABLES")

	if err != nil {
		panic(err.Error())
	}

	for res.Next() {
		var table string
		err = res.Scan(&table)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(table)
	}

	defer db.Close()

	fmt.Println("End Main")
}

func init_db() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err.Error())
	}

	c := mysql.Config{
		DBName:    "golang_app_db",
		User:      "golang_app_user",
		Passwd:    "golang_app_password",
		Net:       "tcp",
		Addr:      "db:3306",
		Collation: "utf8mb4_general_ci",
		Loc:       jst,
	}

	fmt.Println(c.FormatDSN())

	db, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	return db
}
