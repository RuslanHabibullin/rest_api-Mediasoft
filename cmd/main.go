package main

import (
	"log"
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/handler"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/repository"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/service"

	"github.com/spf13/viper"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "12345",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	if err := ConfigInit(); err != nil {
		log.Fatal("Ошибка конфигурации")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(restapimediasoft.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("Не подключились")
	}
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
