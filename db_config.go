package main

import "fmt"

const (
	MYSQL_USERNAME = "admin"
	MYSQL_PASSWORD = "xxxxx"
	MYSQL_HOST     = "localhost"
	MYSQL_PORT     = "3306"
	MYSQL_DATABASE = "employees"
	MYSQL_TIMEOUT  = "5s"
)

var MYSQL_DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=%v",
	MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE, MYSQL_TIMEOUT)

const (
	MONGO_USERNAME   = "admin"
	MONGO_PASSWORD   = "xxxxxxx"
	MONGO_HOST       = "localhost"
	MONGO_PORT       = "27017"
	MONGO_DATABASE   = "users"
	MONGO_COLLECTION = "user"
)

var MONGODB_URI = fmt.Sprintf("mongodb://%v:%v@%v:%v/test?authSource=%v",
	MONGO_USERNAME, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT, MONGO_USERNAME)
