package broker

type Policy struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PolicyTypeId int    `json:"policy_type_id"`
	CreateSchema Schema `json:"creat_schema"`
}


type Schema struct {
	Schema               string     `json:"$schema"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}

type Properties struct {
	Threshold Threshold `json:"threshold"`
}

type Threshold struct {
	Type    string `json:"type"`
	Default int    `json:"default"`
}

type PutBody struct {
	RicID         string `json:"ric_id"`
	PolicyId  	  int    `json:"policy_id"`
	ServiceID     string `json:"service_id"`
	PolicyData	  PolicyData`json:"policy_data"`
	PolicyTypeId  string    `json:"policytype_id"`
}

type PolicyData struct {
	Threshold int   `json:"threshold"`
}