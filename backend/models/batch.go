package models

type Batch struct {
	DocType     string `json:"docType"`
	Id          string `json:"batchId"`
	CoffeeType  string `json:"coffeeType"`
	Location    string `json:"location"`
	CreatedOn   string `json:"createdOn"`
	Status      string `json:"status"`
	Quantity    string `json:"quantity"`
	CostPerKg   string `json:"costPerKg"`
	ProducerId  string `json:"producerId"`
	InspectorId string `json:"inspectorId"`
	ProcessorId string `json:"processorId"`
	ExporterId  string `json:"exporterId"`
	ImporterId  string `json:"importerId"`
}
