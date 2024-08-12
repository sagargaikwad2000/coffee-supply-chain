package models

type Participant struct {
	User   User   `json:"user"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
type Status string

type RequestDataProducer struct {
	BatchID          string `json:"batchID"`
	CurrentHandlerID string `json:"currentHandlerId"`
	CoffeeVariety    string `json:"coffeeVariety"`
	Temperature      string `json:"temperature"`
	Humidity         string `json:"humidity"`
}

type RequestDataInspector struct {
	BatchID          string `json:"batchID"`
	CurrentHandlerID string `json:"currentHandlerID"`
	TypeOfSeed       string `json:"typeOfSeed"`
	CoffeeFamily     string `json:"coffeeFamily"`
	FertilizerUsed   string `json:"fertilizerUsed"`
}

type RequestDataProcessor struct {
	BatchID           string `json:"batchID"`
	CurrentHandlerID  string `json:"currentHandlerID"`
	ProcessedQuantity string `json:"processedQuantity"`
	RoastingTemp      string `json:"roastingTemp"`
	TimeForRoasting   string `json:"timeForRoasting"`
	ProcessorAddr     string `json:"processorAddr"`
}

type RequestDataExporter struct {
	BatchID          string `json:"batchID"`
	CurrentHandlerID string `json:"currentHandlerID"`
	ExportorQuantity string `json:"exportorQuantity"`
	DestAddr         string `json:"destAddr"`
	ShipName         string `json:"shipName"`
	ShipNo           string `json:"shipNo"`
}

type RequestDataImporter struct {
	BatchID          string `json:"batchID"`
	CurrentHandlerID string `json:"currentHandlerID"`
	ImportorQuantity string `json:"importorQuantity"`
	ShipName         string `json:"shipName"`
	ShipNo           string `json:"shipNo"`
	TransportInfo    string `json:"transportInfo"`
	WarehouseName    string `json:"warehouseName"`
}
