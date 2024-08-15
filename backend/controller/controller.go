package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (c *Controller) GetBatches(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getAllBatches",
	}

	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName)
	fmt.Println("evaluateResult1", string(evaluateResult))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get batches")
	}

	// var assets []models.Asset
	var batches []models.Batch

	err = json.Unmarshal(evaluateResult, &batches)
	if err != nil {
		fmt.Println("err", err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get batches")
	}

	return ctx.JSON(http.StatusOK, batches)
}

func (c *Controller) GetBatchById(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getBatch",
	}
	id := ctx.Param("id")
	fmt.Println("Id", id)
	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get asset")
	}
	fmt.Println("evaluateResult2", string(evaluateResult))

	var batch models.Batch

	err = json.Unmarshal(evaluateResult, &batch)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get batch")
	}
	fmt.Sprintf("Batchh %+v \n", batch)

	return ctx.JSON(http.StatusOK, batch)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func (c *Controller) CreateBatch(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "createBatch",
	}

	var batchBody models.Batch
	if err := ctx.Bind(&batchBody); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	fmt.Println("batchBody", batchBody)

	batchBody.Id = GenerateUUID()
	batchBody.DocType = "Batch"
	batchBody.Status = "Created"
	batchBytes, err := json.Marshal(batchBody)
	if err != nil {
		fmt.Println("err1")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("bodyB", batchBody)
	request.Args = []string{batchBody.Id, string(batchBytes)}

	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		fmt.Println("err3")
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Batch created successfully",
		"batchId": batchBody.Id,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) UpdateBatch(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "updateBatch",
	}

	var batchBody models.Batch
	if err := ctx.Bind(&batchBody); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	fmt.Println("batchBody", batchBody)

	// batchBody.Id = GenerateUUID()
	// batchBody.DocType = "Batch"
	batchBytes, err := json.Marshal(batchBody)
	if err != nil {
		fmt.Println("err1")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("bodyB", batchBody)
	request.Args = []string{batchBody.Id, string(batchBytes)}

	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		fmt.Println("err3")
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Batch updated successfully",
		"batchId": batchBody.Id,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) CreateUser(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "createUser",
	}

	var userBody models.User
	if err := ctx.Bind(&userBody); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	fmt.Println("userBody", userBody)

	userBody.DocType = "User"

	userBytes, err := json.Marshal(userBody)
	if err != nil {
		fmt.Println("err1")
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("bodyB", userBody)
	request.Args = []string{userBody.Id, string(userBytes)}

	_, err = c.contract.SubmitTransaction(request.MethodName, request.Args...)
	if err != nil {
		fmt.Println("err3")
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	// Construct the response JSON
	response := map[string]interface{}{
		"message": "User created successfully",
		"userId":  userBody.Id,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetUserById(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getUser",
	}
	id := ctx.Param("id")
	fmt.Println("Id", id)
	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get user")
	}
	fmt.Println("evaluateResult2", string(evaluateResult))

	var user models.User

	err = json.Unmarshal(evaluateResult, &user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get user")
	}

	fmt.Println("Batchh", user)

	return ctx.JSON(http.StatusOK, user)
}

func (c *Controller) UserLogin(ctx echo.Context) error {
	request := models.Request{
		ChaincodeName: chaincodeName,
		ContractName:  contractName,
		MethodName:    "getUser",
	}
	var batchBody models.User
	if err := ctx.Bind(&batchBody); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}
	evaluateResult, err := c.contract.EvaluateTransaction(request.MethodName, batchBody.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get user")
	}
	fmt.Println("evaluateResult2", string(evaluateResult))

	var user models.User
	err = json.Unmarshal(evaluateResult, &user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get user")
	}

	if user.Password == batchBody.Password {
		fmt.Println("user", user)
		return ctx.JSON(http.StatusOK, user)

	} else {
		return ctx.JSON(http.StatusInternalServerError, "Password does not match")
	}
}
