package main

import (
	"demo/config"
	"demo/controller"
	"demo/mapper"
	"demo/model"
	"demo/router"
	"demo/service"
	"log"
)

func main() {
	db := config.InitDB()
	if err := model.Migrate(db); err != nil {
		log.Fatal("failed to migrate: ", err)
	}

	userMapper := &mapper.UserMapper{DB: db}
	userService := &service.UserService{Mapper: userMapper}
	userController := &controller.UserController{Service: userService}

	r := router.SetupRouter(userController)
	r.Run(":8080")
}
