package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	model "github.com/hyperledger/fabric-samples/chaincode/fabcar/go/models"
	utils "github.com/hyperledger/fabric-samples/chaincode/fabcar/go/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Beanchain struct {
	contractapi.Contract
}

// BatchInitiation handles the BatchInitiation transaction.
func (t *Beanchain) BatchInitiation(ctx contractapi.TransactionContextInterface, assetID, userID string) (bool, error) {

	if assetID == "" {
		var err error
		assetID, err = utils.GenerateSecureRandomString()
		if err != nil {
			return false, fmt.Errorf("")
		}
	}

	assetAsBytes, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return false, fmt.Errorf("failed to get the asset data from the ledger")
	}
	if assetAsBytes != nil {
		return false, fmt.Errorf("Batch with ID " + assetID + " already exists")
	}

	userAsBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return false, fmt.Errorf("failed to get user")
	}
	var user model.Participant
	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal")
	}
	if user.Role != "ADMIN" {
		return false, fmt.Errorf("not a valid role")
	}
	if user.Status != "ACTIVE" {
		return false, fmt.Errorf("user is INACTIVE")
	}

	batch := model.Asset{
		Id:               assetID,
		Status:           "ADMIN",
		CreatedOn:        time.Now(),
		Cultivator:       user,
		CultivatorIntime: time.Now(),
	}

	assetAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal")
	}

	err = ctx.GetStub().PutState(assetID, assetAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to put the asset details")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": userID,
		"currentHandler":  userID,
		"Status":          batch.Status,
		"createdOn":       time.Now(),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchFarmInspection handles the BatchFarmInspection transaction.
func (t *Beanchain) BatchFarmInspection(ctx contractapi.TransactionContextInterface, args []string) (bool, error) {
	if len(args) != 6 {
		return false, fmt.Errorf("incorrect number of arguments. Expecting 6")
	}

	batchID := args[0]
	previousHandlerID := args[1]
	currentHandlerID := args[2]
	typeOfSeed := args[3]
	coffeeFamily := args[4]
	fertilizerUsed := args[5]

	batchAsBytes, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to get batch asset")
	}
	if batchAsBytes == nil {
		return false, fmt.Errorf("Batch with ID " + batchID + " does not exist")
	}

	var batch model.Asset
	err = json.Unmarshal(batchAsBytes, &batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal batch data")
	}

	if batch.Status != "ADMIN" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get previous handler")
	}
	var previousHandler model.Participant
	err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal previous handler data")
	}

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "FARMINSPECTOR" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}

	batch.Cultivator = previousHandler
	batch.Inspector = currentHandler
	batch.TypeofSeed = typeOfSeed
	batch.CoffeeFamily = coffeeFamily
	batch.FertilizerType = fertilizerUsed
	batch.Status = "FARMINSPECTOR"
	batch.InspectorIntime = time.Now()

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "FARMINSPECTOR",
		"createdDateTime": time.Now(),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchHarvest handles the BatchHarvest transaction.
func (t *Beanchain) BatchHarvest(ctx contractapi.TransactionContextInterface, args []string) (bool, error) {
	if len(args) != 6 {
		return false, fmt.Errorf("incorrect number of arguments. Expecting 6")
	}

	batchID := args[0]
	previousHandlerID := args[1]
	currentHandlerID := args[2]
	coffeeVariety := args[3]
	temperature := args[4]
	humidity := args[5]

	batchAsBytes, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to get batch asset")
	}
	if batchAsBytes == nil {
		return false, fmt.Errorf("Batch with ID " + batchID + " does not exist")
	}

	var batch model.Asset
	err = json.Unmarshal(batchAsBytes, &batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal batch data")
	}

	if batch.Status != "FARMINSPECTOR" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get previous handler")
	}
	var previousHandler model.Participant
	err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal previous handler data")
	}

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "HARVESTOR" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}

	batch.Inspector = previousHandler
	batch.Processor = currentHandler
	batch.CoffeeVariety = coffeeVariety
	batch.Temperature = temperature
	batch.Humidity = humidity
	batch.Status = "PRODUCER"
	batch.ProducerIntime = time.Now()

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "HARVESTOR",
		"createdDateTime": time.Now().Format(time.RFC3339),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchExport handles the BatchExport transaction.
func (t *Beanchain) BatchExport(ctx contractapi.TransactionContextInterface, args []string) (bool, error) {
	if len(args) != 8 {
		return false, fmt.Errorf("incorrect number of arguments. Expecting 8")
	}

	batchID := args[0]
	previousHandlerID := args[1]
	currentHandlerID := args[2]
	exportorQuantity := args[3]
	destAddr := args[4]
	shipName := args[5]
	shipNo := args[6]
	// estimatedTime := args[7]

	batchAsBytes, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to get batch asset")
	}
	if batchAsBytes == nil {
		return false, fmt.Errorf("Batch with ID " + batchID + " does not exist")
	}

	var batch model.Asset
	err = json.Unmarshal(batchAsBytes, &batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal batch data")
	}

	if batch.Status != "PRODUCER" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get previous handler")
	}
	var previousHandler model.Participant
	err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal previous handler data")
	}

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "EXPORTOR" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}

	batch.Producer = previousHandler
	batch.Exportor = currentHandler
	batch.ExportorQuantity = exportorQuantity
	batch.DestAddr = destAddr
	batch.ShipName = shipName
	batch.ShipNumber = shipNo
	//TO DO NEED TO CHECK IT LATER
	// batch.EstimatedTime = estimatedDatetime
	batch.Status = "EXPORTOR"
	batch.ExporterIntime = time.Now()

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "EXPORTOR",
		"createdDateTime": time.Now().Format(time.RFC3339),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchImport handles the BatchImport transaction.
func (t *Beanchain) BatchImport(ctx contractapi.TransactionContextInterface, args []string) (bool, error) {
	if len(args) != 8 {
		return false, fmt.Errorf("incorrect number of arguments. Expecting 8")
	}

	batchID := args[0]
	previousHandlerID := args[1]
	currentHandlerID := args[2]
	importorQuantity := args[3]
	shipName := args[4]
	shipNo := args[5]
	transportInfo := args[6]
	warehouseName := args[7]

	batchAsBytes, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to get batch asset")
	}
	if batchAsBytes == nil {
		return false, fmt.Errorf("batch with ID " + batchID + " does not exist")
	}

	var batch model.Asset
	err = json.Unmarshal(batchAsBytes, &batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal batch data")
	}

	if batch.Status != "EXPORTOR" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get previous handler")
	}
	var previousHandler model.Participant
	err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal previous handler data")
	}

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "IMPORTOR" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}

	batch.Exportor = previousHandler
	batch.Importor = currentHandler
	batch.ImportorQuantity = importorQuantity
	batch.ShipName = shipName
	batch.ShipNumber = shipNo
	batch.TransportInfo = transportInfo
	batch.WarehouseName = warehouseName
	batch.Status = "IMPORTOR"
	batch.ImporterIntime = time.Now()

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "IMPORTOR",
		"createdDateTime": time.Now().Format(time.RFC3339),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchProcessor handles the BatchProcessor transaction.
func (t *Beanchain) BatchProcessor(ctx contractapi.TransactionContextInterface, args []string) (bool, error) {
	if len(args) != 10 {
		return false, fmt.Errorf("incorrect number of arguments. Expecting 10")
	}

	batchID := args[0]
	previousHandlerID := args[1]
	currentHandlerID := args[2]
	processedQuantity := args[3]
	roastingTemp := args[4]
	timeForRoasting := args[5]
	// packagingDatetime := args[6]
	processorName := args[7]
	processorAddr := args[8]
	batchStatus := args[9]

	batchAsBytes, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to get batch asset")
	}
	if batchAsBytes == nil {
		return false, fmt.Errorf("Batch with ID " + batchID + " does not exist")
	}

	var batch model.Asset
	err = json.Unmarshal(batchAsBytes, &batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal batch data")
	}

	if batch.Status != "IMPORTOR" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get previous handler")
	}
	var previousHandler model.Participant
	err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal previous handler data")
	}

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "PROCESSOR" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}

	batch.Importor = previousHandler
	batch.Processor = currentHandler
	batch.ProcessedQuantity = processedQuantity
	batch.RoastingTemp = roastingTemp
	batch.TimeForRoasting = timeForRoasting
	// batch.PackagingTime = packagingDatetime
	batch.ProcessorName = processorName
	batch.ProcessorAddr = processorAddr
	batch.Status = model.Status(batchStatus)
	batch.ProcessorIntime = time.Now()

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     batchStatus,
		"createdDateTime": time.Now().Format(time.RFC3339),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}
