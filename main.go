package main

import (
	"log"
	"vcd-rental/handler"
	"vcd-rental/vcd"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=vcd_rental port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error!")
	}

	db.AutoMigrate(&vcd.VCD{})

	vcdRepository := vcd.NewRepo(db)
	vcdService := vcd.NewService(vcdRepository)

	vcdHandler := handler.VCDHandler(vcdService)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/", vcdHandler.RootHandler)
		api.GET("/vcd", vcdHandler.GetAllVCD)
		api.GET("/vcd/:id", vcdHandler.GetOneVCD)
		api.POST("/vcd/add", vcdHandler.CreateVCD)
		api.PUT("/vcd/edit/:id", vcdHandler.UpdateVCD)
		api.DELETE("/vcd/delete/:id", vcdHandler.DeleteVCD)
	}

	router.Run()
}
