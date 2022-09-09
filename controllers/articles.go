package controllers

import (
	"fmt"

	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"alpha-indo-soft/models"

	"github.com/go-redis/redis/v8"

	"encoding/json"
)

type ArticleInput struct {
	Author string `json:"author" binding: "required"`
	Title string `json:"title" binding: "required"`
	Body string `json:"body" binding: "required"`
}

var ctx = context.Background()

func CreateArticle(c *gin.Context) {
	var articleInput ArticleInput

	err := c.ShouldBindJSON(&articleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	article := models.Article{Author: articleInput.Author, Title: articleInput.Title, Body: articleInput.Body}
  	models.DB.Create(&article)

	c.JSON(http.StatusOK, gin.H{"data": article})
}

func FindArticles(c *gin.Context){
	var articles []models.Article
	models.DB.Find(&articles)

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
        DB:       0,  // use default DB
	})
	
	// convert struct to json string
	jsonBytes, err := json.Marshal(articles)

    err = rdb.Set(ctx, "data", string(jsonBytes), 0).Err()
    if err != nil {
        panic(err)
    }

    cache, err := rdb.Get(ctx, "data").Result()
    if err != nil {
        panic(err)
	}else{
		c.JSON(http.StatusOK, gin.H{"data": articles})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": cache})
}

func ArticleAuthorHandler(c *gin.Context){
	var author []models.Article

	if err := models.DB.Where("author = ?", c.Param("author")).Find(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
        DB:       0,  // use default DB
	})
	
	// convert struct to json string
	jsonBytes, err := json.Marshal(author)

    err = rdb.Set(ctx, "data", string(jsonBytes), 0).Err()
    if err != nil {
        panic(err)
    }

    cache, err := rdb.Get(ctx, "data").Result()
    if err != nil {
        panic(err)
	}else{
		c.JSON(http.StatusOK, gin.H{"data": cache})
		return
	}

	fmt.Println(string(jsonBytes), err) // {"message":"hello"} <nil>

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func ArticleQueryHandler(c *gin.Context){
	title := c.Query("title")
	body := c.Query("body")

	var author []models.Article

	if (title != "" && body != ""){
		if err := models.DB.Find(&author, "title LIKE ? AND body LIKE ?", "%"+title+"%", "%"+body+"%").Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}
	
	if (title != ""){
		if err := models.DB.Where("title LIKE ?", "%"+title+"%").Find(&author).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}

	if (body != ""){
		if err := models.DB.Where("body LIKE ?", "%"+body+"%").Find(&author).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
        DB:       0,  // use default DB
	})
	
	// convert struct to json string
	jsonBytes, err := json.Marshal(author)

    err = rdb.Set(ctx, "data", string(jsonBytes), 0).Err()
    if err != nil {
        panic(err)
    }

    cache, err := rdb.Get(ctx, "data").Result()
    if err != nil {
        panic(err)
	}else{
		c.JSON(http.StatusOK, gin.H{"data": cache})
		return 
	}

	fmt.Println(string(jsonBytes), err) // {"message":"hello"} <nil>

	c.JSON(http.StatusOK, gin.H{"data": author})
}
 