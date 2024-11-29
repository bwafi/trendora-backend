package main

import (
	"fmt"
	"log"

	"github.com/bwafi/trendora-backend/internal/config"
	"github.com/gofiber/fiber/v3"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(viper)

	webPort := viper.GetInt("web.port")
	fmt.Println(webPort)
	err := app.Listen(fmt.Sprintf(":%d", webPort), fiber.ListenConfig{EnablePrefork: viper.GetBool("web.prefork")})
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}
