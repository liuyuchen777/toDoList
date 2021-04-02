/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-02 22:06:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 13:22:51
 * @Description: backedn of to-do list
 * @FilePath: /toDoList/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

// 日志模板
// 登录验证模块

package main

import (
	"toDoList/db"
	"toDoList/model"
	"toDoList/router"
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
	model.InitModel()
	// create router
	r := router.SetupRouter()
	// run
	r.Run(":9090")
}
