package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/student?charset=utf8"  )
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	insertDB(db)
}

func insertDB(db *sql.DB)  {
	stmt,err := db.Prepare("insert into test(id,name) values (2019210624,'沈怡然')")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec("12222")
	fmt.Println(stmt)
}