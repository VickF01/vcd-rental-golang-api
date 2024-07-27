package main

import (
	"log"
	"vcd-rental/handler"
	"vcd-rental/middleware"
	"vcd-rental/user"
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

	db.AutoMigrate(&vcd.VCD{}, &user.User{})

	vcdRepository := vcd.NewRepo(db)
	vcdService := vcd.NewService(vcdRepository)

	vcdHandler := handler.VCDHandler(vcdService)

	userRepository := user.NewRepo(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.POST("/logout", userHandler.Logout)
	}

	apiAuth := router.Group("/api/v1")
	apiAuth.Use(middleware.AuthMiddleware())
	{
		apiAuth.GET("/", vcdHandler.RootHandler)
		apiAuth.GET("/vcd", vcdHandler.GetAllVCD)
		apiAuth.GET("/vcd/:id", vcdHandler.GetOneVCD)
		apiAuth.POST("/vcd/add", vcdHandler.CreateVCD)
		apiAuth.PUT("/vcd/edit/:id", vcdHandler.UpdateVCD)
		apiAuth.DELETE("/vcd/delete/:id", vcdHandler.DeleteVCD)
	}

	router.Run()
}
