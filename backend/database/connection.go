package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "os"
)

var Connection *gorm.DB
var dsn string

func Setup() {
  dsn_env, exists := os.LookupEnv("DSN")

  if exists {
    dsn = dsn_env
  } else {
    dsn = "host=localhost user=jureg password=password dbname=cideclasse port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  }

  var err error
  Connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
}
