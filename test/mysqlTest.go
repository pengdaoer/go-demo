package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //注册MySQL驱动
	"github.com/jmoiron/sqlx"
	"web-demo-interface-test/model/user"
)

func SelectUserX(username string) (err error) {
	var db *sqlx.DB
	//var userId string
	var user user.User
	db, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/scott")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Get(&user, "select userId from user_table where username=?", username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.UserId)
	return
}

func selectUserY(username string) (err error) {
	var db *sql.DB
	dsn := "root:123456@tcp(localhost:3306)/scott"
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")
	defer db.Close()
	result := db.QueryRow("select userid from user_table where username=?", username)
	var userId string
	// userId 通过Scan()方法之后，才会拿到值
	result.Scan(&userId)
	fmt.Println(userId)
	return
}

func test() (int, int) {
	return 100, 200
}

func main() {
	//selectU("user")
	//SelectUserX("user")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 向ch2通道发送数据的goroutine
	go func() {
		ch2 <- "World" // 发送数据到ch2通道
	}()

	// 向ch1通道发送数据的goroutine
	go func() {
		ch1 <- "Hello" // 发送数据到ch1通道
	}()

	// 使用select监听多个通道
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}
}
