package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ashu23042000/coffee-supply-chain/backend/gatewaycli"
	"github.com/Ashu23042000/coffee-supply-chain/backend/models"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/labstack/echo/v4"
)

// import (
// 	"net/http"

// 	"github.com/Ashu23042000/coffee-supply-chain/backend/models"
// 	"github.com/hyperledger/fabric-gateway/pkg/client"
// 	"github.com/labstack/echo/v4"
// )

var (
	gatewayClient gatewaycli.GatewayClient
)

type Controller struct {
	contract *client.Contract
}

func New(client *client.Contract) *Controller {
	return &Controller{
		contract: client,
	}
}

const (
	chaincodeName = "basic"
	contractName  = "SmartContract"
)

func (c *Controller) Get(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "GetAllAssets",
	}

	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get assets")
	}

	// var assets []models.Asset
	var assets interface{}

	err = json.Unmarshal(evaluateResult, &assets)
	if err != nil {
		fmt.Println("err", err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get assets")
	}

	return ctx.JSON(http.StatusOK, assets)
}

func (c *Controller) GetById(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getAssetById",
	}
	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get asset")
	}

	// var assets []models.Asset
	var assets interface{}

	err = json.Unmarshal(evaluateResult, &assets)
	if err != nil {
		fmt.Println("err", err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get asset")
	}

	return ctx.JSON(http.StatusOK, assets)
}

func (c *Controller) Create(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "InitLedger",
	}
	var args []string
	err := ctx.Bind(args)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	return ctx.JSON(http.StatusOK, "Asset created successfully")
}
