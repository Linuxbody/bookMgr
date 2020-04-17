package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"bookMgr/models"
	"strconv"
)

// 显示书籍列表
func bookList(c *gin.Context) {
	bookList, err := models.QueryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": err,
		})
		return
	}

	// 返回查询的数据
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

// 插入书籍显示页面
func bookNew(c *gin.Context){
	c.HTML(http.StatusOK, "new_book.html", nil)
}

// 插入书籍 post 函数
func bookCreate(c *gin.Context){
	// form 表单获取数据
	var msg string
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price ,err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		msg = "无效的价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	// 查看返回的数据类型
	fmt.Printf("%T %T\n", titleVal, price)

	err = models.InsertBook(titleVal, price)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"msg": "插入书籍信息失败！",
		})
		return
	}

	// 插入书籍成功,返回书籍列表
	c.Redirect(http.StatusMovedPermanently, "/book/list")
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "ok",
	// })
}

// 删除书籍函数
func bookDelete(c *gin.Context){
	// 取Quuery string 参数
	idStr := c.Query("id")
	idVar, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": err,
		})
		return
	}
	// 执行数据删除
	err = models.DeleteBook(idVar)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "删除书籍信息失败！",
		})
		return
	}
	
	// 删除书籍成功,返回书籍列表
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

// 更新书籍信息
func bookEdit(c *gin.Context){
	// 1、取到用户编辑的是哪一本输， 从 QureyString 取到ID值
	idStr := c.Query("id")
	if len(idStr) == 0 {
		c.String(http.StatusBadRequest, "木有此ID书籍")
		return
	}
	// HTTP 请求传过来的参数通常都是string类型，根据自己的需要转换成相对于的数据类型
	bookId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "此ID无效")
		return
	}

	// 判断是否post
	if c.Request.Method == "POST"{
		// 1. 获取用户提交的数据
		titleVal := c.PostForm("title")
		priceStr := c.PostForm("price")

		priceVal, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的价格信息")
			return
		}
		// 2.去数据库更新对应的书籍信息
		err = models.EditBook(titleVal, priceVal, bookId)
		if err != nil {
			c.String(http.StatusInternalServerError, "更新数据失败")
			return
		}

		// 3.跳转回 /book/list 页面，查看是否修改成功
		c.Redirect(http.StatusMovedPermanently, "/book/list")

	}else{
		// 给模板渲染上原来的旧数据
		// 2. 根据id 取到的书籍信息
		bookObj,err := models.QueryBookById(bookId)
		if err != nil {
			c.String(http.StatusBadRequest, "木有此书籍")
			return
		}
		// 3. 把书籍数据渲染到页面上
		c.HTML(http.StatusOK, "edit_book.html", bookObj)
	}
}

func main() {
	// 程序启动就自动链接数据库
	err := models.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	// 定义访问的路由
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // 加载模板页面
	r.GET("/book/list", bookList)  // 查询所有书籍
	r.GET("/book/new", bookNew)   // get 数据添加页面
	r.POST("/book/create", bookCreate) // post 数据添加页面
	r.GET("/book/delete", bookDelete)  // 删除
	r.Any("/book/edit", bookEdit) // get 数据编辑页面

	r.Run(":8000")
}