package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewagius/wallet-api/api/model"
	"github.com/matthewagius/wallet-api/usecase/wallet"
)

type WalletController struct {
	WalletService wallet.WalletUseCase
}

func GetWalletController(service wallet.WalletUseCase) WalletController {
	return WalletController{WalletService: service}
}

func (controller *WalletController) GetWalletBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := controller.WalletService.GetWallet(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toJson := &model.Wallet{
		Id:           data.Id,
		UserId:       data.UserId,
		CurrencyCode: data.CurrencyCode,
		Amount:       data.Amount,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
	c.JSON(http.StatusOK, gin.H{"wallet": toJson})
}

func (controller *WalletController) CreditWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Wallet
	err := c.BindJSON(&data)

	if err != nil && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.WalletService.UpdateWallet(uint(id), "Credit", data.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toJson := &model.Wallet{
		Id:           result.Id,
		UserId:       result.UserId,
		CurrencyCode: result.CurrencyCode,
		Amount:       result.Amount,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"wallet": toJson})
}

func (controller *WalletController) DebitWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Wallet
	err := c.BindJSON(&data)

	if err != nil && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.WalletService.UpdateWallet(uint(id), "Debit", data.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toJson := &model.Wallet{
		Id:           result.Id,
		UserId:       result.UserId,
		CurrencyCode: result.CurrencyCode,
		Amount:       result.Amount,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"wallet": toJson})
}
