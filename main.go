/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 22:06:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 10:43:14
 * @Description:
 * @FilePath: /Local_Lab/toDoList/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// load static files
	r.Static("/static", "./static")
	// load file
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 待办事项
	// 添加
	// 查看
	// 修改
	// 删除

	r.Run(":9090")
}
