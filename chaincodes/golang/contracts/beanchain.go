package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	model "beanchain/models"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type BeanChainContract struct {
	contractapi.Contract
}

// BatchInitiation handles the BatchInitiation transaction.
func (t *BeanChainContract) BatchInitiation(ctx contractapi.TransactionContextInterface, adminID string, assetID string) (bool, error) {

	// assetID := utils.GenerateDeterministicID(utils.SharedSecret)

	// assetID := "123"

	fmt.Println(assetID)
	userAsBytes, err := ctx.GetStub().GetState(adminID)
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
	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	batch := model.Asset{
		Id:               assetID,
		Status:           model.Status("CULTIVATOR"),
		CreatedOn:        formattedTime,
		CultivatorIntime: formattedTime,
	}
	fmt.Println(batch)

	assetAsBytes, err := json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal")
	}
	err = ctx.GetStub().PutState(assetID, assetAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to put the asset details")
	}

	eventPayload := map[string]interface{}{
		"previousHandler": adminID,
		"currentHandler":  adminID,
		"Status":          batch.Status,
		"createdOn":       formattedTime,
		"assetId":         assetID,
	}
	fmt.Println("Asset ID :", eventPayload)

	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcess", eventPayloadBytes)

	return true, nil
}

// BatchHarvest handles the BatchHarvest transaction.
func (t *BeanChainContract) BatchProducer(ctx contractapi.TransactionContextInterface, batchID, currentHandlerId, coffeeVariety, temperature, humidity string) (bool, error) {
	// if len(args) != 6 {
	// 	return false, fmt.Errorf("incorrect number of arguments. Expecting 6")
	// }

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

	if batch.Status != "CULTIVATOR" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	// previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerId)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to get previous handler")
	// }
	// var previousHandler model.Participant
	// err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to unmarshal previous handler data")
	// }

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerId)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "PRODUCER" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}
	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	batch.Producer = currentHandler
	batch.CoffeeVariety = coffeeVariety
	batch.Temperature = temperature
	batch.Humidity = humidity
	batch.Status = "PRODUCER"
	batch.ProducerIntime = formattedTime

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		// "previousHandler": previousHandler,
		"currentHandler": currentHandler,
		"batchStatus":    "PRODUCER",
		// "createdDateTime": "".Format(time.RFC3339),
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("PRODUCER produced", eventPayloadBytes)

	return true, nil
}

// BatchFarmInspection handles the BatchFarmInspection transaction.
func (t *BeanChainContract) BatchFarmInspection(ctx contractapi.TransactionContextInterface, batchID, currentHandlerID, typeOfSeed, coffeeFamily, fertilizerUsed string) (bool, error) {
	// if len(args) != 6 {
	// 	return false, fmt.Errorf("incorrect number of arguments. Expecting 6")
	// }

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

	// previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to get previous handler")
	// }
	// var previousHandler model.Participant
	// err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to unmarshal previous handler data")
	// }

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

	// batch.Producer = previousHandler
	batch.Inspector = currentHandler
	batch.TypeofSeed = typeOfSeed
	batch.CoffeeFamily = coffeeFamily
	batch.FertilizerType = fertilizerUsed
	batch.Status = "FARMINSPECTOR"
	// batch.InspectorIntime = ""

	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		// "previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "FARMINSPECTOR",
		"createdDateTime": formattedTime,
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchInspected", eventPayloadBytes)

	return true, nil
}

// BatchProcessor handles the BatchProcessor transaction.
func (t *BeanChainContract) BatchProcessor(ctx contractapi.TransactionContextInterface, batchID, currentHandlerID, processedQuantity, roastingTemp, timeForRoasting, processorAddr string) (bool, error) {

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

	// previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to get previous handler")
	// }
	// var previousHandler model.Participant
	// err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to unmarshal previous handler data")
	// }

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

	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	// batch.Importor = previousHandler
	batch.Processor = currentHandler
	batch.ProcessedQuantity = processedQuantity
	batch.RoastingTemp = roastingTemp
	batch.TimeForRoasting = timeForRoasting
	//  batch.PackagingTime = packagingDatetime
	// batch.ProcessorName = processorName
	batch.ProcessorAddr = processorAddr
	batch.Status = "PROCESSOR"
	batch.ProcessorIntime = formattedTime

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		// "previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "PROCESSOR",
		"createdDateTime": formattedTime,
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchProcessed", eventPayloadBytes)

	return true, nil
}

// BatchExport handles the BatchExport transaction.
func (t *BeanChainContract) BatchExport(ctx contractapi.TransactionContextInterface, batchID, currentHandlerID, exportorQuantity, destAddr, shipName, shipNo string) (bool, error) {

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

	if batch.Status != "PROCESSOR" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	// previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to get previous handler")
	// }
	// var previousHandler model.Participant
	// err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to unmarshal previous handler data")
	// }

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}
	if currentHandler.Role != "EXPORTER" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}
	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	// batch.Producer = previousHandler
	batch.Exportor = currentHandler
	batch.ExportorQuantity = exportorQuantity
	batch.DestAddr = destAddr
	batch.ShipName = shipName
	batch.ShipNumber = shipNo
	//TO DO NEED TO CHECK IT LATER
	// batch.EstimatedTime = estimatedDatetime
	batch.Status = "EXPORTER"
	batch.ExporterIntime = formattedTime

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		// "previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "EXPORTER",
		"createdDateTime": formattedTime,
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchExported", eventPayloadBytes)

	return true, nil
}

// BatchImport handles the BatchImport transaction.
func (t *BeanChainContract) BatchImport(ctx contractapi.TransactionContextInterface, batchID, currentHandlerID, importorQuantity, shipName, shipNo, transportInfo, warehouseName string) (bool, error) {

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

	if batch.Status != "EXPORTER" {
		return false, fmt.Errorf("cannot perform this transaction with current batch status")
	}

	// previousHandlerAsBytes, err := ctx.GetStub().GetState(previousHandlerID)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to get previous handler")
	// }
	// var previousHandler model.Participant
	// err = json.Unmarshal(previousHandlerAsBytes, &previousHandler)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to unmarshal previous handler data")
	// }

	currentHandlerAsBytes, err := ctx.GetStub().GetState(currentHandlerID)
	if err != nil {
		return false, fmt.Errorf("failed to get current handler")
	}
	var currentHandler model.Participant
	err = json.Unmarshal(currentHandlerAsBytes, &currentHandler)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal current handler data")
	}

	if currentHandler.Role != "IMPORTER" {
		return false, fmt.Errorf("not a valid user role")
	}
	if currentHandler.Status != "ACTIVE" {
		return false, fmt.Errorf("user is inactive")
	}
	currentTime := time.Now()
	// Define the format: "YYYY-MM-DD HH"
	format := "2006-01-02 15"
	// Format the current time
	formattedTime := currentTime.Format(format)

	// batch.Exportor = previousHandler
	batch.Importor = currentHandler
	batch.ImportorQuantity = importorQuantity
	batch.ShipName = shipName
	batch.ShipNumber = shipNo
	batch.TransportInfo = transportInfo
	batch.WarehouseName = warehouseName
	batch.Status = "IMPORTER"
	batch.ImporterIntime = formattedTime

	batchAsBytes, err = json.Marshal(batch)
	if err != nil {
		return false, fmt.Errorf("failed to marshal batch asset")
	}

	err = ctx.GetStub().PutState(batchID, batchAsBytes)
	if err != nil {
		return false, fmt.Errorf("failed to update batch asset")
	}

	eventPayload := map[string]interface{}{
		// "previousHandler": previousHandlerID,
		"currentHandler":  currentHandlerID,
		"batchStatus":     "IMPORTER",
		"createdDateTime": formattedTime,
	}
	eventPayloadBytes, _ := json.Marshal(eventPayload)
	ctx.GetStub().SetEvent("BatchImported", eventPayloadBytes)

	return true, nil
}

// AddParticipant stores the participant data in the ledger
func (s *BeanChainContract) AddParticipant(ctx contractapi.TransactionContextInterface, participantData string) error {
	// Convert JSON string to Participant struct
	var participant model.Participant
	err := json.Unmarshal([]byte(participantData), &participant)
	if err != nil {
		return fmt.Errorf("failed to unmarshal participant data: %v", err)
	}

	fmt.Println("HIIIIIIIIIIIIIIII")
	fmt.Println(participant)

	// Store the participant data in the ledger
	participantJSON, err := json.Marshal(participant)
	if err != nil {
		return fmt.Errorf("failed to marshal participant data: %v", err)
	}

	return ctx.GetStub().PutState(participant.User.UserId, participantJSON)
}

// QueryParticipant retrieves a participant from the ledger by ID
func (s *BeanChainContract) QueryParticipant(ctx contractapi.TransactionContextInterface, participantId string) (model.Participant, error) {
	// Retrieve the state from the ledger
	participantJSON, err := ctx.GetStub().GetState(participantId)
	if err != nil {
		return model.Participant{}, fmt.Errorf("failed to get participant: %v", err)
	}
	if participantJSON == nil {
		return model.Participant{}, fmt.Errorf("participant %s does not exist", participantId)
	}

	// Unmarshal JSON into Participant struct
	var participant model.Participant
	err = json.Unmarshal(participantJSON, &participant)
	if err != nil {
		return model.Participant{}, fmt.Errorf("failed to unmarshal participant data: %v", err)
	}

	return participant, nil
}

// GetAllUsers returns all users in the ledger
func (s *BeanChainContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]*model.Participant, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get state by range: %v", err)
	}
	defer resultsIterator.Close()

	var users []*model.Participant
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate: %v", err)
		}

		var user model.Participant
		err = json.Unmarshal(response.Value, &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetAllAssets returns all assets in the ledger
func (s *BeanChainContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*model.Asset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get state by range: %v", err)
	}
	defer resultsIterator.Close()

	var assets []*model.Asset
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate: %v", err)
		}

		var asset model.Asset
		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// GetAsset retrieves an asset from the ledger by its ID
func (s *BeanChainContract) GetAsset(ctx contractapi.TransactionContextInterface, assetId string) (*model.Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return nil, fmt.Errorf("failed to get state for asset ID %s: %v", assetId, err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset with ID %s does not exist", assetId)
	}

	var asset model.Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON for asset ID %s: %v", assetId, err)
	}
	fmt.Println("ASSET")
	fmt.Println(asset)

	return &asset, nil
}
