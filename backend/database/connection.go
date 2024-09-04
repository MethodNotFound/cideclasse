package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

// var dsn = "root:mysql@tcp(127.0.0.1:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"

var dsn = "host=localhost user=jureg password=password dbname=cideclasse port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var Connection, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
