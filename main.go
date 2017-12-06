package main

import (
	"log"
	"fmt"

	"github.com/kataras/iris"

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type ResponseDelete struct {
	Message string `json:"message"`
}

func userHandler(ctx *iris.Context) {
	// var u []User

	// db.Find(&u)
	// return u
	// ctx.JSON(u)
	fmt.Printf("hahaha")
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:@/test")
	defer db.Close()
	
	if err != nil {
		log.Print(err)
	}

	app := iris.New()

	app.Get("/users", func(ctx iris.Context) {
		var u []User
		db.Find(&u)
		ctx.JSON(u)
	}) 

	app.Get("/users/{userid: int}", func(ctx iris.Context) {
		var u User
		userId, _ := ctx.Params().GetInt("userid")

		db.Where("id = ?", userId).Find(&u)
		ctx.JSON(u)
	}) 

	app.Post("/users", func(ctx iris.Context) {
		var u User
		ctx.ReadJSON(&u)
		db.Create(&u)
		db.Save(&u)
		ctx.JSON(u)

	}) 

	app.Put("/users/{userid: int}", func(ctx iris.Context) {
		var inputStructUser User
		var haveStructUser User
		var zero int

		userId, _ := ctx.Params().GetInt("userid")
		ctx.ReadJSON(&inputStructUser)
		db.Where("id = ?", userId).Find(&haveStructUser)
		if inputStructUser.Age != zero {
			haveStructUser.Age = inputStructUser.Age
		}
		if inputStructUser.Name != "" {
			haveStructUser.Name = inputStructUser.Name
		}
		haveStructUser.Id = userId
		db.Save(&haveStructUser)
		ctx.JSON(haveStructUser)

	})

	app.Delete("users/{userid: int}", func(ctx iris.Context) {
		var u User
		var rd ResponseDelete

		userId, _ := ctx.Params().GetInt("userid")
		db.Where("id = ?", userId).Delete(&u)

		rd.Message = "deleted"
		ctx.JSON(rd)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}