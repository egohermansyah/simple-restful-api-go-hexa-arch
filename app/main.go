package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	apiApp "simple-restful-api-go-hexa-arch/api/app"
	controllerRole "simple-restful-api-go-hexa-arch/api/app/v1/role"
	"simple-restful-api-go-hexa-arch/api/utils/custom_validator"
	"simple-restful-api-go-hexa-arch/app/utils/mongodb"
	"simple-restful-api-go-hexa-arch/business/role"
	"simple-restful-api-go-hexa-arch/config"
	modulesRole "simple-restful-api-go-hexa-arch/modules/role"
	"os"
	"os/signal"
	"time"
)

func serviceRole(db *mongo.Database) role.IService {
	repository, err := modulesRole.NewMongoDBRepository(db)
	if err != nil {
		panic(err)
	}
	service := role.NewService(repository)
	return service
}

func main() {
	mongoDb := mongodb.ConnectDatabase(config.GetConfigs().MasterMongoDb)
	roleService := serviceRole(mongoDb)
	roleController := controllerRole.NewController(roleService)

	e := echo.New()
	apiApp.RegisterPath(e, roleController)

	e.Validator = &custom_validator.BodyRequestValidator{Validator: validator.New()}

	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfigs().Port)

		if err := e.Start(address); err != nil {
			logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
