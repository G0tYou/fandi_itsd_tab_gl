package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	http "app/internal/delivery/http"
	service "app/internal/service"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	//employee routes
	su := service.NewServiceUser()
	http.NewUserHandler(e, su)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
