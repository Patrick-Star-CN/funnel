package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	err := godotenv.Load()
	log.Printf("Load Configs: %v", err)
	if err != nil {
		log.Fatal(err)
	}
}
func SetupConfigs(r *gin.Engine) {
	log.Print("Setup Configs....")
	RedisInit()
	log.Print("Setup Configs Finish....")
}
