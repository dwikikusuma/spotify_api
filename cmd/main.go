package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"spotify_api/internal/configs"
	"spotify_api/pkg/internalsql"
)

func main() {
	fmt.Println("Hello")

	if err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	); err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cnfg := configs.Get()
	fmt.Println(cnfg)

	_, err := internalsql.Connect(cnfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed to Connect With DB")
	}
	r := gin.Default()
	_ = r.Run(cnfg.Service.Port)
}
