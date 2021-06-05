package main

import (
	"GeekTime_Go/mysql"
	"fmt"
)

func main() {
	db, err := mysql.GetMysqlDbConnect()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println(db)
	}
}
