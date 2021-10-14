package mysqllib

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB // set to a global variable

func DBInitializer() {
	db, err := sqlx.Open("mysql", "root:shenal@tcp(localhost:3306)/restuarant?parseTime=true")
	//db, err := sqlx.Open("mysql", "root:coderx@/restuarant")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	DBClient = db
}
