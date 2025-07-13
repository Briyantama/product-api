package routes

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/middlewares"
	"test-case-vhiweb/internal/routes/controllers"
	"test-case-vhiweb/internal/routes/repository"
	"test-case-vhiweb/internal/routes/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	tx := repository.NewTxRepository(db)

	userRepo := repository.NewUserRepository(db)
	vendorRepo := repository.NewVendorRepository(db)
	productRepo := repository.NewProductRepository(db)

	userUC := usecase.NewUserUsecase(userRepo, tx)
	vendorUC := usecase.NewVendorUsecase(vendorRepo, tx)
	productUC := usecase.NewProductUsecase(productRepo, vendorRepo, tx)

	userController := controllers.NewUserController(userUC)
	vendorController := controllers.NewVendorController(vendorUC)
	productController := controllers.NewProductController(productUC)

	r.POST("/auth/register", userController.Register)
	r.POST("/auth/login", userController.Login)

	protected := r.Group("/", middlewares.AuthMiddleware())
	protected.POST("/vendors", vendorController.RegisterVendor)
	protected.GET("/vendors", vendorController.GetVendorsByUserID)

	protected.POST("/products", productController.Create)
	protected.GET("/products/user", productController.GetProductByUserID)
	protected.GET("/products/vendor", productController.GetProductByVendorID)
	protected.PUT("/products/:id", productController.Update)
	protected.DELETE("/products/:id", productController.Delete)

	port := os.Getenv("SERVER_ADDRESS")

	server := &http.Server{
		Addr:    port,
		Handler: r.Handler(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Printf("Shutdown Server ...")
	timeout, strconvErr := strconv.Atoi(os.Getenv(constants.ENV_GRACEFUL_TIMEOUT_KEY))
	const defaultTimeout = 1
	if strconvErr != nil {
		timeout = defaultTimeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Printf("Server Shutdown:", err)
	}

	<-ctx.Done()
	logger.Log.Printf("timeout of %d seconds. \n", timeout)
	logger.Log.Printf("Server exiting")
}
