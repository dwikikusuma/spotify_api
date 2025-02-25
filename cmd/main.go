package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"spotify_api/internal/configs"
	memberhsipHandler "spotify_api/internal/handler/memberships"
	"spotify_api/internal/model/memberhsips"
	membersipRepo "spotify_api/internal/repository/memberships"
	membershipService "spotify_api/internal/service/memberships"
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

	db, err := internalsql.Connect(cnfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed to Connect With DB")
	}

	err = db.AutoMigrate(&memberhsips.User{})
	if err != nil {
		log.Fatalf("failed to migrate user")
	}

	membershipsRepo := membersipRepo.NewRepository(db)
	membershipsService := membershipService.NewService(membershipsRepo, cnfg)
	membershipsHandler := memberhsipHandler.NewHandler(membershipsService)
	membershipsHandler.RegisterRoutes()

	r := gin.Default()
	_ = r.Run(cnfg.Service.Port)
}
