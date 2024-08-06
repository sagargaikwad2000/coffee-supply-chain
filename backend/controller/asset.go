package controller

import (
	"net/http"

	"github.com/Ashu23042000/coffee-supply-chain/backend/client"
	"github.com/Ashu23042000/coffee-supply-chain/backend/models"
	"github.com/labstack/echo/v4"
)

var (
	gatewayClient = client.New()
)

const (
	chaincodeName = "beanChaincode"
	contractName  = "assetContract"
)

func Get(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getAsset",
	}

	assets, err := gatewayClient.Query(request)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, assets)
}

func GetById(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getAssetById",
	}

	asset, err := gatewayClient.Query(request)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, asset)

}

func Create(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "createAsset",
	}
	var args []string
	err := ctx.Bind(args)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	request.Args = args

	err = gatewayClient.Commit(request)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Asset created successfully")
}
