package main

import (
	"GeekTime_Go/mysql"
	"fmt"
)

func main() {
	Test2()
}

// Test2 测试第二周课程
func Test2() {
	db, err := mysql.GetDbConnect()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println(db)
	}

	err = db.UpdateServiceRetryStatusByID(0, 10010)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Println("更新成功")
	}
}
