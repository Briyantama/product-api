package app

import (
	"test-case-vhiweb/internal/config"
	"test-case-vhiweb/internal/middlewares"
	"test-case-vhiweb/internal/routes"

	"github.com/gin-gonic/gin"
)

func App() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.Logger())
	r.Use(middlewares.ErrorMiddleware())

	db := config.InitDB()

	routes.SetupRoutes(r, db)
}
