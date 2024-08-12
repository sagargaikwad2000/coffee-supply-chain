package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Ashu23042000/coffee-supply-chain/backend/models"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/labstack/echo/v4"
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
	Producer      = "PRODUCER"
	Inspector     = "FARMINSTPECTOR"
	Processor     = "PROCESSOR"
	Exporter      = "EXPORTER"
	Importer      = "IMPORTER"
	Admin         = "ADMIN"
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
	var assets []models.Asset

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
		MethodName:    "GetAsset",
	}
	id := ctx.Param("id")
	fmt.Println("ID", id)
	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get asset")
	}

	var assets models.Asset

	err = json.Unmarshal(evaluateResult, &assets)
	if err != nil {
		fmt.Println("err", err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get asset")
	}

	return ctx.JSON(http.StatusOK, assets)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func (c *Controller) Create(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchInitiation",
	}
	assetId := GenerateUUID()
	args := []string{"103", assetId}
	err := ctx.Bind(args)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset created successfully",
		"assetID": assetId,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) ProducerUpdate(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchProducer",
	}
	var requestData models.RequestDataProducer

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	args := []string{requestData.BatchID, "101", requestData.CoffeeVariety, requestData.Temperature, requestData.Humidity}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	var err error

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset Updated successfully",
		"assetID": requestData.BatchID,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) InspectorUpdate(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchFarmInspection",
	}
	var requestData models.RequestDataInspector

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	args := []string{requestData.BatchID, "204", requestData.TypeOfSeed, requestData.CoffeeFamily, requestData.FertilizerUsed}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	var err error

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset Updated successfully",
		"assetID": requestData.BatchID,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) ProcessorUpdate(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchProcessor",
	}
	var requestData models.RequestDataProcessor

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	args := []string{requestData.BatchID, "205", requestData.ProcessedQuantity, requestData.RoastingTemp, requestData.RoastingTemp, requestData.ProcessorAddr}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	var err error

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset Updated successfully",
		"assetID": requestData.BatchID,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) ExporterUpdate(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchExport",
	}
	var requestData models.RequestDataExporter

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	args := []string{requestData.BatchID, "206", requestData.ExportorQuantity, requestData.DestAddr, requestData.ShipName, requestData.ShipNo}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	var err error

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset Updated successfully",
		"assetID": requestData.BatchID,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) ImporterUpdate(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "BatchImport",
	}
	var requestData models.RequestDataImporter

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	args := []string{requestData.BatchID, "207", requestData.ImportorQuantity, requestData.ShipName, requestData.ShipNo, requestData.TransportInfo, requestData.WarehouseName}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	var err error

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Asset Updated successfully",
		"assetID": requestData.BatchID,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) AddUser(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "AddParticipant",
	}
	var requestData models.Participant

	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println("Error reading body", err)
	}

	fmt.Println("body:", string(body))

	// Bind the incoming JSON body to the struct
	// if err := ctx.Bind(&requestData); err != nil {
	// 	return ctx.String(http.StatusBadRequest, "Invalid request payload")
	// }

	if err := json.Unmarshal(body, &requestData); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	requestDataString, err := json.Marshal(requestData)
	if err != nil {

	}

	args := []string{string(requestDataString)}
	// err := ctx.Bind(args)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }

	fmt.Println("args", args)

	request.Args = args
	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "User created successfully ",
		"UserId":  requestData.User.UserId,
	}
	return ctx.JSON(http.StatusOK, response)
}

// QueryParticipantHandler retrieves a participant from the ledger by ID
func (c *Controller) GetUser(ctx echo.Context) error {
	participantId := ctx.Param("id")
	// Query the chaincode
	participantJSON, err := c.contract.EvaluateTransaction("QueryParticipant", participantId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to query chaincode: %v", err))
	}

	var participant models.Participant
	err = json.Unmarshal(participantJSON, &participant)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to unmarshal participant data")
	}

	return ctx.JSON(http.StatusOK, participant)
}

// GetAllAssetsHandler retrieves all assets from the ledger and filters out those without an AssetID
func (c *Controller) GetAllAssets(ctx echo.Context) error {
	// Query the chaincode
	assetsJSON, err := c.contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to query chaincode: %v", err))
	}

	// Unmarshal the JSON byte string into a slice of Asset structs
	var assets []*models.Asset
	err = json.Unmarshal(assetsJSON, &assets)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to unmarshal asset data")
	}

	// Filter assets to include only those with a non-empty AssetID
	var filteredAssets []*models.Asset
	for _, asset := range assets {
		if strings.TrimSpace(asset.Id) != "" {
			filteredAssets = append(filteredAssets, asset)
		}
	}

	return ctx.JSON(http.StatusOK, filteredAssets)
}

func (c *Controller) GetAllUsers(ctx echo.Context) error {
	// Query the chaincode
	participantJSON, err := c.contract.EvaluateTransaction("GetAllUsers")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to query chaincode: %v", err))
	}

	var participants []*models.Participant
	err = json.Unmarshal(participantJSON, &participants)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to unmarshal participant data")
	}
	var filteredUsers []*models.Participant
	for _, asset := range participants {
		if strings.TrimSpace(asset.Role) != "" {
			filteredUsers = append(filteredUsers, asset)
		}
	}

	return ctx.JSON(http.StatusOK, filteredUsers)
}

func (c *Controller) LoginUser(ctx echo.Context) error {
	// Bind request data to LoginRequest struct
	var loginRequest struct {
		UserID   string `json:"userId"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	// Query all users from the ledger
	usersJSON, err := c.contract.EvaluateTransaction("GetAllUsers")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to query chaincode: %v", err))
	}

	// Unmarshal the JSON byte string into a slice of User structs
	var users []models.Participant
	err = json.Unmarshal(usersJSON, &users)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to unmarshal user data")
	}

	// Check if the user exists and the password matches
	for _, user := range users {
		if user.User.UserId == loginRequest.UserID && user.User.Password == loginRequest.Password && user.Role != Admin {
			return ctx.JSON(http.StatusOK, user)
		}
	}

	return ctx.String(http.StatusUnauthorized, "Invalid user ID or password")
}

func (c *Controller) LoginAdmin(ctx echo.Context) error {
	// Bind request data to LoginRequest struct
	var loginRequest struct {
		UserID   string `json:"userId"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	// Query all users from the ledger
	usersJSON, err := c.contract.EvaluateTransaction("GetAllUsers")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to query chaincode: %v", err))
	}

	// Unmarshal the JSON byte string into a slice of User structs
	var users []models.Participant
	err = json.Unmarshal(usersJSON, &users)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to unmarshal user data")
	}

	// Check if the user exists and the password matches
	for _, user := range users {
		if user.User.UserId == loginRequest.UserID && user.User.Password == loginRequest.Password && user.Role == Admin {
			return ctx.JSON(http.StatusOK, user)
		}
	}

	return ctx.String(http.StatusUnauthorized, "Invalid User")
}
