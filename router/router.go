/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-03 22:46:24
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 13:18:19
 * @Description:
 * @FilePath: /toDoList/router/router.go
 * @GitHub: https://github.com/liuyuchen777
 */

package router

import (
	"toDoList/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// load static files
	r.Static("/static", "./static")
	// load file
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", controller.IndexHandler)

	// v1 group
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateHandler)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetToDoHandler)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateHandler)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteHandler)
	}

	return r
}
