package models

import (
	"time"
)

type Asset struct {
	Id                string      `json:"batchId"`
	Inspector         Participant `json:"Inspector,omitempty"`
	Producer          Participant `json:"producer,omitempty"`
	Exportor          Participant `json:"exportor,omitempty"`
	Importor          Participant `json:"importor,omitempty"`
	Processor         Participant `json:"processor,omitempty"`
	Cultivator        Participant `json:"cultivator,omitempty"`
	Status            Status      `json:"status"`
	CreatedOn         time.Time   `json:"createdOn,omitempty"`
	CoffeeFamily      string      `json:"coffeeFamily,omitempty"`
	TypeofSeed        string      `json:"typeofSeed,omitempty"`
	FertilizerType    string      `json:"fertilizerType,omitempty"`
	CoffeeVariety     string      `json:"coffeeVariety,omitempty"`
	Temperature       string      `json:"temperature,omitempty"`
	Humidity          string      `json:"humidity,omitempty"`
	DestAddr          string      `json:"destAddr,omitempty"`
	ShipName          string      `json:"shipName,omitempty"`
	ShipNumber        string      `json:"shipNo,omitempty"`
	EstimatedTime     time.Time   `json:"estimatedTime,omitempty"`
	ExporterId        string      `json:"exporterId,omitempty"`
	TransportInfo     string      `json:"transportInfo,omitempty"`
	WarehouseName     string      `json:"warehouseName,omitempty"`
	WarehouseAddr     string      `json:"warehouseAddr,omitempty"`
	ImporterId        string      `json:"importerId,omitempty"`
	RoastingTemp      string      `json:"roastingTemp,omitempty"`
	PackagingTime     time.Time   `json:"packagingTime,omitempty"`
	TimeForRoasting   string      `json:"timeForRoasting,omitempty"`
	ProcessorName     string      `json:"processorName,omitempty"`
	ProcessorAddr     string      `json:"processorAddr,omitempty"`
	ExportorQuantity  string      `json:"exportorQuantity,omitempty"`
	ImportorQuantity  string      `json:"importorQuantity,omitempty"`
	ProcessedQuantity string      `json:"processedQuantity,omitempty"`
	CultivatorIntime  time.Time   `json:"cultivatorIntime,omitempty"`
	InspectorIntime   time.Time   `json:"InspectorIntime,omitempty"`
	ProducerIntime    time.Time   `json:"producerIntime,omitempty"`
	ExporterIntime    time.Time   `json:"exporterIntime,omitempty"`
	ImporterIntime    time.Time   `json:"importerIntime,omitempty"`
	ProcessorIntime   time.Time   `json:"processorIntime,omitempty"`
}
