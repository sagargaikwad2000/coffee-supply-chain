package models

type Asset struct {
	Id                string      `json:"batchId"`
	Inspector         Participant `json:"Inspector"`
	Producer          Participant `json:"producer"`
	Exportor          Participant `json:"exportor"`
	Importor          Participant `json:"importor"`
	Processor         Participant `json:"processor"`
	Cultivator        Participant `json:"cultivator"`
	Status            Status      `json:"status"`
	CreatedOn         string      `json:"createdOn"`
	CoffeeFamily      string      `json:"coffeeFamily"`
	TypeofSeed        string      `json:"typeofSeed"`
	FertilizerType    string      `json:"fertilizerType"`
	CoffeeVariety     string      `json:"coffeeVariety"`
	Temperature       string      `json:"temperature"`
	Humidity          string      `json:"humidity"`
	DestAddr          string      `json:"destAddr"`
	ShipName          string      `json:"shipName"`
	ShipNumber        string      `json:"shipNo"`
	EstimatedTime     string      `json:"estimatedTime"`
	ExporterId        string      `json:"exporterId"`
	TransportInfo     string      `json:"transportInfo"`
	WarehouseName     string      `json:"warehouseName"`
	WarehouseAddr     string      `json:"warehouseAddr"`
	ImporterId        string      `json:"importerId"`
	RoastingTemp      string      `json:"roastingTemp"`
	PackagingTime     string      `json:"packagingTime"`
	TimeForRoasting   string      `json:"timeForRoasting"`
	ProcessorName     string      `json:"processorName"`
	ProcessorAddr     string      `json:"processorAddr"`
	ExportorQuantity  string      `json:"exportorQuantity"`
	ImportorQuantity  string      `json:"importorQuantity"`
	ProcessedQuantity string      `json:"processedQuantity"`
	CultivatorIntime  string      `json:"cultivatorIntime"`
	InspectorIntime   string      `json:"InspectorIntime"`
	ProducerIntime    string      `json:"producerIntime"`
	ExporterIntime    string      `json:"exporterIntime"`
	ImporterIntime    string      `json:"importerIntime"`
	ProcessorIntime   string      `json:"processorIntime"`
}
