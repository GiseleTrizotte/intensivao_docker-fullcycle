// Abrindo uma conexão com o banco de dados e realizando im insert em go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Query("Insert into exemplo values(1, 'Full Cycle')")

}
