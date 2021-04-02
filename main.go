/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 22:06:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 13:05:26
 * @Description: backedn of to-do list
 * @FilePath: /toDoList/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

// 日志模板
// 登录验证模块

package main

import (
	"toDoList/controller"
	"toDoList/db"
	"toDoList/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// create database
	// sql: CREATE DATABASE todo;
	// link db
	if err := db.InitMySQL(); err != nil {
		panic(err)
	}
	defer db.Close()
	// create table
	db.DB.AutoMigrate(&model.Todo{})
	// gin
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

	r.Run(":9090")
}
