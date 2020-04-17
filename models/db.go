package models
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库连接操作
var DB *sqlx.DB

func InitDB()(err error) {
	dsn := "root:ts123456@tcp(127.0.0.1:3306)/book"
	DB,err = sqlx.Open("mysql", dsn)
	if err != nil{
		fmt.Println(err)
	}
	DB.SetMaxOpenConns(20) // 最大连接数
	DB.SetMaxIdleConns(10) // 最大空闲连接数
	return
}

// 查询数据
func QueryAllBook()(bookList []*Book, err error){
	sqlStr := "select id,title,price from book"
	err = DB.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍信息失败！")
		return
	}
	return 
}

// 插入数据
func InsertBook(title string, price float64)(err error){
	sqlStr := "insert into book(title,price) values (?,?)"
	_, err = DB.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("书籍信息插入失败！！！")
		return
	}
	return
}

// 删除数据
func DeleteBook(id int64)(err error){
	sqlStr := "delete from book where id = ?"
	_,err = DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除书籍信息失败")
		return
	}
	return
}

// 查询单个书籍
func QueryBookById(id int64)(book Book, err error){
	sqlStr := "select id,title,price from book where id =?"
	err = DB.Get(&book, sqlStr, id)
	if err != nil{
		fmt.Println("查询书籍信息失败")
		return
	}
	return
}


// 编辑修改信息
func EditBook(title string, price float64, id int64)(err error){
	sqlStr := "update book set title=?, price=? where id =?"
	_,err = DB.Exec(sqlStr, title, price, id)
	if err != nil {
		fmt.Println("编辑书籍信息失败")
		return
	} 
	return
}