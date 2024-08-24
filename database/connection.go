package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var dsn = "root:mysql@tcp(127.0.0.1:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"
var Connection, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
