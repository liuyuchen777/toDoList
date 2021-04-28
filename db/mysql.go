/*
 * @Author: Liu Yuchen
 * @Date: 2021-04-03 21:46:49
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-04-02 13:22:15
 * @Description:
 * @FilePath: /toDoList/db/mysql.go
 * @GitHub: https://github.com/liuyuchen777
 */

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:LYClyc7758321321!@(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	return err
}

func Close() {
	DB.Close()
}