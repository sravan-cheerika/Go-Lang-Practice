package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection_Link() (db *sql.DB){
	db,err:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/MyDB" )
	if err!= nil {
		fmt.Println(db, "connection error ")
	} else{
		fmt.Println(db, "connection success ")
	}
	return db
}

func CreateTable(db *sql.DB){
	query:=`create table table6(id int primary key auto_increment, name varchar(50))`
	ctx, cancelfun:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfun()
	resp, err:=db.ExecContext(ctx,query)
	if err != nil{
		fmt.Println(err, "query error ")
	} 
	row,err:=resp.RowsAffected()
	fmt.Println("Table created ",row)
}

func main(){
	db:=Connection_Link()
	CreateTable(db)
}