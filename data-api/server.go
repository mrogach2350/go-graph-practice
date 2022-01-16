package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"net/http"
)

type User struct {
	UUID      string `faker:"uuid_digit"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Email     string `faker:"email"`
	Password  string `faker:"password"`
}

type Post struct {
	UUID   string `faker:"uuid_digit"`
	Body   string `faker:"paragraph"`
	Title  string `faker:"sentence"`
	UserId string
}

type Comment struct {
	UUID   string `faker:"uuid_digit"`
	Body   string `faker:"paragraph"`
	Name   string `faker:"name"`
	PostId string
}

type Channel struct {
	UUID string `faker:"uuid_digit"`
}

type Message struct {
	UUID      string `faker:"uuid_digit"`
	Body      string `faker:"paragraph"`
	ChannelId string
}

func createUser() User {
	var u User
	err := faker.FakeData(&u)
	if err != nil {
		fmt.Println(err)
	}
	return u
}

func createPost(uuid string) Post {
	var p Post
	err := faker.FakeData(&p)
	if err != nil {
		fmt.Println(err)
	}
	p.UserId = uuid
	return p
}

func createComment(uuid string) Comment {
	var c Comment
	err := faker.FakeData(&c)
	if err != nil {
		fmt.Println(err)
	}
	c.PostId = uuid
	return c
}

func main() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	userMigrateError := db.AutoMigrate(&User{})
	if userMigrateError != nil {
		fmt.Println(userMigrateError)
	}
	postMigrateError := db.AutoMigrate(&Post{})
	if postMigrateError != nil {
		fmt.Println(postMigrateError)
	}
	commentMigrateError := db.AutoMigrate(&Comment{})
	if commentMigrateError != nil {
		fmt.Println(commentMigrateError)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	api.GET("/users", func(context *gin.Context) {
		var users []User
		db.Find(&users)
		context.JSON(http.StatusOK, gin.H{"data": users})
	})

	api.GET("/users/:uuid", func(context *gin.Context) {
		var user User
		db.Where("uuid = ?", context.Param("uuid")).First(&user)
		context.JSON(http.StatusOK, gin.H{"data": user})
	})

	api.GET("/users/seed", func(context *gin.Context) {
		var users []User
		for i := 0; i < 10; i++ {
			users = append(users, createUser())
		}
		db.Create(users)
		context.JSON(http.StatusOK, gin.H{"success": true})
	})

	api.GET("/posts", func(context *gin.Context) {
		var posts []Post
		db.Find(&posts)
		context.JSON(http.StatusOK, gin.H{"data": posts})
	})

	api.GET("/posts/seed", func(context *gin.Context) {
		var posts []Post
		var users []User
		db.Find(&users)
		for _, user := range users {
			for i := 0; i < 10; i++ {
				posts = append(posts, createPost(user.UUID))
			}
		}
		db.Create(posts)
		context.JSON(http.StatusOK, gin.H{"success": true})
	})

	api.GET("/comments", func(context *gin.Context) {
		var comments []Comment
		db.Find(&comments)
		context.JSON(http.StatusOK, gin.H{"data": comments})
	})

	api.GET("/comments/seed", func(context *gin.Context) {
		var posts []Post
		var comments []Comment
		db.Find(&posts)
		for _, post := range posts {
			for i := 0; i < 15; i++ {
				comments = append(comments, createComment(post.UUID))
			}
		}
		db.Create(comments)
		context.JSON(http.StatusOK, gin.H{"success": true})
	})

	api.GET("/comments/delete", func(context *gin.Context) {
		var comments []Comment
		db.Find(&comments)
		for _, comment := range comments {
			db.Where("uuid = ?", comment.UUID).Delete(comment)
		}
		context.JSON(http.StatusOK, gin.H{"success": true})
	})

	r.Run(":7000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
