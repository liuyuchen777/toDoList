/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 22:06:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 10:54:18
 * @Description:
 * @FilePath: /Local_Lab/toDoList/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// create database
	// sql: CREATE DATABASE todo;
	// link db
	db, err := gorm.Open("mysql", "root:lyc7758321321@(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gin
	r := gin.Default()

	// load static files
	r.Static("/static", "./static")
	// load file
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1 group
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	r.Run(":9090")
}
