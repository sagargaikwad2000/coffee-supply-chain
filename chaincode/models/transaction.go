package models

import (
	"time"
)

type Status string

const (
	ADMIN     Status = "ADMIN"
	Inspector Status = "Inspector"
	PRODUCER  Status = "PRODUCER"
	EXPORTOR  Status = "EXPORTOR"
	IMPORTOR  Status = "IMPORTOR"
	PROCESSOR Status = "PROCESSOR"
)

// Inspection transaction
type Inspection struct {
	Asset           `json:"batch"`
	PreviousHandler Participant `json:"previousHandler"`
	CurrentHandler  Participant `json:"currentHandler"`
	TypeofSeed      string      `json:"typeofSeed,omitempty"`
	CoffeeFamily    string      `json:"coffeeFamily,omitempty"`
	FertilizerType  string      `json:"fertilizerType,omitempty"`
	Status          Status      `json:"batchStatus"`
	InspectorIntime time.Time   `json:"InspectorIntime,omitempty"`
}

// Producer transaction
type Producer struct {
	Asset           `json:"batch"`
	PreviousHandler Participant `json:"previousHandler"`
	CurrentHandler  Participant `json:"currentHandler"`
	CoffeeVariety   string      `json:"coffeeVariety,omitempty"`
	Temprature      string      `json:"temprature,omitempty"`
	Humidity        string      `json:"humidity,omitempty"`
	Status          Status      `json:"batchStatus"`
	ProducerIntime  time.Time   `json:"ProducerIntime,omitempty"`
}

// Export transaction
type Export struct {
	Asset            `json:"batch"`
	PreviousHandler  Participant `json:"previousHandler"`
	CurrentHandler   Participant `json:"currentHandler"`
	ExportorQuantity string      `json:"exportorQuantity,omitempty"`
	DestAddr         string      `json:"destAddr,omitempty"`
	ShipName         string      `json:"shipName,omitempty"`
	ShipNumber       string      `json:"shipNo,omitempty"`
	EstimatedTime    time.Time   `json:"estimatedTime"`
	Status           Status      `json:"batchStatus"`
	ExporterIntime   time.Time   `json:"exporterIntime,omitempty"`
}

// Import transaction
type Import struct {
	Asset            `json:"batch"`
	PreviousHandler  Participant `json:"previousHandler"`
	CurrentHandler   Participant `json:"currentHandler"`
	ImportorQuantity string      `json:"importorQuantity,omitempty"`
	ShipName         string      `json:"shipName,omitempty"`
	ShipNumber       string      `json:"shipNo,omitempty"`
	TransportInfo    string      `json:"transportInfo,omitempty"`
	WarehouseName    string      `json:"warehouseName,omitempty"`
	WarehouseAddr    string      `json:"warehouseAddr,omitempty"`
	Status           Status      `json:"batchStatus"`
	ImporterIntime   time.Time   `json:"importerIntime,omitempty"`
}

// Processor transaction
type Processor struct {
	Asset             `json:"batch"`
	PreviousHandler   Participant `json:"previousHandler"`
	CurrentHandler    Participant `json:"currentHandler"`
	ProcessedQuantity string      `json:"processedQuantity,omitempty"`
	RoastingTemp      string      `json:"roastingTemp,omitempty"`
	PackagingTime     time.Time   `json:"packagingTime"`
	TimeForRoasting   string      `json:"timeForRoasting,omitempty"`
	ProcessorName     string      `json:"processorName,omitempty"`
	ProcessorAddr     string      `json:"processorAddr,omitempty"`
	Status            Status      `json:"batchStatus"`
	ProcessorIntime   time.Time   `json:"processorIntime,omitempty"`
}

// Process event
type Process struct {
	PreviousHandler Participant `json:"previousHandler"`
	CurrentHandler  Participant `json:"currentHandler"`
	Status          Status      `json:"batchStatus"`
	CreatedOn       time.Time   `json:"createdOn"`
}

// Initiation transaction
type Initiation struct {
	Cultivator       Participant `json:"cultivator"`
	Id               string      `json:"batchId,omitempty"`
	Status           Status      `json:"batchStatus"`
	CultivatorIntime time.Time   `json:"cultivatorIntime,omitempty"`
}
