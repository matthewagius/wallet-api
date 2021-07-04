package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matthewagius/wallet-api/api/controller"
	"github.com/matthewagius/wallet-api/config"
	"github.com/matthewagius/wallet-api/entity"
	"github.com/matthewagius/wallet-api/repository"
	"github.com/matthewagius/wallet-api/usecase/wallet"
	"github.com/matthewagius/wallet-api/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func main() {
	db := initDB()
	db.AutoMigrate(&entity.Wallet{})
	//TODO add data on load
	walletRepo := repository.NewWalletMySQL(db)
	walletService := wallet.NewService(walletRepo)
	walletController := controller.GetWalletController(walletService)

	router := gin.Default()
	api := router.Group("/api")
	version := api.Group("/v1")
	controllers := version.Group("/wallets")
	{
		controllers.GET("/ping", controller.HealthLiveliness)
		controllers.GET("/:id/balance", walletController.GetWalletBalance)
		controllers.POST("/:id/credit", walletController.CreditWallet)
		controllers.POST("/:id/debit", walletController.DebitWallet)
	}
	//running
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
