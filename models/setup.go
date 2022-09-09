package models

import (
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:pwd123@tcp(127.0.0.1:3310)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error")
	}

	// rdb := redis.NewClient(&redis.Options{
    //     Addr:     "localhost:6379",
    //     Password: "", // no password set
    //     DB:       0,  // use default DB
	// })
	
	// err := rdb.Set(ctx, "key", "value", 0).Err()
    // if err != nil {
    //     panic(err)
    // }

    // val, err := rdb.Get(ctx, "key").Result()
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println("key", val)


	db.AutoMigrate(&Article{})

    DB = db
}