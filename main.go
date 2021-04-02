/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 22:06:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 12:05:00
 * @Description: backedn of to-do list
 * @FilePath: /Local_Lab/toDoList/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

// 日志模板
// 登录验证模块

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
	Status bool   `json:"status"` // false: not finished, true: finished
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:lyc7758321321@(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	return err
}

func main() {
	// create database
	// sql: CREATE DATABASE todo;
	// link db
	if err := initMySQL(); err != nil {
		panic(err)
	}
	defer DB.Close()
	// create table
	DB.AutoMigrate(&Todo{})
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
			// 前端页面填写待办事项
			// 1. 从请求中拿出数据存入数据库
			var todo Todo
			c.BindJSON(&todo)
			// 2.存入数据库
			if err := DB.Create(&todo).Error; err != nil {
				// 3. 返回响应
				// error occur
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				// success
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var toDoList []Todo
			if err := DB.Find(&toDoList).Error; err != nil {
				// error occur
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				// success
				c.JSON(http.StatusOK, toDoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "invalid id",
				})
				return
			} else {
				var todo Todo
				if err := DB.Where("id=?", id).First(&todo).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{
						"error": err.Error(),
					})
					return
				} else {
					c.BindJSON(&todo)
					if err := DB.Save(&todo).Error; err != nil {
						c.JSON(http.StatusOK, gin.H{
							"error": err.Error(),
						})
						return
					}
				}
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "invalid id",
				})
				return
			} else {
				// 查找是否在数据库存在这条记录
				if err := DB.Where("id=?", id).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{
						"error": err.Error(),
					})
					return
				} else {
					// 删除
					if err := DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
						c.JSON(http.StatusOK, gin.H{
							"error": err.Error(),
						})
					}
				}
			}
		})
	}

	r.Run(":9090")
}
