package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewagius/wallet-api/api/model"
	"github.com/matthewagius/wallet-api/usecase/wallet"
	util "github.com/matthewagius/wallet-api/util"
)

var standardLogger = util.NewLogger()

type WalletController struct {
	WalletService wallet.WalletUseCase
}

func GetWalletController(service wallet.WalletUseCase) WalletController {
	return WalletController{WalletService: service}
}

func (controller *WalletController) GetWalletBalance(c *gin.Context) {
	standardLogger.IncomingRequest(c.Request.RequestURI)
	_, err := util.Validate(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		standardLogger.Error(err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := controller.WalletService.GetWallet(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
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
	_, err := util.Validate(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		standardLogger.Error(err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var data model.BodyData
	err = c.BindJSON(&data)

	if err != nil && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
		return
	}
	standardLogger.IncomingRequestWithBody(c.Request.RequestURI, string(jsonData))

	result, err := controller.WalletService.UpdateWallet(uint(id), "Credit", data.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
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
	_, err := util.Validate(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		standardLogger.Error(err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var data model.BodyData
	err = c.BindJSON(&data)

	if err != nil && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
		return
	}

	standardLogger.IncomingRequestWithBody(c.Request.RequestURI, string(jsonData))
	result, err := controller.WalletService.UpdateWallet(uint(id), "Debit", data.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		standardLogger.Error(err.Error())
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
