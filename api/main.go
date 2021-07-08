package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewagius/wallet-api/api/controller"
	_ "github.com/matthewagius/wallet-api/api/docs"
	"github.com/matthewagius/wallet-api/config"
	"github.com/matthewagius/wallet-api/entity"
	"github.com/matthewagius/wallet-api/repository"
	"github.com/matthewagius/wallet-api/usecase/wallet"
	util "github.com/matthewagius/wallet-api/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

// @securityDefinitions.apikey BearerAuth
// @Param Authorization header string true default(Bearer <Add access token here>)
// @in header
// @name Authorization
func main() {

	db := initDB()
	db.AutoMigrate(&entity.Wallet{})

	//if the config env is dev Insert Sample data to wallet table
	if env := config.ENV; env != "" {
		util.InsertSampleData(db)
	}
	walletRepo := repository.NewWalletMySQL(db)
	walletService := wallet.NewService(walletRepo)
	walletController := controller.GetWalletController(walletService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	version := api.Group("/v1")
	auth := version.Group("/auth")
	{
		auth.POST("/", controller.Login)
		auth.POST("/token/refresh", controller.RefreshToken)
	}
	controllers := version.Group("/wallets")
	{
		controllers.GET("/ping", controller.HealthLiveliness)
		controllers.GET("/:id/balance", TokenAuthMiddleware(), walletController.GetWalletBalance)
		controllers.POST("/:id/credit", TokenAuthMiddleware(), walletController.CreditWallet)
		controllers.POST("/:id/debit", TokenAuthMiddleware(), walletController.DebitWallet)
	}

	err := router.Run(config.API_PORT)
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	db, err = gorm.Open(mysql.Open(util.DBUrlConnection(
		util.NewDatabaseConfig())))

	if err != nil {
		fmt.Println("Status:", err)
	}

	return db
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
