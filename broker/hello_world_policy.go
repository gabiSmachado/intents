package main


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