package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/SimpleBookManagementAPI/internal/config"
	"github.com/rishad004/SimpleBookManagementAPI/internal/controllers"
	"github.com/rishad004/SimpleBookManagementAPI/internal/repository"
	"github.com/rishad004/SimpleBookManagementAPI/internal/routes"
	"github.com/rishad004/SimpleBookManagementAPI/internal/service"
	"github.com/spf13/viper"
)

func main() {

	r := gin.Default()

	if err := config.Config(); err != nil {
		log.Fatal("env error!", err)
	}

	db := config.DbConn()

	repo := repository.NewRepo(db)
	svc := service.NewSvc(repo)
	handlers := controllers.NewControllers(svc)

	routes.Routes(r, handlers)

	r.Run(viper.GetString("PORT"))

}
