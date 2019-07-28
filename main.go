// helloworld project main.go
package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mippm_1"
)

//Db数据库连接池
var DB *sql.DB

var Info = struct {
	inforId     string
	heading     string
	summary     string
	content     string
	iamurl      string
	imghoturl   string
	publishDate string
	source      string
	autor       string
	newstype    string
	presdate    string
	updateDate  string
}{}

//注意方法名大写，就是public
//1. 需要先按照diver,go get -u github.com/go-sql-driver/mysql,并记得引入_ "github.com/go-sql-driver/mysql"会做初始化
func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
	addInfo()
	getCountInformation()
	getInformationById("aaa")
	deleteInbyId("aaa")
}

func addInfo() {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("failed to begin something")
	}
	stmt, err := tx.Prepare("insert into information (`infoid`,`heading`,`summary`,`content`,`imgurl`,`imghoturl`,`publishDate`,`source`,`autor`,`newstype`,`prestate`,`update_date`)values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("prepare sql fail")
	}
	resq, err := stmt.Exec("aaa", "2", "3", "4", "5", "1", "2017-07-05", "3", "4", "5", "1", "2017-07-05")
	if err != nil {
		fmt.Println(err.Error())
	}
	tx.Commit()
	fmt.Println(resq.LastInsertId())
}

func deleteInbyId(id string) {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("failed to begin something")
	}
	stmt, err := tx.Prepare("delete from information where infoid = ?")
	if err != nil {
		fmt.Println("prepare sql fail")
	}
	resq, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("exec sql failed")
	}
	fmt.Println(resq.LastInsertId())
	tx.Commit()
}

func getCountInformation() {
	var count int
	row := DB.QueryRow("select count(*) from information")
	row.Scan(&count)
	fmt.Println(count)
}

func getInformationById(id string) {
	in := struct {
		inforId     string
		heading     string
		summary     string
		content     string
		iamurl      string
		imghoturl   string
		publishDate string
		source      string
		autor       string
		newstype    string
		presdate    string
		updateDate  string
	}{}
	in.inforId = id
	fmt.Println(in.inforId)
	row := DB.QueryRow("select * from information where infoid = ?", id)
	fmt.Println(row)
	err := row.Scan(&in.inforId, &in.heading, &in.summary, &in.content, &in.iamurl, &in.imghoturl, &in.publishDate, &in.source, &in.autor, &in.newstype, &in.presdate, &in.updateDate)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(in)
}
