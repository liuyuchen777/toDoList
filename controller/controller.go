/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-03 21:40:58
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 13:09:03
 * @Description:
 * @FilePath: /toDoList/controller/controller.go
 * @GitHub: https://github.com/liuyuchen777
 */

package controller

import (
	"net/http"
	"toDoList/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
url -> controller -> logic -> model
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetToDoHandler(c *gin.Context) {
	var toDoList []model.Todo
	if err := model.GetTodoList(&toDoList); err != nil {
		// error occur
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// success
		c.JSON(http.StatusOK, toDoList)
	}
}

func CreateHandler(c *gin.Context) {
	// 前端页面填写待办事项
	// 1. 从请求中拿出数据存入数据库
	var todo model.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	if err := model.CreateTodo(&todo); err != nil {
		// 3. 返回响应
		// error occur
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// success
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	// judge whether have id in request
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		var todo model.Todo
		if err := model.SearchTodo1(id, &todo); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.BindJSON(&todo) // 从前端返回的JSON中获取对象
			if err := model.UpdateTodo(&todo); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}

func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		// search id in db
		if err := model.SearchTodo2(id); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			// do delete operation
			if err := model.DeleteTodo(id); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}
