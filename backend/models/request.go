package models

type Request struct {
	ChaincodeName string        `json:"chaincodeName"`
	ContractName  string        `json:"contractName"`
	MethodName    string        `json:"methodName"`
	Args          []string      `json:"args"`
	TrasientData  []interface{} `json:"transientData"`
}
