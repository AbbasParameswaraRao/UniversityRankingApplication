package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	InitLogger()
	InitCache()

	config, err := LoadConfig()
	if err != nil {
		LogError(err)
		return
	}

	db, err := ConnectDB(config)
	if err != nil {
		LogError(err)
		return
	}
	DB = db
	defer DB.Close()
	ReadAndLoad()

	router := gin.Default()
	router.Use(LoggerMiddleware())

	router.GET("/universities", GetUniversities)
	router.GET("/university/:rank", GetUnivByRank)
	router.DELETE("/universitydel/:rank",deleteUniversity)
	router.PUT("/updateuniversity/:rank",updateUniversity)

	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Printf("Server is running on %s\n", addr)

	err = router.Run(addr)
	if err != nil {
		LogError(err)
		return
	}
}
