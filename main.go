package main

import (
	"roketin/entity"
	"roketin/handler"
	"roketin/repository"
	"roketin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/roketin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate([]entity.Film{})
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()
	url := router.Group("/api/v1")
	url.POST("/create", handler.Create)
	url.PUT("/update/:id", handler.Update)

	router.Run(":8080")
}
